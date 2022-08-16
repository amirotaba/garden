package tag

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	UseCase domain.TagUseCase
}

func NewHandler(e *echo.Echo, u domain.TagUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("tag/create", handler.CreateTag)
	res.GET("tag/read", handler.ReadTag)
	res.GET("tag/readID", handler.ReadTagID)
	res.PATCH("tag/update", handler.UpdateTag)
	res.DELETE("tag/delete", handler.DeleteTag)
}

func (m *Handler) CreateTag(e echo.Context) error {
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(tag)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Tag added successfully")
}

func (m *Handler) ReadTag(e echo.Context) error {
	pageNumber := e.QueryParam("page")
	t, err := m.UseCase.Read(pageNumber)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, t)
}

func (m *Handler) ReadTagID(e echo.Context) error {
	id := e.QueryParam("id")
	t, err := m.UseCase.ReadID(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, t)
}

func (m *Handler) UpdateTag(e echo.Context) error {
	tag := new(domain.TagForm)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(tag)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Tag updated successfully")
}

func (m *Handler) DeleteTag(e echo.Context) error {
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(tag)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Tag deleted successfully")
}
