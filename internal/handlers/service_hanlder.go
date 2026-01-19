package handlers

import (

	"log"
	"github.com/gin-gonic/gin"

	"net/http"
    "distributed-health-monitor/internal/db"
	"distributed-health-monitor/internal/dtos"
	"distributed-health-monitor/internal/models"

)


// Register service   

func RegisterService(c *gin.Context){

	var servicedto dtos.ServiceCreateDTO

	if err := c.ShouldBindJSON(&servicedto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // return { "error": "error.message" }
		log.Println("Error binding JSON:", err)
		return
	}
	// mapping 
	service := models.Service{
		Name: servicedto.Name,
		URL:  servicedto.URL,
		Interval: servicedto.Interval,
		Timeout: servicedto.Timeout, 

	}

	if err := db.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
	    log.Println("Error creating service:", err)
		return
	}

	c.JSON(http.StatusCreated,service)
}

// List all services
func ListServices (c *gin.Context) {


	var services []models.Service

	if err := db.DB.Find(&services).Error;  err != nil {
		c.JSON (http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		log.Println("Error retrieving services:", err)
		return
	}

	var response []dtos.ServiceResponseDTO
	for _, s := range services {
		response = append(response, dtos.ServiceResponseDTO{
			ID:         s.ID,
			Name:       s.Name,
			URL:        s.URL,
			Interval:   s.Interval,
			LastStatus: s.LastStatus,
			CreatedAt: s.CreatedAt,
            LastState:  s.LastCheck,
			Timeout: s.Timeout,

		})
	}


	c.JSON (http.StatusOK, response)



}