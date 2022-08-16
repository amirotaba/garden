package treeRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlTreeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.TreeRepository {
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

func (m *mysqlTreeRepository) ReadID(form domain.ReadTreeID) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Where(form.Query, form.ID).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadByType(form domain.ReadTreeType) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Limit(form.Span).Where("type = ?", form.ID).First(&tType).Error; err != nil {
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
