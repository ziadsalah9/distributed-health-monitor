package dtos

import "time"

type HealthLogResponseDTO struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`     
	State     string    `json:"state"`      
	LatencyMs int       `json:"latency_ms"`
	CheckedAt time.Time `json:"checked_at"`
}