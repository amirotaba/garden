package userTypeRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.UserTypeRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func (m *mysqlUserRepository) Create(usertype *domain.UserType) error {
	if err := m.Conn.Create(usertype).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Read() ([]domain.UserType, error) {
	var uType []domain.UserType
	if err := m.Conn.Find(&uType).Error; err != nil {
		return []domain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) ReadID(id uint) (domain.UserType, error) {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return domain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) Update(userType *domain.UserTypeForm) error {
	if err := m.Conn.Model(domain.UserType{}).Where("id = ?", userType.ID).Updates(userType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Delete(id uint) error {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}
