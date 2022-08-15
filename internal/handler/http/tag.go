package deliver

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (m *Handler) CreateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Tag.Create(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase added successfully")
}

func (m *Handler) ReadTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	pageNumber := e.QueryParam("page")
	t, code, err := m.Tag.Read(pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) ReadTagID(e echo.Context) error {
	id := e.QueryParam("id")
	t, code, err := m.Tag.ReadID(id)
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

	code, err := m.Tag.Update(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase updated successfully")
}

func (m *Handler) DeleteTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Tag.Delete(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase deleted successfully")
}
