package userTypeDomain

import (
	"gorm.io/gorm"
)

type UserType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	AccessList  string `json:"access_list"`
}

type UserTypeForm struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AccessList  string `json:"access_list"`
}

type TypeStruct struct {
	ID   uint
	Name string
}

type AccessForm struct {
	ID     uint   `json:"id"`
	TypeID string `json:"type_id"`
}

type UserTypeUseCase interface {
	Create(usertype *UserType) error
	Read(id string) ([]UserType, error)
	UpdateAccess(access *AccessForm) error
	Update(usertype *UserTypeForm) error
	Delete(usertype *UserType) error
}
type UserTypeRepository interface {
	Create(usertype *UserType) error
	Read() ([]UserType, error)
	ReadID(id uint) (UserType, error)
	Update(userType *UserTypeForm) error
	Delete(id uint) error
}
