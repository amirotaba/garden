package commentRepo

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) domain.CommentRepository {
	return &mysqlCommentRepository{
		Conn: Conn,
	}
}

func (m *mysqlCommentRepository) Create(comment *domain.Comment) error {
	if err := m.Conn.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlCommentRepository) Read(n int) ([]domain.Comment, error) {
	var comment []domain.Comment
	if err := m.Conn.Limit(n).Find(&comment).Error; err != nil {
		return []domain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlCommentRepository) ReadID(id uint, q string, n int) ([]domain.Comment, error) {
	var comment []domain.Comment
	if err := m.Conn.Limit(n).Where(q, id).First(&comment).Error; err != nil {
		return []domain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlCommentRepository) Update(comment *domain.CommentForm) error {
	if err := m.Conn.Model(domain.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlCommentRepository) Delete(id uint) error {
	var comment domain.Comment
	if err := m.Conn.Where("id = ?", id).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
