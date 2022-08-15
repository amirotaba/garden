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
}

func (m *Handler) CreateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.Create(garden, uid)
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
	g, code, err := m.UseCase.Read(form)
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

	code, err := m.UseCase.Update(garden, uid)
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

	code, err := m.UseCase.Delete(garden, uid)
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

	code, err := m.UseCase.CreateLocation(location, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location added successfully.")
}

func (m *Handler) ReadLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gid := e.QueryParam("garden_id")
	pageNumber := e.QueryParam("page")
	t, code, err := m.UseCase.ReadLocation(gid, pageNumber, uid)
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

	code, err := m.UseCase.UpdateLocation(loc, uid)
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

	code, err := m.UseCase.DeleteLocation(loc, uid)
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

	code, err := m.UseCase.CreateType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type added successfully")
}

func (m *Handler) ReadGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.UseCase.ReadType(id, uid)
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

	code, err := m.UseCase.UpdateType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type updated successfully")
}

func (m *Handler) DeleteGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.DeleteType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type deleted successfully")
}
