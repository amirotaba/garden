package tagDomain

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	History        string `json:"history"`
	Detoxification string `json:"detoxification"`
	Image          string `json:"image"`
}

type TagForm struct {
	ID             uint   `json:"id"`
	History        string `json:"history"`
	Detoxification string `json:"detoxification"`
	Image          string `json:"image"`
}

type TagUseCase interface {
	Create(tag *Tag) error
	Read(pageNumber string) ([]Tag, error)
	ReadID(id string) ([]Tag, error)
	Update(tag *TagForm) error
	Delete(tag *Tag) error
}

type TagRepository interface {
	Create(tag *Tag) error
	Read(n int) ([]Tag, error)
	ReadID(u uint) ([]Tag, error)
	Update(tag *TagForm) error
	Delete(id uint) error
}
