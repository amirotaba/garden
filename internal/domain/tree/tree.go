package treeDomain

import (
	"garden/internal/domain/treeType"
	"gorm.io/gorm"
	"time"
)

type Tree struct {
	gorm.Model
	FullName    string    `json:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Type        uint      `json:"type"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Qr          string
	Length      float64 `json:"length"`
	Image       string  `json:"image"`
	GardenId    uint    `json:"garden_id"`
	Description string  `json:"description"`
}

type TreeForm struct {
	ID          uint      `json:"id"`
	FullName    string    `json:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Type        uint      `json:"type"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Qr          string
	Length      float64 `json:"length"`
	Image       string  `json:"image"`
	GardenId    uint    `json:"garden_id"`
	Description string  `json:"description"`
}

type ReadTreeForm struct {
	Uid        string
	GardenID   string
	Tp         string
	PageNumber string
}

type ReadTreeUserForm struct {
	ID       string
	GardenID string
}

type ReadTreeID struct {
	Query string
	ID    uint
}

type TreeUseCase interface {
	Create(tree *Tree) error
	Read(form ReadTreeForm) ([]Tree, error)
	ReadUser(form ReadTreeUserForm) ([]Tree, error)
	Update(tree *TreeForm) error
	Delete(tree *Tree) error
}

type TreeRepository interface {
	Create(tree *Tree) error
	Read(n int) ([]Tree, error)
	ReadID(readForm ReadTreeID) ([]Tree, error)
	ReadByType(readForm treeTypeDomain.ReadTreeType) ([]Tree, error)
	Update(tree *TreeForm) error
	Delete(id uint) error
}
