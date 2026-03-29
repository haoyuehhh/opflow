package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "opflow-backend/internal/models"
    "opflow-backend/internal/services/content"
    "opflow-backend/internal/services/crawler"
)

type HotspotHandler struct {
    crawler   *crawler.Service
    generator *content.Generator
}

func NewHotspotHandler(c *crawler.Service, g *content.Generator) *HotspotHandler {
    return &HotspotHandler{crawler: c, generator: g}
}

func (h *HotspotHandler) List(c *gin.Context) {
    zhihu, _ := h.crawler.FetchZhihu()
    weibo, _ := h.crawler.FetchWeibo()
    
    all := append(zhihu, weibo...)
    
    if len(all) == 0 {
        all = []models.Hotspot{
            {ID: "1", Title: "Test Hotspot", Source: "zhihu", Score: 100},
        }
    }
    
    c.JSON(http.StatusOK, gin.H{"data": all})
}

func (h *HotspotHandler) Generate(c *gin.Context) {
    id := c.Param("id")
    content, _ := h.generator.Generate(models.Hotspot{Title: "Hotspot " + id}, "xiaohongshu")
    c.JSON(http.StatusOK, gin.H{"content": content})
}
