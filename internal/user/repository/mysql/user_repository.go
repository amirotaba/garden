package uRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func (m *mysqlUserRepository) Account(n int) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(n).Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) AccountUsername(username string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) AccountID(id uint) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) AccountType(n int, tp uint) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(n).Where("type = ?", tp).Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) SignUp(user *domain.User) error {
	if err := m.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) SignIn(form *domain.LoginForm) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", form.Username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) UpdateUser(user *domain.UserForm) error {
	if err := m.Conn.Model(domain.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteUser(id uint) error {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateUserType(usertype *domain.UserType) error {
	if err := m.Conn.Create(usertype).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadUserType() ([]domain.UserType, error) {
	var uType []domain.UserType
	if err := m.Conn.Find(&uType).Error; err != nil {
		return []domain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) ReadUserTypeID(id uint) ([]domain.UserType, error) {
	var uType []domain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return []domain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) UpdateUserType(userType *domain.UserTypeForm) error {
	if err := m.Conn.Model(domain.UserType{}).Where("id = ?", userType.ID).Updates(userType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteUserType(id uint) error {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) UserType(id uint) (string, error) {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return "", err
	}
	return uType.Name, nil
}
