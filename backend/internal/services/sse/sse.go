package sse

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"opflow-backend/internal/models"
)

type SSEService struct {
	clients      map[chan string]struct{}
	newClients   chan chan string
	deadClients  chan chan string
	messages     chan models.Hotspot
}

var (
	instance *SSEService
	once     sync.Once
)

func GetSSEService() *SSEService {
	once.Do(func() {
		instance = &SSEService{
			clients:     make(map[chan string]struct{}),
			newClients:  make(chan chan string),
			deadClients: make(chan chan string),
			messages:    make(chan models.Hotspot, 10),
		}
		go instance.handleMessages()
	})
	return instance
}

func (s *SSEService) handleMessages() {
	for {
		select {
		case client := <-s.newClients:
			s.clients[client] = struct{}{}
			fmt.Println("Added new client")

		case client := <-s.deadClients:
			delete(s.clients, client)
			close(client)
			fmt.Println("Removed client")

		case message := <-s.messages:
			for client := range s.clients {
				select {
				case client <- fmt.Sprintf("data: %s\n\n", message.Title):
				default:
					go func() { s.deadClients <- client }()
				}
			}
		}
	}
}

func (s *SSEService) SendMessage(hotspot models.Hotspot) {
	s.messages <- hotspot
}

func (s *SSEService) ServeHTTP(c *gin.Context) {
	clientChan := make(chan string)
	s.newClients <- clientChan
	defer func() {
		s.deadClients <- clientChan
	}()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	for {
		select {
		case msg := <-clientChan:
			c.Writer.Write([]byte(msg))
			c.Writer.Flush()
		case <-c.Done():
			return
		}
	}
}