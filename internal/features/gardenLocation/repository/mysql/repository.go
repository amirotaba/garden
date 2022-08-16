package gardenLocRepo

import (
	"garden/internal/domain/gardenLocation"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) gardenLocationDomain.GardenLocRepository {
	return &mysqlRepository{
		Conn: Conn,
	}
}

func (m *mysqlRepository) Create(location *gardenLocationDomain.GardenLocation) error {
	if err := m.Conn.Create(location).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Read(n int) ([]gardenLocationDomain.GardenLocation, error) {
	var loc []gardenLocationDomain.GardenLocation
	if err := m.Conn.Limit(n).Find(&loc).Error; err != nil {
		return []gardenLocationDomain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlRepository) ReadID(id uint) ([]gardenLocationDomain.GardenLocation, error) {
	var loc []gardenLocationDomain.GardenLocation
	if err := m.Conn.Where("garden_id = ?", id).First(&loc).Error; err != nil {
		return []gardenLocationDomain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlRepository) Update(loc *gardenLocationDomain.GardenLocationForm) error {
	if err := m.Conn.Model(gardenLocationDomain.GardenLocation{}).Where("id = ?", loc.ID).Updates(loc).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Delete(id uint) error {
	var loc gardenLocationDomain.GardenLocation
	if err := m.Conn.Where("id = ?", id).Delete(&loc).Error; err != nil {
		return err
	}
	return nil
}
