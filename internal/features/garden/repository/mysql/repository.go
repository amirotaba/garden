package gardenRepo

import (
	"garden/internal/domain/garden"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) gardenDomain.GardenRepository {
	return &mysqlRepository{
		Conn: Conn,
	}
}

func (m *mysqlRepository) Create(garden *gardenDomain.Garden) error {
	if err := m.Conn.Create(garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Read(n int) ([]gardenDomain.Garden, error) {
	var garden []gardenDomain.Garden
	if err := m.Conn.Limit(n).Find(&garden).Error; err != nil {
		return []gardenDomain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) ReadID(id uint) ([]gardenDomain.Garden, error) {
	var garden []gardenDomain.Garden
	if err := m.Conn.Where("id = ?", id).First(&garden).Error; err != nil {
		return []gardenDomain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) ReadUID(id uint) ([]gardenDomain.Garden, error) {
	var garden []gardenDomain.Garden
	if err := m.Conn.Where("user_id = ?", id).First(&garden).Error; err != nil {
		return []gardenDomain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) Update(g *gardenDomain.GardenForm) error {
	if err := m.Conn.Model(gardenDomain.Garden{}).Where("id = ?", g.ID).Updates(g).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Delete(id uint) error {
	var garden gardenDomain.Garden
	if err := m.Conn.Where("id = ?", id).Delete(&garden).Error; err != nil {
		return err
	}
	return nil
}
