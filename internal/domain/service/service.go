package serviceDomain

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name   string `json:"name"`
	Url    string `json:"url"`
	Method string `json:"method"`
}

type ServiceForm struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ServiceUseCase interface {
	Create(service *Service) error
	Read() ([]Service, error)
	Update(usertype *ServiceForm) error
	Delete(service *Service) error
}

type ServiceRepository interface {
	Create(service *Service) error
	Read() ([]Service, error)
	ReadURL(url string) (Service, error)
	Update(service *ServiceForm) error
	Delete(id uint) error
}
