package tagDel

import (
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"

	"garden/internal/domain"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCase  domain.TagUseCase
}

func NewHandler(e *echo.Echo, u domain.TagUseCase) {
	handler := &Handler{
		UseCase:    u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("tag/create", handler.CreateTag)
	res.GET("tag/read", handler.ReadTag)
	res.GET("tag/readID", handler.ReadTagID)
	res.PATCH("tag/update", handler.UpdateTag)
	res.DELETE("tag/delete", handler.DeleteTag)

	e.Logger.Fatal(e.Start(":4000"))
}

func (m *Handler) CreateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.CreateTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag added successfully")
}

func (m *Handler) ReadTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	pageNumber := e.QueryParam("page")
	t, code, err := m.UseCase.ReadTag(pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) ReadTagID(e echo.Context) error {
	id := e.QueryParam("id")
	t, code, err := m.UseCase.ReadTagID(id)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) UpdateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.TagForm)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.UpdateTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag updated successfully")
}

func (m *Handler) DeleteTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.DeleteTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag deleted successfully")
}
