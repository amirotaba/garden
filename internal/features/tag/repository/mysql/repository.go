package tagRepo

import (
	"garden/internal/domain/tag"
	"gorm.io/gorm"
)

type mysqlTagRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) tagDomain.TagRepository {
	return &mysqlTagRepository{
		Conn: Conn,
	}
}

func (m *mysqlTagRepository) Create(tag *tagDomain.Tag) error {
	if err := m.Conn.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTagRepository) Read(n int) ([]tagDomain.Tag, error) {
	var tag []tagDomain.Tag
	if err := m.Conn.Limit(n).Find(&tag).Error; err != nil {
		return []tagDomain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlTagRepository) ReadID(id uint) ([]tagDomain.Tag, error) {
	var tag []tagDomain.Tag
	if err := m.Conn.Where("id = ?", id).First(&tag).Error; err != nil {
		return []tagDomain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlTagRepository) Update(t *tagDomain.TagForm) error {
	if err := m.Conn.Model(tagDomain.Tag{}).Where("id = ?", t.ID).Updates(t).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTagRepository) Delete(id uint) error {
	var tag tagDomain.Tag
	if err := m.Conn.Where("id = ?", id).Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}
