package mysql

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlTagRepository struct {
	Conn *gorm.DB
}

func NewMysqlTagRepository(Conn *gorm.DB) domain.TagRepository {
	return &mysqlTagRepository{
		Conn: Conn,
	}
}

func (m *mysqlTagRepository) CreateTag(tag *domain.Tag) error {
	if err := m.Conn.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTagRepository) ReadTag(n int) ([]domain.Tag, error) {
	var tag []domain.Tag
	if err := m.Conn.Limit(n).Find(&tag).Error; err != nil {
		return []domain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlTagRepository) ReadTagID(id uint) ([]domain.Tag, error) {
	var tag []domain.Tag
	if err := m.Conn.Where("id = ?", id).First(&tag).Error; err != nil {
		return []domain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlTagRepository) UpdateTag(tag *domain.TagForm) error {
	if err := m.Conn.Model(domain.Tag{}).Where("id = ?", tag.ID).Updates(tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTagRepository) DeleteTag(id uint) error {
	var tag domain.Tag
	if err := m.Conn.Where("id = ?", id).Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}
