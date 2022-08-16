package gardenTypeRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.GardenTypeRepository {
	return &mysqlRepository{
		Conn: Conn,
	}
}

func (m *mysqlRepository) Create(gardenType *domain.GardenType) error {
	if err := m.Conn.Create(gardenType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Read() ([]domain.GardenType, error) {
	var gType []domain.GardenType
	if err := m.Conn.Find(&gType).Error; err != nil {
		return []domain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlRepository) ReadID(id uint) ([]domain.GardenType, error) {
	var gType []domain.GardenType
	if err := m.Conn.Where("id = ?", id).First(&gType).Error; err != nil {
		return []domain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlRepository) Update(gardenType *domain.GardenTypeForm) error {
	if err := m.Conn.Model(domain.GardenType{}).Where("id = ?", gardenType.ID).Updates(gardenType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Delete(id uint) error {
	var gType domain.GardenType
	if err := m.Conn.Where("id = ?", id).Delete(&gType).Error; err != nil {
		return err
	}
	return nil
}
