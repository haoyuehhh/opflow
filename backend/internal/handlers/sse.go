package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opflow-backend/internal/models"
	"opflow-backend/internal/services/sse"
)

type SSEHandler struct {
	sseService *sse.SSEService
}

func NewSSEHandler(sseService *sse.SSEService) *SSEHandler {
	return &SSEHandler{
		sseService: sseService,
	}
}

func (h *SSEHandler) StreamHotspots(c *gin.Context) {
	h.sseService.ServeHTTP(c)
}
