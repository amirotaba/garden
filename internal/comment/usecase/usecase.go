package cUsecase

import (
	"errors"
	"strconv"
	"strings"

	"garden/internal/domain"
)

type commentUsecase struct {
	CommentRepo domain.CommentRepository
	ServiceRepo domain.ServiceRepository
	UserRepo    domain.UserRepository
}

func NewCommentUseCase(r domain.Repositories) domain.CommentUseCase {
	return &commentUsecase{
		CommentRepo: r.Comment,
		ServiceRepo: r.Service,
		UserRepo:    r.User,
	}
}

func (a *commentUsecase) CreateComment(comment *domain.Comment) (int, error) {
	if err := a.CommentRepo.CreateComment(comment); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *commentUsecase) ReadComment(form domain.ReadCommentForm) ([]domain.Comment, int, error) {
	if form.ID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.ID)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		q := "id = ?"
		t, err := a.CommentRepo.ReadCommentID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		return t, 200, nil
	} else if form.TreeID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.TreeID)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		q := "tree_id = ?"
		t, err := a.CommentRepo.ReadCommentID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		return t, 200, nil
	} else if form.TagID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.ID)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		q := "tag_id = ?"
		t, err := a.CommentRepo.ReadCommentID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		return t, 200, nil
	} else if form.UserID != "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		span := nInt * 10
		idInt, err := strconv.Atoi(form.UserID)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		q := "user_id = ?"
		t, err := a.CommentRepo.ReadCommentID(uint(idInt), q, span)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		return t, 200, nil
	}
	var boolean bool
	uidInt, err := strconv.Atoi(form.Uid)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/comment/read")
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.Comment{}, 403, errors.New("you can't access to this page")
	}
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	span := nInt * 10
	c, err := a.CommentRepo.ReadComment(span)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	return c, 200, nil
}

func (a *commentUsecase) UpdateComment(comment *domain.CommentForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tree/update")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	c, err := a.CommentRepo.ReadCommentID(comment.ID, "id", 1)
	if err != nil {
		return 400, err
	}
	if boolean || int(c[0].ID) == uidInt {
		if err := a.CommentRepo.UpdateComment(comment); err != nil {
			return 400, err
		}
		return 201, nil
	}
	return 403, errors.New("you can't access to this page")
}

func (a *commentUsecase) DeleteComment(comment *domain.Comment, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tree/update")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	c, err := a.CommentRepo.ReadCommentID(comment.ID, "id", 1)
	if boolean || int(c[0].ID) == uidInt {
		if err := a.CommentRepo.DeleteComment(comment.ID); err != nil {
			return 400, err
		}
		return 204, nil
	}
	return 403, errors.New("you can't access to this page")
}
