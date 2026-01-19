package dtos

import "time"

type HealthLogResponseDTO struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`     
	ResponseTimeAsMs int       `json:"response_time"`
	Timestamp time.Time `json:"timestamp"`
}