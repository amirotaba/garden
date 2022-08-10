package http

import (
	"garden/internal/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"garden/internal/domain"
)

func (m *UserHandler) CreateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.CreateGarden(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden added successfully.")
}

func (m *UserHandler) ReadGarden(e echo.Context) error {
	form := domain.ReadGardenForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		UserID:     e.QueryParam("user_id"),
		PageNumber: e.QueryParam("page"),
		ID:         e.QueryParam("id"),
	}
	g, code, err := m.GUseCase.ReadGarden(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, g)
}

func (m *UserHandler) UpdateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.GardenForm)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.UpdateGarden(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden updated successfully")
}

func (m *UserHandler) DeleteGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.DeleteGarden(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden has been removed.")
}

func (m *UserHandler) CreateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	location := new(domain.GardenLocation)
	if err := e.Bind(location); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.CreateLocation(location, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location added successfully.")
}

func (m *UserHandler) ReadLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gid := e.QueryParam("garden_id")
	pageNumber := e.QueryParam("page")
	t, code, err := m.GUseCase.ReadLocation(gid, pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocationForm)
	if err := e.Bind(loc); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.UpdateLocation(loc, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location updated successfully")
}

func (m *UserHandler) DeleteLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocation)
	if err := e.Bind(loc); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.DeleteLocation(loc, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location deleted successfully")
}

func (m *UserHandler) CreateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.CreateGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type added successfully")
}

func (m *UserHandler) ReadGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.GUseCase.ReadGardenType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenTypeForm)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.UpdateGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type updated successfully")
}

func (m *UserHandler) DeleteGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.GUseCase.DeleteGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type deleted successfully")
}
