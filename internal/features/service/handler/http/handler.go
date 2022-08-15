package service

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase domain.ServiceUseCase
}

func NewHandler(e *echo.Echo, u domain.ServiceUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("service/create", handler.CreateService)
	res.GET("service/read", handler.ReadService)
	res.PATCH("service/update", handler.UpdateService)
	res.DELETE("service/delete", handler.DeleteService)

	handler.addRoutes(e)
}

func (m *Handler) CreateService(e echo.Context) error {
	service := new(domain.Service)
	if err := e.Bind(service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.Create(service)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Service added successfully")
}

func (m *Handler) ReadService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	t, code, err := m.UseCase.Read(uid)
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

	code, err := m.UseCase.Update(service, uid)
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

	code, err := m.UseCase.Delete(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Service deleted successfully")
}

func (m *Handler) addRoutes(e *echo.Echo) {
	r := e.Routes()
	for i := range r {
		service := &domain.Service{
			Name:   r[i].Name,
			Url:    r[i].Path,
			Method: r[i].Method,
		}
		_, _ = m.UseCase.Create(service)
	}
}
