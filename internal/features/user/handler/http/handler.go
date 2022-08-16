package user

import (
	"garden/internal/domain/user"
	"garden/internal/middleware/access"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase userDomain.UserUseCase
}

func NewHandler(e *echo.Echo, u userDomain.UserUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))
	s := access.NewStats()
	res.Use(s.Process)

	e.POST("signUp", handler.SignUp)
	e.POST("signIn", handler.SignIn)
	res.GET("account", handler.Account)
	res.GET("useraccount", handler.UserAccount)
	res.PATCH("update", handler.UpdateUser)
	res.DELETE("delete", handler.DeleteUser)
}

func (m *Handler) SignIn(e echo.Context) error {
	var loginForm userDomain.LoginForm
	if err := e.Bind(&loginForm); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	u, err := m.UseCase.SignIn(&loginForm)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	msg := userDomain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: u,
	}
	return e.JSON(http.StatusOK, msg)
}

func (m *Handler) SignUp(e echo.Context) error {
	var user userDomain.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	u, err := m.UseCase.Create(user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, u)
}

func (m *Handler) Account(e echo.Context) error {
	form := userDomain.AccountForm{
		Uid:        jwt.UserID(e),
		Tp:         e.QueryParam("type"),
		PageNumber: e.QueryParam("page"),
	}
	users, err := m.UseCase.Read(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, users)
}

func (m *Handler) UserAccount(e echo.Context) error {
	// use struct instead of map
	form := userDomain.UserAccountForm{
		Uid:      strconv.Itoa(int(jwt.UserID(e))),
		Username: e.QueryParam("username"),
		ID:       e.QueryParam("id"),
	}
	users, err := m.UseCase.UserRead(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, users)
}

func (m *Handler) UpdateUser(e echo.Context) error {
	var user userDomain.UserForm
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Update(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User updated successfully")
}

func (m *Handler) DeleteUser(e echo.Context) error {
	var user userDomain.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err := m.UseCase.Delete(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User deleted successfully")
}
