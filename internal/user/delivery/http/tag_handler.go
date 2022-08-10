package http

import (
	"garden/internal/pkg/jwt"
	"net/http"
	"strconv"

	"garden/internal/domain"
	"github.com/labstack/echo/v4"
)

func (m *UserHandler) CreateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TagUseCase.CreateTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag added successfully")
}

func (m *UserHandler) ReadTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	pageNumber := e.QueryParam("page")
	t, code, err := m.TagUseCase.ReadTag(pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) ReadTagID(e echo.Context) error {
	id := e.QueryParam("id")
	t, code, err := m.TagUseCase.ReadTagID(id)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.TagForm)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TagUseCase.UpdateTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag updated successfully")
}

func (m *UserHandler) DeleteTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TagUseCase.DeleteTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag deleted successfully")
}
