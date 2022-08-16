package userDomain

import (
	"garden/internal/domain/userType"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Type     uint   `json:"type"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type UserForm struct {
	ID       uint   `json:"id"`
	Type     uint   `json:"type"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active" gorm:"default:True"`
}

type CheckAccessForm struct {
	AccessList string
	ServiceID  uint
}

type LoginForm struct {
	Type     uint   `json:"type"`
	Username string `json:"user_name"`
	Password string `json:"pass_word"`
}

type UserResponse struct {
	UserName string
	Type     userTypeDomain.TypeStruct
	Token    string
}

type AuthMessage struct {
	Text     string
	UserInfo UserResponse
}

type SignUpMessage struct {
	Text     string
	UserName string
	Email    string
}

type AccountForm struct {
	Uid        uint
	Tp         string
	PageNumber string
}

type UserAccountForm struct {
	Uid      string
	Username string
	ID       string
}

type UserReadForm struct {
	Span   int
	TypeID uint
}

type UserUseCase interface {
	Create(newUser User) (UserResponse, error)
	SignIn(form *LoginForm) (UserResponse, error)
	Read(form AccountForm) ([]UserResponse, error)
	UserRead(form UserAccountForm) (UserResponse, error)
	Update(user *UserForm) error
	Delete(user *User) error
}

type UserRepository interface {
	Create(newUser User) error
	Read(n int) ([]User, error)
	ReadID(id uint) (User, error)
	ReadByType(readForm UserReadForm) ([]User, error)
	ReadUsername(username string) (User, error)
	Update(user *UserForm) error
	Delete(id uint) error
}
