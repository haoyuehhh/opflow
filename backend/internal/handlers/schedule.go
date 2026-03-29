package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"opflow-backend/internal/models"
	"opflow-backend/internal/services/database"
)

type ScheduleHandler struct {
	dbService *database.DatabaseService
}

func NewScheduleHandler(dbService *database.DatabaseService) *ScheduleHandler {
	return &ScheduleHandler{
		dbService: dbService,
	}
}

func (h *ScheduleHandler) GetSchedules(c *gin.Context) {
	var schedules []models.Schedule
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	result := h.dbService.GetDB().Limit(limit).Offset(offset).Find(&schedules)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": schedules,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": result.RowsAffected,
		},
	})
}
