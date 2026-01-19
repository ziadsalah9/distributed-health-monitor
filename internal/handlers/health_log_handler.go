package handlers

import (
	"distributed-health-monitor/internal/db"
	"distributed-health-monitor/internal/dtos"
	"distributed-health-monitor/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get  Logs by Service ID
func GetServiceLogs(c *gin.Context) {
	serviceID := c.Param("id")
	var logs []models.HealthLog

	result := db.DB.Where("service_id = ?", serviceID).Order("checked_at desc").Limit(100).Find(&logs)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch logs"})
		return
	}

	var response []dtos.HealthLogResponseDTO
	for _, l := range logs {
		response = append(response, dtos.HealthLogResponseDTO{
			ID:        l.ID,
			Status:    l.Status,
			ResponseTimeAsMs: l.LatencyMs,
			Timestamp: l.CheckedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}