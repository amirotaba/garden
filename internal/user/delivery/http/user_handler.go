package uDel

import (
	"garden/internal/domain"
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase  domain.UserUseCase
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
	res.GET("user/account", handler.UserAccount)
	res.PATCH("update", handler.UpdateUser)
	res.DELETE("delete", handler.DeleteUser)

	res.POST("userType/create", handler.CreateUserType)
	res.GET("user/type/read", handler.ReadUserType)
	res.PATCH("user/type/update", handler.UpdateUserType)
	res.PATCH("user/type/addAccess", handler.AddAccess)
	res.DELETE("user/type/delete", handler.DeleteUsertype)

	e.Logger.Fatal(e.Start(":4000"))
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
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(403, err.Error())
	}

	code, err := m.UseCase.SignUp(user)
	if err != nil {
		return e.JSON(code, err.Error())
	}

	u := domain.UserResponse{UserName: user.UserName}
	msg := &domain.SignUpMessage{
		Text:     "you signed up as, ",
		UserName: u.UserName,
	}
	return e.JSON(http.StatusOK, msg)
}

func (m *Handler) Account(e echo.Context) error {
	form := domain.AccountForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		Tp:         e.QueryParam("type"),
		PageNumber: e.QueryParam("page"),
	}
	users, code, err := m.UseCase.Account(form)
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
	users, code, err := m.UseCase.UserAccount(form)
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

	code, err := m.UseCase.UpdateUser(user, uid)
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

	code, err := m.UseCase.DeleteUser(user, uid)
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

	code, err := m.UseCase.CreateUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type added successfully")
}

func (m *Handler) ReadUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.UseCase.ReadUserType(id, uid)
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

	code, err := m.UseCase.UpdateUserType(usertype, uid)
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

	code, err := m.UseCase.DeleteUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type deleted successfully")
}
