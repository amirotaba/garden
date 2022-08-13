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

func (m *mysqlTreeRepository) CreateTree(tree *domain.Tree) error {
	if err := m.Conn.Create(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) ReadTree(n int) ([]domain.Tree, error) {
	var tree []domain.Tree
	if err := m.Conn.Limit(n).Find(&tree).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlTreeRepository) ReadTreeID(id uint, q string) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Where(q, id).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadTreeByType(t uint, n int) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Limit(n).Where("type = ?", t).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) UpdateTree(tree *domain.TreeForm) error {
	if err := m.Conn.Model(domain.Tree{}).Where("id = ?", tree.ID).Updates(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) DeleteTree(id uint) error {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).Delete(&tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) CreateTreeType(treeType *domain.TreeType) error {
	if err := m.Conn.Create(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) ReadTreeType() ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Find(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadTreeTypeID(id uint) ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Where("id = ?", id).First(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) UpdateTreeType(treeType *domain.TreeTypeForm) error {
	if err := m.Conn.Model(domain.TreeType{}).Where("id = ?", treeType.ID).Updates(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) DeleteTreeType(id uint) error {
	var tType domain.TreeType
	if err := m.Conn.Where("id = ?", id).Delete(&tType).Error; err != nil {
		return err
	}
	return nil
}
