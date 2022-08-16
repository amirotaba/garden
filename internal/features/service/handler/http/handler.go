package service

import (
	"garden/internal/domain/service"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	UseCase serviceDomain.ServiceUseCase
}

func NewHandler(e *echo.Echo, u serviceDomain.ServiceUseCase) {
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
	var service serviceDomain.Service
	if err := e.Bind(&service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(&service)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Service added successfully")
}

func (m *Handler) ReadService(e echo.Context) error {
	t, err := m.UseCase.Read()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, t)
}

func (m *Handler) UpdateService(e echo.Context) error {
	var service serviceDomain.ServiceForm
	if err := e.Bind(&service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(&service)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User type updated successfully")
}

func (m *Handler) DeleteService(e echo.Context) error {
	var service serviceDomain.Service
	if err := e.Bind(&service); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(&service)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Service deleted successfully")
}

func (m *Handler) addRoutes(e *echo.Echo) {
	r := e.Routes()
	for i := range r {
		service := &serviceDomain.Service{
			Name:   r[i].Name,
			Url:    r[i].Path,
			Method: r[i].Method,
		}
		_ = m.UseCase.Create(service)
	}
}
