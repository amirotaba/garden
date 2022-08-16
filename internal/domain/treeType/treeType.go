package treeTypeDomain

import "gorm.io/gorm"

type TreeType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TreeTypeForm struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ReadTreeType struct {
	ID   uint
	Span int
}

type TreeTypeUseCase interface {
	Create(treeType *TreeType) error
	Read(id string) ([]TreeType, error)
	Update(treeType *TreeTypeForm) error
	Delete(treeType *TreeType) error
}

type TreeTypeRepository interface {
	Create(treeType *TreeType) error
	Read() ([]TreeType, error)
	ReadID(u uint) ([]TreeType, error)
	Update(treeType *TreeTypeForm) error
	Delete(id uint) error
}
