package worker

import (
    "net/http"
	"distributed-health-monitor/internal/db"
	"distributed-health-monitor/internal/models"
	"log"
	"strconv"
	"time"
	"distributed-health-monitor/internal/websocket"
	
amqp "github.com/rabbitmq/amqp091-go"
)

func StartWorker(c*amqp.Connection, hub*websocket.Hub) {

	ch,err:=c.Channel()
	if err!=nil{
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	
	q,err:= ch.QueueDeclare("health_checks",true,false,false,false,nil)

	msg,err :=ch.Consume(q.Name,"",true,false,false,false,nil)

	log.Println(" Worker is waiting for messages...")


	for d:=range msg{

     serviceId,_ := strconv.Atoi(string(d.Body))
	 log.Printf("Received a message for Service ID: %d", serviceId)

	 processCheck(uint(serviceId),hub)
	}



}


func processCheck(id uint, hub *websocket.Hub) {


	var service models.Service

	if err := db.DB.First(&service,id).Error; err != nil {
		log.Printf("Service with ID %d not found", id)
		return
	}

	timeout := 10 // default timeout
	if service.Timeout > 0 {
		timeout = service.Timeout
	}

	client := http.Client{ 
		Timeout: time.Duration(timeout) * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(service.URL)  // Perform HTTP GET request statuscode : 2** or 4**
	latency := time.Since(start).Milliseconds()  // now() - start

	status := "DOWN"
	statusCode:= "Error"

	
	if err == nil {
		statusCode = resp.Status
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			status = "UP"
		}
		resp.Body.Close()
	}


    // check if status has changed
	if status != service.LastStatus {
		log.Printf(" State Changed for %s: %s -> %s", service.Name, service.LastStatus, status)
	
		// send websocket notification about status change
		hub.BroadcastUpdate(map[string]interface{}{
			"service_id": service.ID,
            "name":       service.Name,
            "new_status": status,
            "old_status": service.LastStatus,
            "latency":    latency,
		})


	
	}

	db.DB.Model(&service).Updates(map[string]interface{}{
		"last_status": status,
		"last_check":  time.Now(),
	})

	healthLog := models.HealthLog{
		ServiceID: service.ID,
		Status:    status,
		LatencyMs: int(latency),
		CheckedAt: time.Now(),
	}
	db.DB.Create(&healthLog)

log.Printf("Finished check for %s: Result [%s], Latency [%dms]", service.Name, status, latency )

log.Printf("Status Code: %s", statusCode)

}

