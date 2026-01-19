package models

import "time"

type HealthLog struct {
	ID        uint `gorm:"primaryKey"`
	ServiceID uint
	Status    string // UP - DOWN
	LatencyMs int
	CheckedAt time.Time

	//optional Navigation property
	Service   Service   `gorm:"foreignKey:ServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}