package http

import (
	"garden/internal/pkg/jwt"
	"net/http"
	"strconv"

	"garden/internal/domain"
	"github.com/labstack/echo/v4"
)

func (m *UserHandler) CreateService(e echo.Context) error {
	service := new(domain.Service)
	if err := e.Bind(service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.SUseCase.CreateService(service)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Service added successfully")
}

func (m *UserHandler) ReadService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	t, code, err := m.SUseCase.ReadService(uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	service := new(domain.ServiceForm)
	if err := e.Bind(service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.SUseCase.UpdateService(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *UserHandler) DeleteService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	service := new(domain.Service)
	if err := e.Bind(service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.SUseCase.DeleteService(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type deleted successfully")
}

func (m *UserHandler) addRoutes(e *echo.Echo) {
	r := e.Routes()
	for i := range r {
		service := &domain.Service{
			Name:   r[i].Name,
			Url:    r[i].Path,
			Method: r[i].Method,
		}
		_, _ = m.SUseCase.CreateService(service)
	}
}
