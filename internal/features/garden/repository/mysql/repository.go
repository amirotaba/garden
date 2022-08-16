package gardenRepo

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

func (m *mysqlRepository) Create(garden *domain.Garden) error {
	if err := m.Conn.Create(garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Read(n int) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Limit(n).Find(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) ReadID(id uint) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Where("id = ?", id).First(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) ReadUID(id uint) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Where("user_id = ?", id).First(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlRepository) Update(garden *domain.GardenForm) error {
	if err := m.Conn.Model(domain.Garden{}).Where("id = ?", garden.ID).Updates(garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Delete(id uint) error {
	var garden domain.Garden
	if err := m.Conn.Where("id = ?", id).Delete(&garden).Error; err != nil {
		return err
	}
	return nil
}
