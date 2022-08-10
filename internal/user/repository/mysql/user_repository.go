package mysql

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type repository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &repository{
		Conn: Conn,
	}
}

func (m *repository) Account(n int) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(n).Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *repository) AccountUsername(username string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *repository) AccountID(id uint) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *repository) AccountType(n int, tp uint) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(n).Where("type = ?", tp).Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *repository) SignUp(user *domain.User) error {
	if err := m.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *repository) SignIn(form *domain.LoginForm) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", form.Username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *repository) UpdateUser(user *domain.UserForm) error {
	if err := m.Conn.Model(domain.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *repository) DeleteUser(id uint) error {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (m *repository) CreateUserType(usertype *domain.UserType) error {
	if err := m.Conn.Create(usertype).Error; err != nil {
		return err
	}
	return nil
}

func (m *repository) ReadUserType() ([]domain.UserType, error) {
	var uType []domain.UserType
	if err := m.Conn.Find(&uType).Error; err != nil {
		return []domain.UserType{}, err
	}
	return uType, nil
}

func (m *repository) ReadUserTypeID(id uint) ([]domain.UserType, error) {
	var uType []domain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return []domain.UserType{}, err
	}
	return uType, nil
}

func (m *repository) UpdateUserType(userType *domain.UserTypeForm) error {
	if err := m.Conn.Model(domain.UserType{}).Where("id = ?", userType.ID).Updates(userType).Error; err != nil {
		return err
	}
	return nil
}

func (m *repository) DeleteUserType(id uint) error {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}

func (m *repository) UserType(id uint) (string, error) {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return "", err
	}
	return uType.Name, nil
}
