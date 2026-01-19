package main

import (
	"log"
	"github.com/gin-gonic/gin"

    "distributed-health-monitor/internal/db"
	"distributed-health-monitor/internal/models"
	"distributed-health-monitor/internal/repository"
	"distributed-health-monitor/internal/handlers"

	"distributed-health-monitor/internal/scheduler"
	"distributed-health-monitor/internal/worker"
	"distributed-health-monitor/internal/websocket"

	 amqp "github.com/rabbitmq/amqp091-go"

)


	func main() {
	
	
		db.ConnectPostgres()

		err:= db.DB.AutoMigrate(&models.Service{}, &models.HealthLog{})

		if err != nil {
			log.Fatal("Auto Migration failed:", err)
		}


		db.SeedData()


 		// RabbitMQ connection

		conn , err :=amqp.Dial("amqp://guest:guest@localhost:5672/")

		if err!= nil{
			log.Fatal( "Failed to connect to RabbitMQ:", err)
		}
		defer conn.Close()

		channel, err := conn.Channel()
		if err != nil {
			log.Fatal("Failed to open a channel:", err)
		}
		defer channel.Close()

		// Declare a queue for health checks
		channel.QueueDeclare(
			"health_checks", // name
			true,           
			false,		  
			false,          
			false,          
			nil,            
		)


		// inject repo
		serviceRepo := repository.NewServiceRepository(db.DB)
		servicehandlers := handlers.NewServiceHandler(serviceRepo)
		
		hub:= websocket.NewHub()
		go hub.Run()


	// make it as background process (sheduler job in java  , background job in c#)
	  go scheduler.StartScheduler(channel)
	  go worker.StartWorker(conn,hub)

		r := gin.Default()

		r.GET("/health", func(c *gin.Context) {
		
		  // gin.H is a shortcut for map[string]interface{}
			c.JSON(200, gin.H{
			"status": "API running",
				})
			})

			r.GET("/ws", func(c *gin.Context) {
		hub.HandleWS(c.Writer, c.Request)
	})

	    		// Service routes
		serviceRoutes := r.Group("/services")
		{
			serviceRoutes.POST("/", servicehandlers.RegisterService)
			serviceRoutes.GET("/", servicehandlers.ListServices)
			serviceRoutes.GET("/:id/logs", servicehandlers.GetServiceLogs) //  show logs for a specific service
		}		

			r.Run(":8088")
	}