package garden

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase domain.GardenUseCase
}

func NewHandler(e *echo.Echo, useCase domain.GardenUseCase) {
	handler := &Handler{
		UseCase: useCase,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("garden/create", handler.Create)
	res.GET("garden/read", handler.Read)
	res.PATCH("garden/update", handler.Update)
	res.DELETE("garden/delete", handler.Delete)
}

func (m *Handler) Create(e echo.Context) error {
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(garden)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Garden added successfully.")
}

func (m *Handler) Read(e echo.Context) error {
	form := domain.ReadGardenForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		UserID:     e.QueryParam("user_id"),
		PageNumber: e.QueryParam("page"),
		ID:         e.QueryParam("id"),
	}
	g, err := m.UseCase.Read(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, g)
}

func (m *Handler) Update(e echo.Context) error {
	garden := new(domain.GardenForm)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(garden)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Garden updated successfully")
}

func (m *Handler) Delete(e echo.Context) error {
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(garden)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Garden has been removed.")
}
