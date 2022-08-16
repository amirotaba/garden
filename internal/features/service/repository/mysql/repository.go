package serviceRepo

import (
	"garden/internal/domain/service"
	"gorm.io/gorm"
)

type mysqlSerivceRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) serviceDomain.ServiceRepository {
	return &mysqlSerivceRepository{
		Conn: Conn,
	}
}

func (m *mysqlSerivceRepository) Create(service *serviceDomain.Service) error {
	if err := m.Conn.Create(service).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlSerivceRepository) Read() ([]serviceDomain.Service, error) {
	var service []serviceDomain.Service
	if err := m.Conn.Find(&service).Error; err != nil {
		return []serviceDomain.Service{}, err
	}
	return service, nil
}

func (m *mysqlSerivceRepository) ReadURL(url string) (serviceDomain.Service, error) {
	var service serviceDomain.Service
	if err := m.Conn.Where("url = ?", url).First(&service).Error; err != nil {
		return serviceDomain.Service{}, err
	}
	return service, nil
}

func (m *mysqlSerivceRepository) Update(s *serviceDomain.ServiceForm) error {
	if err := m.Conn.Model(serviceDomain.Service{}).Where("id = ?", s.ID).Updates(s).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlSerivceRepository) Delete(id uint) error {
	var uType serviceDomain.Service
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}
