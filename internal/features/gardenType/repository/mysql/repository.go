package gardenTypeRepo

import (
	"garden/internal/domain/gardenType"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) gardenTypeDomain.GardenTypeRepository {
	return &mysqlRepository{
		Conn: Conn,
	}
}

func (m *mysqlRepository) Create(gardenType *gardenTypeDomain.GardenType) error {
	if err := m.Conn.Create(gardenType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Read() ([]gardenTypeDomain.GardenType, error) {
	var gType []gardenTypeDomain.GardenType
	if err := m.Conn.Find(&gType).Error; err != nil {
		return []gardenTypeDomain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlRepository) ReadID(id uint) ([]gardenTypeDomain.GardenType, error) {
	var gType []gardenTypeDomain.GardenType
	if err := m.Conn.Where("id = ?", id).First(&gType).Error; err != nil {
		return []gardenTypeDomain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlRepository) Update(Type *gardenTypeDomain.GardenTypeForm) error {
	if err := m.Conn.Model(gardenTypeDomain.GardenType{}).Where("id = ?", Type.ID).Updates(Type).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) Delete(id uint) error {
	var gType gardenTypeDomain.GardenType
	if err := m.Conn.Where("id = ?", id).Delete(&gType).Error; err != nil {
		return err
	}
	return nil
}
