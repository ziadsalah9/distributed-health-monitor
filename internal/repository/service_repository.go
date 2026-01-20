
package repository

import (


	"distributed-health-monitor/internal/models"
    "gorm.io/gorm"

)

type ServiceRepository interface{
	CreateService(service *models.Service) error
	GetAllServices() ([]models.Service, error)
    GetLogsByServiceID(id uint) ([]models.HealthLog, error)
}

type ServiceRepositoryimpl struct{
	db *gorm.DB
}
func NewServiceRepository(db *gorm.DB) ServiceRepository {
    return &ServiceRepositoryimpl{db: db}
}

func (r *ServiceRepositoryimpl) CreateService(service *models.Service) error {
    return r.db.Create(service).Error
}

func (r *ServiceRepositoryimpl) GetAllServices() ([]models.Service, error) {
    var services []models.Service
    err := r.db.Find(&services).Error
    return services, err
}

func (r *ServiceRepositoryimpl) GetLogsByServiceID(id uint) ([]models.HealthLog, error) {
    var logs []models.HealthLog
    err := r.db.Where("service_id = ?", id).Order("checked_at desc").Find(&logs).Error
    return logs, err
}