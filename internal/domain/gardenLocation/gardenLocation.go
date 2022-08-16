package gardenLocationDomain

import "gorm.io/gorm"

type GardenLocation struct {
	gorm.Model
	Lat1     float64 `json:"lat_1"`
	Lat2     float64 `json:"lat_2"`
	Lat3     float64 `json:"lat_3"`
	Lat4     float64 `json:"lat_4"`
	Long1    float64 `json:"long_1"`
	Long2    float64 `json:"long_2"`
	Long3    float64 `json:"long_3"`
	Long4    float64 `json:"long_4"`
	GardenId uint    `json:"garden_id"`
	UserId   uint    `json:"user_id"`
}

type GardenLocationForm struct {
	ID       uint    `json:"id"`
	Lat1     float64 `json:"lat_1"`
	Lat2     float64 `json:"lat_2"`
	Lat3     float64 `json:"lat_3"`
	Lat4     float64 `json:"lat_4"`
	Long1    float64 `json:"long_1"`
	Long2    float64 `json:"long_2"`
	Long3    float64 `json:"long_3"`
	Long4    float64 `json:"long_4"`
	GardenId uint    `json:"garden_id"`
	UserId   uint    `json:"user_id"`
}

type GardenLocRead struct {
	GardenID   string
	PageNumber string
}

type GardenLocUseCase interface {
	Create(location *GardenLocation) error
	Read(form GardenLocRead) ([]GardenLocation, error)
	Update(loc *GardenLocationForm) error
	Delete(loc *GardenLocation) error
}

type GardenLocRepository interface {
	Create(location *GardenLocation) error
	Read(n int) ([]GardenLocation, error)
	ReadID(u uint) ([]GardenLocation, error)
	Update(loc *GardenLocationForm) error
	Delete(id uint) error
}
