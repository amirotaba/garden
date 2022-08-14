package deliver

import (
	"garden/internal/pkg/jwt"
	"net/http"
	"strconv"

	"garden/internal/domain"
	"github.com/labstack/echo/v4"
)

func (m *Handler) CreateService(e echo.Context) error {
	service := new(domain.Service)
	if err := e.Bind(service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Service.Create(service)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Service added successfully")
}

func (m *Handler) ReadService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	t, code, err := m.Service.Read(uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) UpdateService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	service := new(domain.ServiceForm)
	if err := e.Bind(service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Service.Update(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *Handler) DeleteService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	service := new(domain.Service)
	if err := e.Bind(service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Service.Delete(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type deleted successfully")
}

func (m *Handler) addRoutes(e *echo.Echo) {
	r := e.Routes()
	for i := range r {
		service := &domain.Service{
			Name:   r[i].Name,
			Url:    r[i].Path,
			Method: r[i].Method,
		}
		_, _ = m.Service.Create(service)
	}
}
