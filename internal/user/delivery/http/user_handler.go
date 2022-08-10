package http

import (
	"garden/internal/domain"
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (m *UserHandler) SignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return e.JSON(403, err.Error())
	}

	u, code, err := m.UUseCase.SignIn(loginForm)
	if err != nil {
		return e.JSON(code, err.Error())
	}

	msg := domain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: u,
	}
	return e.JSON(code, msg)
}

func (m *UserHandler) SignUp(e echo.Context) error {
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(403, err.Error())
	}

	code, err := m.UUseCase.SignUp(user)
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

func (m *UserHandler) Account(e echo.Context) error {
	form := domain.AccountForm{
		Uid:        strconv.Itoa(int(jwt.UserID(e))),
		Tp:         e.QueryParam("type"),
		PageNumber: e.QueryParam("page"),
	}
	users, code, err := m.UUseCase.Account(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, users)
}

func (m *UserHandler) UserAccount(e echo.Context) error {
	// use struct instead of map
	form := domain.UserAccountForm{
		Uid:      strconv.Itoa(int(jwt.UserID(e))),
		Username: e.QueryParam("username"),
		ID:       e.QueryParam("id"),
	}
	users, code, err := m.UUseCase.UserAccount(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, users)
}

func (m *UserHandler) UpdateUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.UserForm)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UUseCase.UpdateUser(user, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User updated successfully")
}

func (m *UserHandler) DeleteUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UUseCase.DeleteUser(user, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User deleted successfully")
}

func (m *UserHandler) CreateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UUseCase.CreateUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type added successfully")
}

func (m *UserHandler) ReadUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.UUseCase.ReadUserType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserTypeForm)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UUseCase.UpdateUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *UserHandler) AddAccess(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	add := new(domain.AccessForm)
	if err := e.Bind(add); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UUseCase.UpdateAccess(add, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *UserHandler) DeleteUsertype(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	code, err := m.UUseCase.DeleteUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type deleted successfully")
}
