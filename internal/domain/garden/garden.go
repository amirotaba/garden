package gardenDomain

import "gorm.io/gorm"

type Garden struct {
	gorm.Model
	Name        string  `json:"name"`
	Type        uint    `json:"type"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	UserID      uint    `json:"user_id"`
	Description string  `json:"description"`
}

type GardenForm struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	UserID      uint    `json:"user_id"`
	Description string  `json:"description"`
}

type ReadGardenForm struct {
	Uid        string
	UserID     string
	PageNumber string
	ID         string
}

type GardenUseCase interface {
	Create(garden *Garden) error
	Read(form ReadGardenForm) ([]Garden, error)
	Update(garden *GardenForm) error
	Delete(garden *Garden) error
}

type GardenRepository interface {
	Create(garden *Garden) error
	Read(n int) ([]Garden, error)
	ReadID(u uint) ([]Garden, error)
	ReadUID(id uint) ([]Garden, error)
	Update(garden *GardenForm) error
	Delete(id uint) error
}
