package treeRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlTreeRepository struct {
	Conn *gorm.DB
}

func NewMysqlTreeRepository(Conn *gorm.DB) domain.TreeRepository {
	return &mysqlTreeRepository{
		Conn: Conn,
	}
}

func (m *mysqlTreeRepository) Create(tree *domain.Tree) error {
	if err := m.Conn.Create(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Read(n int) ([]domain.Tree, error) {
	var tree []domain.Tree
	if err := m.Conn.Limit(n).Find(&tree).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlTreeRepository) ReadID(id uint, q string) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Where(q, id).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadByType(t uint, n int) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Limit(n).Where("type = ?", t).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) Update(tree *domain.TreeForm) error {
	if err := m.Conn.Model(domain.Tree{}).Where("id = ?", tree.ID).Updates(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Delete(id uint) error {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).Delete(&tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) CreateType(treeType *domain.TreeType) error {
	if err := m.Conn.Create(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) ReadType() ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Find(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadTypeID(id uint) ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Where("id = ?", id).First(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) UpdateType(treeType *domain.TreeTypeForm) error {
	if err := m.Conn.Model(domain.TreeType{}).Where("id = ?", treeType.ID).Updates(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) DeleteType(id uint) error {
	var tType domain.TreeType
	if err := m.Conn.Where("id = ?", id).Delete(&tType).Error; err != nil {
		return err
	}
	return nil
}
