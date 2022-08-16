package treeType

import (
	"garden/internal/domain/treeType"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	UseCase treeTypeDomain.TreeTypeUseCase
}

func NewHandler(e *echo.Echo, u treeTypeDomain.TreeTypeUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("tree/type/create", handler.Create)
	res.GET("tree/type/read", handler.Read)
	res.PATCH("tree/type/update", handler.Update)
	res.DELETE("tree/type/delete", handler.Delete)
}

func (m *Handler) Create(e echo.Context) error {
	var treeType treeTypeDomain.TreeType
	if err := e.Bind(&treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(&treeType)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Tree type added successfully")
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
	var treeType treeTypeDomain.TreeTypeForm
	if err := e.Bind(&treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(&treeType)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Tree type updated successfully")
}

func (m *Handler) Delete(e echo.Context) error {
	var treeType treeTypeDomain.TreeType
	if err := e.Bind(&treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(&treeType)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Tree type deleted successfully")
}
