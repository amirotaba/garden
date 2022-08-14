package user

import (
	"garden/internal/domain"
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase domain.UserUseCase
}

func NewHandler(e *echo.Echo, u domain.UserUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	e.POST("signUp", handler.SignUp)
	e.POST("signIn", handler.SignIn)
	res.GET("account", handler.Account)
	res.GET("useraccount", handler.UserAccount)
	res.PATCH("update", handler.UpdateUser)
	res.DELETE("delete", handler.DeleteUser)

	res.POST("usertype/create", handler.CreateUserType)
	res.GET("usertype/read", handler.ReadUserType)
	res.PATCH("usertype/update", handler.UpdateUserType)
	res.PATCH("usertype/addAccess", handler.AddAccess)
	res.DELETE("usertype/delete", handler.DeleteUsertype)
}

func (m *Handler) SignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return e.JSON(403, err.Error())
	}

	u, code, err := m.UseCase.SignIn(loginForm)
	if err != nil {
		return e.JSON(code, err.Error())
	}

	msg := domain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: u,
	}
	return e.JSON(code, msg)
}

func (m *Handler) SignUp(e echo.Context) error {
	var user domain.User
	if err := e.Bind(user); err != nil {
		return e.JSON(403, err.Error())
	}

	u, err := m.UseCase.Create(user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, u)
}

func (m *Handler) Account(e echo.Context) error {
	form := domain.AccountForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		Tp:         e.QueryParam("type"),
		PageNumber: e.QueryParam("page"),
	}
	users, code, err := m.UseCase.Read(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, users)
}

func (m *Handler) UserAccount(e echo.Context) error {
	// use struct instead of map
	form := domain.UserAccountForm{
		Uid:      strconv.Itoa(int(jwt.UserID(e))),
		Username: e.QueryParam("username"),
		ID:       e.QueryParam("id"),
	}
	users, code, err := m.UseCase.UserRead(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, users)
}

func (m *Handler) UpdateUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.UserForm)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.Update(user, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User updated successfully")
}

func (m *Handler) DeleteUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.Delete(user, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User deleted successfully")
}

func (m *Handler) CreateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.CreateType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type added successfully")
}

func (m *Handler) ReadUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.UseCase.ReadType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *Handler) UpdateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserTypeForm)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.UpdateType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *Handler) AddAccess(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	add := new(domain.AccessForm)
	if err := e.Bind(add); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.UpdateAccess(add, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *Handler) DeleteUsertype(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UseCase.DeleteType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type deleted successfully")
}
