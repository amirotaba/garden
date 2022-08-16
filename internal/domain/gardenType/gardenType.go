package gardenTypeDomain

import "gorm.io/gorm"

type GardenType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GardenTypeForm struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GardenTypeUseCase interface {
	Create(gardenType *GardenType) error
	Read(id string) ([]GardenType, error)
	Update(gardenType *GardenTypeForm) error
	Delete(gardenType *GardenType) error
}

type GardenTypeRepository interface {
	Create(gardenType *GardenType) error
	Read() ([]GardenType, error)
	ReadID(u uint) ([]GardenType, error)
	Update(gardenType *GardenTypeForm) error
	Delete(id uint) error
}
