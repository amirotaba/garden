package tree

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase domain.TreeUseCase
}

func NewHandler(e *echo.Echo, u domain.TreeUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("tree/create", handler.CreateTree)
	res.GET("tree/read", handler.ReadTree)
	res.GET("tree/readUser", handler.ReadTreeUser)
	res.PATCH("tree/update", handler.UpdateTree)
	res.DELETE("tree/delete", handler.DeleteTree)

	res.POST("tree/type/create", handler.CreateTreeType)
	res.GET("tree/type/read", handler.ReadTreeType)
	res.PATCH("tree/type/update", handler.UpdateTreeType)
	res.DELETE("tree/type/delete", handler.DeleteTreeType)
}

func (m *Handler) CreateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.Create(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree Added successfully.")
}

func (m *Handler) ReadTree(e echo.Context) error {
	form := domain.ReadTreeForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		GardenID:   e.QueryParam("garden_id"),
		Tp:         e.QueryParam("type"),
		PageNumber: e.QueryParam("page"),
	}
	tree, code, err := m.UseCase.Read(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, tree)
}

func (m *Handler) ReadTreeUser(e echo.Context) error {
	form := domain.ReadTreeUserForm{
		ID:       e.QueryParam("id"),
		GardenID: e.QueryParam("garden_id"),
	}
	tree, code, err := m.UseCase.ReadUser(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, tree)
}

func (m *Handler) UpdateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.TreeForm)
	if err := e.Bind(tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.Update(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree updated successfully")
}

func (m *Handler) DeleteTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.Delete(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree has been removed.")
}

func (m *Handler) CreateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.CreateType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type added successfully")
}

func (m *Handler) ReadTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.UseCase.ReadType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) UpdateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeTypeForm)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.UpdateType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type updated successfully")
}

func (m *Handler) DeleteTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.DeleteType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type deleted successfully")
}
