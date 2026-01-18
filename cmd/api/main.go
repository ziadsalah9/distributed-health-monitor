package main

import (
	"log"
	"github.com/gin-gonic/gin"

    "distributed-health-monitor/internal/db"
	"distributed-health-monitor/internal/models"
	"distributed-health-monitor/internal/handlers"
)


	func main() {
	
	
		db.ConnectPostgres()

		err:= db.DB.AutoMigrate(&models.Service{}, &models.HealthLog{})

		if err != nil {
			log.Fatal("Auto Migration failed:", err)
		}


		db.SeedData()
		r := gin.Default()

		r.GET("/health", func(c *gin.Context) {
		
		  // gin.H is a shortcut for map[string]interface{}
			c.JSON(200, gin.H{
			"status": "API running",
				})
			})

	    		// Service routes
		serviceRoutes := r.Group("/services")
		{
			serviceRoutes.POST("/", handlers.RegisterService)
			serviceRoutes.GET("/", handlers.ListServices)
			serviceRoutes.GET("/:id/logs", handlers.GetServiceLogs) //  show logs for a specific service
		}		

			r.Run(":8088")
	}