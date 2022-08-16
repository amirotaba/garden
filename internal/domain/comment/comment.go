package commentDomain

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text   string `json:"text"`
	TreeId uint   `json:"tree_id"`
	TagId  uint   `json:"tag_id"`
	Image  string `json:"image"`
	UserId uint   `json:"user_id"`
}

type CommentForm struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	TreeID uint   `json:"tree_id"`
	TagID  uint   `json:"tag_id"`
	Image  string `json:"image"`
	UserID uint   `json:"user_id"`
}

type ReadCommentForm struct {
	ID         string
	TreeID     string
	TagID      string
	UserID     string
	PageNumber string
	Uid        string
}

type UpdateCommentForm struct {
	Comment CommentForm
	Uid     uint
}

type ReadComment struct {
	ID    uint
	Query string
	Span  int
}

type CommentUseCase interface {
	Create(comment *Comment) error
	Read(form ReadCommentForm) ([]Comment, error)
	Update(form *UpdateCommentForm) error
	Delete(form *UpdateCommentForm) error
}

type CommentRepository interface {
	Create(comment *Comment) error
	Read(n int) ([]Comment, error)
	ReadID(readForm ReadComment) ([]Comment, error)
	Update(comment CommentForm) error
	Delete(id uint) error
}
