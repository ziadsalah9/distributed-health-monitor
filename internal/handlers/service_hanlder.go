package handlers

import (

	"log"
	"github.com/gin-gonic/gin"
	"distributed-health-monitor/internal/repository"
	"net/http"
	"distributed-health-monitor/internal/dtos"
	"distributed-health-monitor/internal/models"
	"strconv"	

)


type ServiceHandler struct {
    repo repository.ServiceRepository
}

// injection
func NewServiceHandler(r repository.ServiceRepository) *ServiceHandler {
	return &ServiceHandler{repo: r}
}



// Register service   

func (h *ServiceHandler) RegisterService(c *gin.Context){

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

	if err := h.repo.CreateService(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}

	c.JSON(http.StatusCreated,service)
}

// List all services
func (h *ServiceHandler) ListServices(c *gin.Context) {


	services, err := h.repo.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		return
	}
	// var services []models.Service

	// if err := db.DB.Find(&services).Error;  err != nil {
	// 	c.JSON (http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
	// 	log.Println("Error retrieving services:", err)
	// 	return
	// }

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


func (h *ServiceHandler) GetServiceLogs(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	logs, err := h.repo.GetLogsByServiceID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch logs"})
		return
	}

	c.JSON(http.StatusOK, logs)
}