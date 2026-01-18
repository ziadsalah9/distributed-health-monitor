package scheduler

import (
	"context"
	"distributed-health-monitor/internal/db"
	"distributed-health-monitor/internal/models"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)


func StartScheduler(channel *amqp.Channel) {	

	log.Println("Scheduler started...")

	// // Run every 10 seconds
	// ticker := time.NewTicker(10 * time.Second)

	// Run every second
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		var services []models.Service
		//db.DB.Find(&services)   // get all services from db then push it to queue
	
		db.DB.Where("extract(epoch from (now() - last_check)) >= interval OR last_check IS NULL").Find(&services)

		
		for _ , s := range services {


			// update last_check time
			db.DB.Model(&s).Update("last_check", time.Now())

			message := fmt.Sprintf("%d", s.ID)
			err:= channel.PublishWithContext(context.Background(),"","health_checks",false,false,amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message)})
		
			if err != nil {
				log.Printf("Failed to publish message for service ID %d: %v", s.ID, err)
			}
		
		}
		
if len(services) > 0 {
            log.Printf("Dispatched %d services that were due for check", len(services))
        }

	}
	

}