package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"opflow-backend/internal/models"
	"opflow-backend/internal/services/database"
)

type HotspotHandler struct {
	dbService *database.DatabaseService
}

func NewHotspotHandler(dbService *database.DatabaseService) *HotspotHandler {
	return &HotspotHandler{
		dbService: dbService,
	}
}

func (h *HotspotHandler) GetHotspots(c *gin.Context) {
	var hotspots []models.Hotspot
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	result := h.dbService.GetDB().Limit(limit).Offset(offset).Find(&hotspots)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hotspots,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": result.RowsAffected,
		},
	})
}

func (h *HotspotHandler) GenerateContent(c *gin.Context) {
	id := c.Param("id")
	var hotspot models.Hotspot

	if err := h.dbService.GetDB().First(&hotspot, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotspot not found"})
		return
	}

	// 这里应该调用LLM服务生成内容
	// 模拟生成过程
	time.Sleep(2 * time.Second)

	content := models.Content{
		HotspotID: hotspot.ID,
		Title:     "Generated: " + hotspot.Title,
		Content:   "Generated content based on hotspot: " + hotspot.Title,
		Platform:  "xiaohongshu",
		Status:    "draft",
	}

	if err := h.dbService.GetDB().Create(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Content generation started",
		"content_id": content.ID,
	})
}
