package commentUsecase

import (
	"garden/internal/domain/comment"
	"strconv"
)

type usecase struct {
	CommentRepo commentDomain.CommentRepository
}

func NewUseCase(r commentDomain.CommentRepository) commentDomain.CommentUseCase {
	return &usecase{
		CommentRepo: r,
	}
}

func (a *usecase) Create(comment *commentDomain.Comment) error {
	if err := a.CommentRepo.Create(comment); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(form commentDomain.ReadCommentForm) ([]commentDomain.Comment, error) {
	var readForm commentDomain.ReadComment
	if form.ID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return nil, err
		}
		idInt, err := strconv.Atoi(form.ID)
		if err != nil {
			return nil, err
		}

		readForm = commentDomain.ReadComment{
			ID:    uint(idInt),
			Query: "id = ?",
			Span:  nInt * 10,
		}
		t, err := a.CommentRepo.ReadID(readForm)
		if err != nil {
			return nil, err
		}
		return t, nil
	} else if form.TreeID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return nil, err
		}
		idInt, err := strconv.Atoi(form.TreeID)
		if err != nil {
			return nil, err
		}

		readForm = commentDomain.ReadComment{
			ID:    uint(idInt),
			Query: "tree_id = ?",
			Span:  nInt * 10,
		}

		t, err := a.CommentRepo.ReadID(readForm)
		if err != nil {
			return nil, err
		}
		return t, nil
	} else if form.TagID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return nil, err
		}
		idInt, err := strconv.Atoi(form.ID)
		if err != nil {
			return nil, err
		}

		readForm = commentDomain.ReadComment{
			ID:    uint(idInt),
			Query: "tag_id = ?",
			Span:  nInt * 10,
		}

		t, err := a.CommentRepo.ReadID(readForm)
		if err != nil {
			return nil, err
		}
		return t, nil
	} else if form.UserID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return nil, err
		}
		idInt, err := strconv.Atoi(form.UserID)
		if err != nil {
			return nil, err
		}

		readForm = commentDomain.ReadComment{
			ID:    uint(idInt),
			Query: "user_id = ?",
			Span:  nInt * 10,
		}

		t, err := a.CommentRepo.ReadID(readForm)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return nil, err
	}
	span := nInt * 10
	c, err := a.CommentRepo.Read(span)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (a *usecase) Update(form *commentDomain.UpdateCommentForm) error {
	readForm := commentDomain.ReadComment{
		ID:    form.Comment.ID,
		Query: "id = ?",
		Span:  1,
	}

	c, err := a.CommentRepo.ReadID(readForm)
	if err != nil {
		return err
	}

	if int(c[0].ID) == int(form.Uid) {
		if err := a.CommentRepo.Update(form.Comment); err != nil {
			return err
		}
	}
	return nil
}

func (a *usecase) Delete(form *commentDomain.UpdateCommentForm) error {
	readForm := commentDomain.ReadComment{
		ID:    form.Comment.ID,
		Query: "id = ?",
		Span:  1,
	}

	c, err := a.CommentRepo.ReadID(readForm)
	if err != nil {
		return err
	}
	if int(c[0].ID) == int(form.Uid) {
		if err := a.CommentRepo.Delete(form.Comment.ID); err != nil {
			return err
		}
	}
	return nil
}
