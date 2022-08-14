package tagRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlTagRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.TagRepository {
	return &mysqlTagRepository{
		Conn: Conn,
	}
}

func (m *mysqlTagRepository) Create(tag *domain.Tag) error {
	if err := m.Conn.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTagRepository) Read(n int) ([]domain.Tag, error) {
	var tag []domain.Tag
	if err := m.Conn.Limit(n).Find(&tag).Error; err != nil {
		return []domain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlTagRepository) ReadID(id uint) ([]domain.Tag, error) {
	var tag []domain.Tag
	if err := m.Conn.Where("id = ?", id).First(&tag).Error; err != nil {
		return []domain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlTagRepository) Update(tag *domain.TagForm) error {
	if err := m.Conn.Model(domain.Tag{}).Where("id = ?", tag.ID).Updates(tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlTagRepository) Delete(id uint) error {
	var tag domain.Tag
	if err := m.Conn.Where("id = ?", id).Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}
