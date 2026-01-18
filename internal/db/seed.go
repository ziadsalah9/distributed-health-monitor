package db
import (
	"log"
	"distributed-health-monitor/internal/models"
	"time"

)
func SeedData() {
	// Check if there are any existing services
	var count int64
	DB.Model(&models.Service{}).Count(&count)	
	if count > 0 {
		log.Println("Database already seeded")
		return
	}
	// Seed initial services
	services := []models.Service{
		{Name: "Google", URL: "https://www.google.com", Interval: 60, LastStatus: "UP", CreatedAt: time.Now(),LastCheck: time.Now(), Timeout: 5},
		{Name: "GitHub", URL: "https://www.github.com", Interval: 120, LastStatus: "UP", CreatedAt: time.Now() ,LastCheck: time.Now(),Timeout: 10},
	  	{Name: "NonExistent", URL: "http://nonexistent.example.com", Interval: 90, LastStatus: "DOWN", CreatedAt: time.Now() ,LastCheck: time.Now(), Timeout: 30},
		{Name: "DelayLinkToTest",URL:"https://httpbin.org/delay/15",Interval: 20 , LastStatus: "UP",CreatedAt: time.Now(),LastCheck: time.Now() , Timeout: 5},
	
		}   
	if err := DB.Create(&services).Error; err != nil {
		log.Printf("Error seeding services: %v", err)
		return
	}
	log.Println("Database seeded successfully")
}
