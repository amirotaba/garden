package gardenType

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	UseCase domain.GardenTypeUseCase
}

func NewHandler(e *echo.Echo, useCase domain.GardenTypeUseCase) {
	handler := &Handler{
		UseCase: useCase,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("garden/type/create", handler.Create)
	res.GET("garden/type/read", handler.Read)
	res.PATCH("garden/type/update", handler.Update)
	res.DELETE("garden/type/delete", handler.Delete)
}

func (m *Handler) Create(e echo.Context) error {
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(gardenType)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Garden type added successfully")
}

func (m *Handler) Read(e echo.Context) error {
	id := e.QueryParam("id")
	t, err := m.UseCase.Read(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, t)
}

func (m *Handler) Update(e echo.Context) error {
	gardenType := new(domain.GardenTypeForm)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(gardenType)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Garden type updated successfully")
}

func (m *Handler) Delete(e echo.Context) error {
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(gardenType)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Garden type deleted successfully")
}
