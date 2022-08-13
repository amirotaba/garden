package gRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.GardenRepository {
	return &mysqlRepository{
		Conn: Conn,
	}
}

func (m *mysqlRepository) CreateGarden(garden *domain.Garden) error {
	if err := m.Conn.Create(garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) ReadGarden(n int) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Limit(n).Find(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) ReadGardenID(id uint) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Where("id = ?", id).First(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) ReadGardenUID(id uint) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Where("user_id = ?", id).First(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) UpdateGarden(garden *domain.GardenForm) error {
	if err := m.Conn.Model(domain.Garden{}).Where("id = ?", garden.ID).Updates(garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) DeleteGarden(id uint) error {
	var garden domain.Garden
	if err := m.Conn.Where("id = ?", id).Delete(&garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) CreateLocation(location *domain.GardenLocation) error {
	if err := m.Conn.Create(location).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) ReadLocation(n int) ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Limit(n).Find(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlRepository) ReadLocationID(id uint) ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Where("garden_id = ?", id).First(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlRepository) UpdateLocation(loc *domain.GardenLocationForm) error {
	if err := m.Conn.Model(domain.GardenLocation{}).Where("id = ?", loc.ID).Updates(loc).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) DeleteLocation(id uint) error {
	var loc domain.GardenLocation
	if err := m.Conn.Where("id = ?", id).Delete(&loc).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) CreateGardenType(gardenType *domain.GardenType) error {
	if err := m.Conn.Create(gardenType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) ReadGardenType() ([]domain.GardenType, error) {
	var gType []domain.GardenType
	if err := m.Conn.Find(&gType).Error; err != nil {
		return []domain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlRepository) ReadGardenTypeID(id uint) ([]domain.GardenType, error) {
	var gType []domain.GardenType
	if err := m.Conn.Where("id = ?", id).First(&gType).Error; err != nil {
		return []domain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlRepository) UpdateGardenType(gardenType *domain.GardenTypeForm) error {
	if err := m.Conn.Model(domain.GardenType{}).Where("id = ?", gardenType.ID).Updates(gardenType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) DeleteGardenType(id uint) error {
	var gType domain.GardenType
	if err := m.Conn.Where("id = ?", id).Delete(&gType).Error; err != nil {
		return err
	}
	return nil
}
