package cDel

import (
	"garden/internal/domain"
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase  domain.CommentUseCase
}

func NewHandler(e *echo.Echo, u domain.CommentUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("comment/create", handler.CreateComment)
	res.GET("comment/read", handler.ReadComment)
	res.PATCH("comment/update", handler.UpdateComment)
	res.DELETE("comment/delete", handler.DeleteComment)

	e.Logger.Fatal(e.Start(":4000"))
}


func (m *Handler) CreateComment(e echo.Context) error {
	form := new(domain.Comment)
	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	code, err := m.UseCase.CreateComment(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Comment added successfully")
}

func (m *Handler) ReadComment(e echo.Context) error {
	form := domain.ReadCommentForm{
		ID:         e.QueryParam("id"),
		TreeID:     e.QueryParam("tree_id"),
		TagID:      e.QueryParam("tag_id"),
		UserID:     e.QueryParam("user_id"),
		PageNumber: e.QueryParam("page"),
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
	}
	comments, code, err := m.UseCase.ReadComment(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, comments)
}

func (m *Handler) UpdateComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.CommentForm)
	if err := e.Bind(comment); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.UpdateComment(comment, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "comment updated successfully")
}

func (m *Handler) DeleteComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.Comment)
	if err := e.Bind(comment); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.DeleteComment(comment, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Comment has been removed.")
}
