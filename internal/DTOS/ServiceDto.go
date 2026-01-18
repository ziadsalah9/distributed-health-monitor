package dtos

type ServiceCreateDTO struct {
	Name     string `json:"name" binding:"required,min=3"`
	URL      string `json:"url" binding:"required,url"`
	Interval int    `json:"interval" binding:"required,min=10"`
}

type ServiceResponseDTO struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Interval   int    `json:"interval"`
	LastStatus string `json:"last_status"`
	LastState  string `json:"last_state"`
}