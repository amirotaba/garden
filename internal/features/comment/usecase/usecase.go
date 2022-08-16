package commentUsecase

import (
	"strconv"

	"garden/internal/domain"
)

type usecase struct {
	CommentRepo domain.CommentRepository
}

func NewUseCase(r domain.CommentRepository) domain.CommentUseCase {
	return &usecase{
		CommentRepo: r,
	}
}

func (a *usecase) Create(comment *domain.Comment) error {
	if err := a.CommentRepo.Create(comment); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(form domain.ReadCommentForm) ([]domain.Comment, error) {
	if form.ID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.ID)
		if err != nil {
			return []domain.Comment{}, err
		}
		q := "id = ?"
		t, err := a.CommentRepo.ReadID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, err
		}
		return t, nil
	} else if form.TreeID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.TreeID)
		if err != nil {
			return []domain.Comment{}, err
		}
		q := "tree_id = ?"
		t, err := a.CommentRepo.ReadID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, err
		}
		return t, nil
	} else if form.TagID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.ID)
		if err != nil {
			return []domain.Comment{}, err
		}
		q := "tag_id = ?"
		t, err := a.CommentRepo.ReadID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, err
		}
		return t, nil
	} else if form.UserID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.UserID)
		if err != nil {
			return []domain.Comment{}, err
		}
		q := "user_id = ?"
		t, err := a.CommentRepo.ReadID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, err
		}
		return t, nil
	}
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return []domain.Comment{}, err
	}
	span := nInt * 10
	c, err := a.CommentRepo.Read(span)
	if err != nil {
		return []domain.Comment{}, err
	}
	return c, nil
}

func (a *usecase) Update(comment *domain.CommentForm, uid uint) error {
	c, err := a.CommentRepo.ReadID(comment.ID, "id", 1)
	if err != nil {
		return err
	}
	if int(c[0].ID) == int(uid) {
		if err := a.CommentRepo.Update(comment); err != nil {
			return err
		}
	}
	return nil
}

func (a *usecase) Delete(comment *domain.Comment, uid uint) error {
	c, err := a.CommentRepo.ReadID(comment.ID, "id", 1)
	if err != nil {
		return err
	}
	if int(c[0].ID) == int(uid) {
		if err := a.CommentRepo.Delete(comment.ID); err != nil {
			return err
		}
	}
	return nil
}
