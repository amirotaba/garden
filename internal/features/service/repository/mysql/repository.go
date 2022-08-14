package serviceRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlSerivceRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.ServiceRepository {
	return &mysqlSerivceRepository{
		Conn: Conn,
	}
}

func (m *mysqlSerivceRepository) Create(service *domain.Service) error {
	if err := m.Conn.Create(service).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlSerivceRepository) Read() ([]domain.Service, error) {
	var service []domain.Service
	if err := m.Conn.Find(&service).Error; err != nil {
		return []domain.Service{}, err
	}
	return service, nil
}

func (m *mysqlSerivceRepository) ReadURL(url string) (domain.Service, error) {
	var service domain.Service
	if err := m.Conn.Where("url = ?", url).First(&service).Error; err != nil {
		return domain.Service{}, err
	}
	return service, nil
}

func (m *mysqlSerivceRepository) Update(service *domain.ServiceForm) error {
	if err := m.Conn.Model(domain.Service{}).Where("id = ?", service.ID).Updates(service).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlSerivceRepository) Delete(id uint) error {
	var uType domain.Service
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}
