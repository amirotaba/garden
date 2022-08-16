package userType

import (
	"garden/internal/domain/userType"
	"garden/internal/middleware/access"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	UseCase userTypeDomain.UserTypeUseCase
}

func NewHandler(e *echo.Echo, u userTypeDomain.UserTypeUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))
	s := access.NewStats()
	res.Use(s.Process)

	res.POST("usertype/create", handler.Create)
	res.GET("usertype/read", handler.Read)
	res.PATCH("usertype/update", handler.Update)
	res.PATCH("usertype/addAccess", handler.AddAccess)
	res.DELETE("usertype/delete", handler.Delete)
}

func (m *Handler) Create(e echo.Context) error {
	var usertype userTypeDomain.UserType
	if err := e.Bind(&usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(&usertype)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "User type added successfully")
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
	var usertype userTypeDomain.UserTypeForm
	if err := e.Bind(&usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(&usertype)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User type updated successfully")
}

func (m *Handler) AddAccess(e echo.Context) error {
	var add userTypeDomain.AccessForm
	if err := e.Bind(&add); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	add.ID = jwt.UserID(e)
	err := m.UseCase.UpdateAccess(&add)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User type updated successfully")
}

func (m *Handler) Delete(e echo.Context) error {
	var usertype userTypeDomain.UserType
	if err := e.Bind(&usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(&usertype)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User type deleted successfully")
}
