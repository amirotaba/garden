package treeRepo

import (
	"garden/internal/domain/tree"
	"garden/internal/domain/treeType"
	"gorm.io/gorm"
)

type mysqlTreeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) treeDomain.TreeRepository {
	return &mysqlTreeRepository{
		Conn: Conn,
	}
}

func (m *mysqlTreeRepository) Create(tree *treeDomain.Tree) error {
	if err := m.Conn.Create(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Read(n int) ([]treeDomain.Tree, error) {
	var tree []treeDomain.Tree
	if err := m.Conn.Limit(n).Find(&tree).Error; err != nil {
		return []treeDomain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlTreeRepository) ReadID(form treeDomain.ReadTreeID) ([]treeDomain.Tree, error) {
	var tType []treeDomain.Tree
	if err := m.Conn.Where(form.Query, form.ID).First(&tType).Error; err != nil {
		return []treeDomain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) ReadByType(form treeTypeDomain.ReadTreeType) ([]treeDomain.Tree, error) {
	var tType []treeDomain.Tree
	if err := m.Conn.Limit(form.Span).Where("type = ?", form.ID).First(&tType).Error; err != nil {
		return []treeDomain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlTreeRepository) Update(t *treeDomain.TreeForm) error {
	if err := m.Conn.Model(treeDomain.Tree{}).Where("id = ?", t.ID).Updates(t).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTreeRepository) Delete(id uint) error {
	var tree treeDomain.Tree
	if err := m.Conn.Where("id = ?", id).Delete(&tree).Error; err != nil {
		return err
	}
	return nil
}
