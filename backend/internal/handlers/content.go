package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"opflow-backend/internal/models"
	"opflow-backend/internal/services/database"
)

type ContentHandler struct {
	dbService *database.DatabaseService
}

func NewContentHandler(dbService *database.DatabaseService) *ContentHandler {
	return &ContentHandler{
		dbService: dbService,
	}
}

func (h *ContentHandler) GetContents(c *gin.Context) {
	var contents []models.Content
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	result := h.dbService.GetDB().Limit(limit).Offset(offset).Find(&contents)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": contents,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": result.RowsAffected,
		},
	})
}

func (h *ContentHandler) ScheduleContent(c *gin.Context) {
	id := c.Param("id")
	var content models.Content

	if err := h.dbService.GetDB().First(&content, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		return
	}

	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content.ScheduledAt = &schedule.ScheduledAt
	content.Status = "scheduled"

	if err := h.dbService.GetDB().Save(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Content scheduled successfully",
		"content": content,
	})
}
