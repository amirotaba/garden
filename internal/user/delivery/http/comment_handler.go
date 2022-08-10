package http

import (
	"garden/internal/domain"
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (m *UserHandler) CreateComment(e echo.Context) error {
	form := new(domain.Comment)
	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	code, err := m.CUseCase.CreateComment(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Comment added successfully")
}

func (m *UserHandler) ReadComment(e echo.Context) error {
	form := domain.ReadCommentForm{
		ID:         e.QueryParam("id"),
		TreeID:     e.QueryParam("tree_id"),
		TagID:      e.QueryParam("tag_id"),
		UserID:     e.QueryParam("user_id"),
		PageNumber: e.QueryParam("page"),
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
	}
	comments, code, err := m.CUseCase.ReadComment(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, comments)
}

func (m *UserHandler) UpdateComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.CommentForm)
	if err := e.Bind(comment); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.CUseCase.UpdateComment(comment, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "comment updated successfully")
}

func (m *UserHandler) DeleteComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.Comment)
	if err := e.Bind(comment); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.CUseCase.DeleteComment(comment, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Comment has been removed.")
}
