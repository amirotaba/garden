package userRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func (m *mysqlUserRepository) Create(user domain.User) error {
	if err := m.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Read(n int) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(n).Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) ReadUsername(username string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) ReadID(id uint) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) ReadByType(readForm domain.UserReadForm) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(readForm.Span).Where("type = ?", readForm.TypeID).Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) Update(user *domain.UserForm) error {
	if err := m.Conn.Model(domain.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Delete(id uint) error {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
