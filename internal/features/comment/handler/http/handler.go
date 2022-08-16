package comment

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase domain.CommentUseCase
}

func NewHandler(e *echo.Echo, useCase domain.CommentUseCase) {
	handler := &Handler{
		UseCase: useCase,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("comment/create", handler.CreateComment)
	res.GET("comment/read", handler.ReadComment)
	res.PATCH("comment/update", handler.UpdateComment)
	res.DELETE("comment/delete", handler.DeleteComment)
}

func (m *Handler) CreateComment(e echo.Context) error {
	form := new(domain.Comment)
	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	err := m.UseCase.Create(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Comment added successfully")
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
	comments, err := m.UseCase.Read(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, comments)
}

func (m *Handler) UpdateComment(e echo.Context) error {
	form := new(domain.UpdateCommentForm)
	if err := e.Bind(form.Comment); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	form.Uid = jwt.UserID(e)
	err := m.UseCase.Update(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "comment updated successfully")
}

func (m *Handler) DeleteComment(e echo.Context) error {
	form := new(domain.UpdateCommentForm)
	if err := e.Bind(form.Comment); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	form.Uid = jwt.UserID(e)
	err := m.UseCase.Delete(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Comment has been removed.")
}
