package http

import (
	"garden/internal/pkg/jwt"
	"net/http"
	"strconv"

	"garden/internal/domain"
	"github.com/labstack/echo/v4"
)

func (m *UserHandler) CreateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TreeUseCase.CreateTree(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree Added successfully.")
}

func (m *UserHandler) ReadTree(e echo.Context) error {
	form := domain.ReadTreeForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		GardenID:   e.QueryParam("garden_id"),
		Tp:         e.QueryParam("type"),
		PageNumber: e.QueryParam("page"),
	}
	tree, code, err := m.TreeUseCase.ReadTree(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, tree)
}

func (m *UserHandler) ReadTreeUser(e echo.Context) error {
	form := domain.ReadTreeUserForm{
		ID:       e.QueryParam("id"),
		GardenID: e.QueryParam("garden_id"),
	}
	tree, code, err := m.TreeUseCase.ReadTreeUser(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, tree)
}

func (m *UserHandler) UpdateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.TreeForm)
	if err := e.Bind(tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TreeUseCase.UpdateTree(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree updated successfully")
}

func (m *UserHandler) DeleteTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TreeUseCase.DeleteTree(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree has been removed.")
}

func (m *UserHandler) CreateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TreeUseCase.CreateTreeType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type added successfully")
}

func (m *UserHandler) ReadTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.TreeUseCase.ReadTreeType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeTypeForm)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TreeUseCase.UpdateTreeType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type updated successfully")
}

func (m *UserHandler) DeleteTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.TreeUseCase.DeleteTreeType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type deleted successfully")
}
