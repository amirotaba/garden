package gardenLoc

import (
	"garden/internal/domain/gardenLocation"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	UseCase gardenLocationDomain.GardenLocUseCase
}

func NewHandler(e *echo.Echo, useCase gardenLocationDomain.GardenLocUseCase) {
	handler := &Handler{
		UseCase: useCase,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("loc/create", handler.Create)
	res.GET("loc/read", handler.Read)
	res.PATCH("loc/update", handler.Update)
	res.DELETE("loc/delete", handler.Delete)
}

func (m *Handler) Create(e echo.Context) error {
	var loc gardenLocationDomain.GardenLocation
	if err := e.Bind(&loc); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(&loc)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Location added successfully.")
}

func (m *Handler) Read(e echo.Context) error {
	form := gardenLocationDomain.GardenLocRead{
		GardenID:   e.QueryParam("garden_id"),
		PageNumber: e.QueryParam("page"),
	}
	t, err := m.UseCase.Read(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, t)
}

func (m *Handler) Update(e echo.Context) error {
	var loc gardenLocationDomain.GardenLocationForm
	if err := e.Bind(&loc); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(&loc)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Location updated successfully")
}

func (m *Handler) Delete(e echo.Context) error {
	var loc gardenLocationDomain.GardenLocation
	if err := e.Bind(&loc); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(&loc)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Location deleted successfully")
}
