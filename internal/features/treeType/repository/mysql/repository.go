package treeTypeRepo

import (
	"garden/internal/domain/treeType"
	"gorm.io/gorm"
)

type mysqlTreeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) treeTypeDomain.TreeTypeRepository {
	return &mysqlTreeRepository{
		Conn: Conn,
	}
}

func (m *mysqlTreeRepository) Create(treeType *treeTypeDomain.TreeType) error {
	if err := m.Conn.Create(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Read() ([]treeTypeDomain.TreeType, error) {
	var tType []treeTypeDomain.TreeType
	if err := m.Conn.Find(&tType).Error; err != nil {
		return []treeTypeDomain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadID(id uint) ([]treeTypeDomain.TreeType, error) {
	var tType []treeTypeDomain.TreeType
	if err := m.Conn.Where("id = ?", id).First(&tType).Error; err != nil {
		return []treeTypeDomain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) Update(Type *treeTypeDomain.TreeTypeForm) error {
	if err := m.Conn.Model(treeTypeDomain.TreeType{}).Where("id = ?", Type.ID).Updates(Type).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Delete(id uint) error {
	var tType treeTypeDomain.TreeType
	if err := m.Conn.Where("id = ?", id).Delete(&tType).Error; err != nil {
		return err
	}
	return nil
}
