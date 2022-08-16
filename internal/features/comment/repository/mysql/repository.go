package commentRepo

import (
	"garden/internal/domain/comment"
	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(Conn *gorm.DB) commentDomain.CommentRepository {
	return &mysqlCommentRepository{
		Conn: Conn,
	}
}

func (m *mysqlCommentRepository) Create(comment *commentDomain.Comment) error {
	if err := m.Conn.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlCommentRepository) Read(n int) ([]commentDomain.Comment, error) {
	var comment []commentDomain.Comment
	if err := m.Conn.Limit(n).Find(&comment).Error; err != nil {
		return []commentDomain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlCommentRepository) ReadID(form commentDomain.ReadComment) ([]commentDomain.Comment, error) {
	var comment []commentDomain.Comment
	if err := m.Conn.Limit(form.Span).Where(form.Query, form.ID).First(&comment).Error; err != nil {
		return []commentDomain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlCommentRepository) Update(c commentDomain.CommentForm) error {
	if err := m.Conn.Model(commentDomain.Comment{}).Where("id = ?", c.ID).Updates(c).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlCommentRepository) Delete(id uint) error {
	var comment commentDomain.Comment
	if err := m.Conn.Where("id = ?", id).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
