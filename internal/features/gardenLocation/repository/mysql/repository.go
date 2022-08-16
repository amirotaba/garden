package gardenLocRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.GardenLocRepository {
	return &mysqlRepository{
		Conn: Conn,
	}
}

func (m *mysqlRepository) Create(location *domain.GardenLocation) error {
	if err := m.Conn.Create(location).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Read(n int) ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Limit(n).Find(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlRepository) ReadID(id uint) ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Where("garden_id = ?", id).First(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlRepository) Update(loc *domain.GardenLocationForm) error {
	if err := m.Conn.Model(domain.GardenLocation{}).Where("id = ?", loc.ID).Updates(loc).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Delete(id uint) error {
	var loc domain.GardenLocation
	if err := m.Conn.Where("id = ?", id).Delete(&loc).Error; err != nil {
		return err
	}
	return nil
}
