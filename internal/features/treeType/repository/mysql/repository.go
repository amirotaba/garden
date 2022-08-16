package treeTypeRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlTreeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.TreeTypeRepository {
	return &mysqlTreeRepository{
		Conn: Conn,
	}
}

func (m *mysqlTreeRepository) Create(treeType *domain.TreeType) error {
	if err := m.Conn.Create(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Read() ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Find(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadID(id uint) ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Where("id = ?", id).First(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) Update(treeType *domain.TreeTypeForm) error {
	if err := m.Conn.Model(domain.TreeType{}).Where("id = ?", treeType.ID).Updates(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Delete(id uint) error {
	var tType domain.TreeType
	if err := m.Conn.Where("id = ?", id).Delete(&tType).Error; err != nil {
		return err
	}
	return nil
}
