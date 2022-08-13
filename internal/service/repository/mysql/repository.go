package sRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlSerivceRepository struct {
	Conn *gorm.DB
}

func NewMysqlSerivceRepository(Conn *gorm.DB) domain.ServiceRepository {
	return &mysqlSerivceRepository{
		Conn: Conn,
	}
}

func (m *mysqlSerivceRepository) CreateService(service *domain.Service) error {
	if err := m.Conn.Create(service).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlSerivceRepository) ReadService() ([]domain.Service, error) {
	var service []domain.Service
	if err := m.Conn.Find(&service).Error; err != nil {
		return []domain.Service{}, err
	}
	return service, nil
}

func (m *mysqlSerivceRepository) ReadServiceUrl(url string) (domain.Service, error) {
	var service domain.Service
	if err := m.Conn.Where("url = ?", url).First(&service).Error; err != nil {
		return domain.Service{}, err
	}
	return service, nil
}

func (m *mysqlSerivceRepository) UpdateService(service *domain.ServiceForm) error {
	if err := m.Conn.Model(domain.Service{}).Where("id = ?", service.ID).Updates(service).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlSerivceRepository) DeleteService(id uint) error {
	var uType domain.Service
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}
