package models

import "time"

type Service struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100);not null"`
	URL         string
    Interval 	int
    LastStatus  string
    CreatedAt   time.Time

	HealthLogs  []HealthLog `gorm:"foreignKey:ServiceID"`
}