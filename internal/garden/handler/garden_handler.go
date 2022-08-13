package gDel

import (
	"garden/internal/domain"
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	GUseCase    domain.GardenUseCase
}

func NewHandler(e *echo.Echo, u domain.GardenUseCase) {
	handler := &Handler{
		GUseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	res.POST("garden/create", handler.CreateGarden)
	res.GET("garden/read", handler.ReadGarden)
	res.PATCH("garden/update", handler.UpdateGarden)
	res.DELETE("garden/delete", handler.DeleteGarden)

	res.POST("loc/create", handler.CreateLocation)
	res.GET("loc/read", handler.ReadLocation)
	res.PATCH("loc/update", handler.UpdateLocation)
	res.DELETE("loc/delete", handler.DeleteLocation)

	res.POST("garden/type/create", handler.CreateGardenType)
	res.GET("garden/type/read", handler.ReadGardenType)
	res.PATCH("garden/type/update", handler.UpdateGardenType)
	res.DELETE("garden/type/delete", handler.DeleteGardenType)

	e.Logger.Fatal(e.Start(":4000"))
}

func (m *Handler) CreateGarden(e echo.Context) error {
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

func (m *Handler) ReadGarden(e echo.Context) error {
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

func (m *Handler) UpdateGarden(e echo.Context) error {
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

func (m *Handler) DeleteGarden(e echo.Context) error {
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

func (m *Handler) CreateLocation(e echo.Context) error {
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

func (m *Handler) ReadLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gid := e.QueryParam("garden_id")
	pageNumber := e.QueryParam("page")
	t, code, err := m.GUseCase.ReadLocation(gid, pageNumber, uid)
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

	code, err := m.GUseCase.UpdateLocation(loc, uid)
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

	code, err := m.GUseCase.DeleteLocation(loc, uid)
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

	code, err := m.GUseCase.CreateGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type added successfully")
}

func (m *Handler) ReadGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.GUseCase.ReadGardenType(id, uid)
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

	code, err := m.GUseCase.UpdateGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type updated successfully")
}

func (m *Handler) DeleteGardenType(e echo.Context) error {
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
