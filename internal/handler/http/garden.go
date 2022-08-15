package deliver

import (
	"garden/internal/domain"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (m *Handler) CreateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.Create(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase added successfully.")
}

func (m *Handler) ReadGarden(e echo.Context) error {
	form := domain.ReadGardenForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		UserID:     e.QueryParam("user_id"),
		PageNumber: e.QueryParam("page"),
		ID:         e.QueryParam("id"),
	}
	g, code, err := m.Garden.Read(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, g)
}

func (m *Handler) UpdateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.GardenForm)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.Update(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase updated successfully")
}

func (m *Handler) DeleteGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.Delete(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase has been removed.")
}

func (m *Handler) CreateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	location := new(domain.GardenLocation)
	if err := e.Bind(location); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.CreateLocation(location, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location added successfully.")
}

func (m *Handler) ReadLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gid := e.QueryParam("garden_id")
	pageNumber := e.QueryParam("page")
	t, code, err := m.Garden.ReadLocation(gid, pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) UpdateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocationForm)
	if err := e.Bind(loc); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.UpdateLocation(loc, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location updated successfully")
}

func (m *Handler) DeleteLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocation)
	if err := e.Bind(loc); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.DeleteLocation(loc, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location deleted successfully")
}

func (m *Handler) CreateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.CreateType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase type added successfully")
}

func (m *Handler) ReadGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.Garden.ReadType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) UpdateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenTypeForm)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.UpdateType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase type updated successfully")
}

func (m *Handler) DeleteGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.Garden.DeleteType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "UseCase type deleted successfully")
}
