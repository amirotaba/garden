package userRepo

import (
	"garden/internal/domain/user"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) userDomain.UserRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func (m *mysqlUserRepository) Create(user userDomain.User) error {
	if err := m.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Read(n int) ([]userDomain.User, error) {
	var user []userDomain.User
	if err := m.Conn.Limit(n).Find(&user).Error; err != nil {
		return []userDomain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) ReadUsername(username string) (userDomain.User, error) {
	var user userDomain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return userDomain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) ReadID(id uint) (userDomain.User, error) {
	var user userDomain.User
	if err := m.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return userDomain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) ReadByType(readForm userDomain.UserReadForm) ([]userDomain.User, error) {
	var user []userDomain.User
	if err := m.Conn.Limit(readForm.Span).Where("type = ?", readForm.TypeID).Find(&user).Error; err != nil {
		return []userDomain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) Update(u *userDomain.UserForm) error {
	if err := m.Conn.Model(userDomain.User{}).Where("id = ?", u.ID).Updates(u).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Delete(id uint) error {
	var user userDomain.User
	if err := m.Conn.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
