package tree

import (
	"garden/internal/domain/tree"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase treeDomain.TreeUseCase
}

func NewHandler(e *echo.Echo, u treeDomain.TreeUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("tree/create", handler.Create)
	res.GET("tree/read", handler.Read)
	res.GET("tree/readUser", handler.ReadUser)
	res.PATCH("tree/update", handler.Update)
	res.DELETE("tree/delete", handler.Delete)
}

func (m *Handler) Create(e echo.Context) error {
	var tree treeDomain.Tree
	if err := e.Bind(&tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Create(&tree)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, "Tree Added successfully.")
}

func (m *Handler) Read(e echo.Context) error {
	form := treeDomain.ReadTreeForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		GardenID:   e.QueryParam("garden_id"),
		Tp:         e.QueryParam("type"),
		PageNumber: e.QueryParam("page"),
	}
	tree, err := m.UseCase.Read(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, tree)
}

func (m *Handler) ReadUser(e echo.Context) error {
	form := treeDomain.ReadTreeUserForm{
		ID:       e.QueryParam("id"),
		GardenID: e.QueryParam("garden_id"),
	}
	tree, err := m.UseCase.ReadUser(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, tree)
}

func (m *Handler) Update(e echo.Context) error {
	var tree treeDomain.TreeForm
	if err := e.Bind(&tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(&tree)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Tree updated successfully")
}

func (m *Handler) Delete(e echo.Context) error {
	var tree treeDomain.Tree
	if err := e.Bind(&tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(&tree)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Tree has been removed.")
}
