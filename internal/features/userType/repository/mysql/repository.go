package userTypeRepo

import (
	"garden/internal/domain/userType"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) userTypeDomain.UserTypeRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func (m *mysqlUserRepository) Create(usertype *userTypeDomain.UserType) error {
	if err := m.Conn.Create(usertype).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Read() ([]userTypeDomain.UserType, error) {
	var uType []userTypeDomain.UserType
	if err := m.Conn.Find(&uType).Error; err != nil {
		return []userTypeDomain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) ReadID(id uint) (userTypeDomain.UserType, error) {
	var uType userTypeDomain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return userTypeDomain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) Update(Type *userTypeDomain.UserTypeForm) error {
	if err := m.Conn.Model(userTypeDomain.UserType{}).Where("id = ?", Type.ID).Updates(Type).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Delete(id uint) error {
	var uType userTypeDomain.UserType
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}
