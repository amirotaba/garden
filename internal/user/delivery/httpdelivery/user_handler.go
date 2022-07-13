package httpdelivery

import (
	"garden/internal/domain"
	"sudoku/internal/domain"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	AUsecae domain.UserUsecase
}

func NewuserHandler(e *echo.Echo, us domain.UserUsecase, ad domain.AdminUsecase, fr domain.FarmerUsecase)  {
	handler := &UserHandler{
		AUsecase: us,
	}
	e.POST("user/signup", handler.SignUp)
	e.POST("user/signin", handler.SignIn)
	e.GET("user/account/:username", handler.Account)
	e.POST("user/addcmt", handler.Comment)

	e.POST("admin/login", handler.Login)
	e.GET("admin/grd/shw", handler.ShowGarden)
	e.POST("admin/grd/rmv", handler.RemoveGarden)
	e.POST("admin/grd/add", handler.AddGarden)
	e.POST("admin/frm/add", handler.AddFarmer)

	e.POST("farmer/login", handler.Login)
	e.GET("farmer/trees/shw/:id", handler.ShowTrees)
	e.POST("farmer/cmt/shw", handler.ShowComments)
	e.POST("farmer/tree/add", handler.AddTree)
	e.POST("farmer/tree/rmv", handler.RemoveTree)
	e.POST("farmer/tree/atnd", handler.AddAttend)
}

//user

func (m *UserHandler) SignUp(e echo.Context) error {
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return err
	}
	if err := m.AUsecase.SignUp(user);err != nil {
		return err
	}
	u := domain.UserResponse{UserName: user.UserName, Email: user.Email}
	msg := &domain.AuthMessage{
		Text:     "you logged in as, ",
		UserInfo: &u,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) SignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return err
	}
	u, err := m.AUsecase.SignIn(loginForm.PassWord, loginForm.Email)
	if err != nil {
		return err
	}
	if u.UserName == "" {
		msg := domain.AuthMessage{
			Text: "incorrect password",
		}
		return e.JSON(200, msg)
	}
	msg := domain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: &u,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) Account(e echo.Context) error {
	username := e.Param("username")
	user, err := m.AUsecase.Account(username)
	if err != nil {
		msg := &domain.AuthMessage{
			Text: "User not found",
		}
		return e.JSON(200, msg)
	}
	msg := &domain.AuthMessage{
		Text:     "User info: ",
		UserInfo: &user,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) Comment(e echo.Context) error {
	var form *domain.CommentForm
	if err := e.Bind(form); err != nil {
		return err
	}
	if err := m.AUsecase.Comment(form.ID, form.TreeID, form.Text); err !- nil {
		return err
	}
	return e.JSON(200, "Comment saved")
}

//admin

(m *AdminHandler) ShowGarden(e echo.Context) error {
	g, err := m.AUsecase.ShowGarden()
	if err != nil {
		return err
	}
	return e.JSON(200, g)
}

(m *AdminHandler) RemoveGarden(e echo.Context) error {
	var id int
	if err := e.Bind(id); err != nil {
		return err
	}
	if err := m.AUsecase.RemoveGarden(id); err != nil {
		return err
	}
	return e.JSON(200, "Garden has been removed.")
}

(m *AdminHandler) AddGarden(e echo.Context) error {
	var garden domain.Garden
	if err := e.Bind(garden); err != nil {
		return err
	}
	if err := m.AUsecase.AddGarden(garden); err != nil {
		return err
	}
	return e.JSON(200, "Garden added successfuly.")
}

(m *AdminHandler) AddFarmer(e echo.Context) error {
	var farmer domain.Farmer
	if err := e.Bind(farmer); err != nil {
		return err
	}
	if err != m.AUsecase.AddFarmer(farmer); err != nil {
		return err
	}
	return e.JSON(200, "farmer added successfuly.")
}

//farmer

(m *FarmerHandler) ShowTrees(e echo.Context) error {
	id := e.Param("id")
	t, err := m.AUsecase.ShowTrees(id)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

(m *FarmerHandler) ShowComment(e echo.Context) error {
	var tree domain.Tree
	if err := e.Bind(tree); err != nil {
		return err
	}
	comment, err := m.AUsecase.ShowComments(tree.ID)
	if err != nil {
		return err
	}
	return e.JSON(200, comment)
}

(m *FarmerHandler) AddTree(e echo.Context) error {
	var tree domain.Tree
	if err := e.Bind(tree); err != nil {
		return err
	}
	if err != m.AUsecase.AddTree(tree); err != nil {
		return err
	}
	return e.JSON(200, "Tree Added successfuly.")
}

(m *FarmerHandler) RemoveTree(e echo.Context) error {
	var tree domain.Tree
	if err := e.Bind(tree); err != nil {
		return err
	}
	if err := m.AUsecase.RemoveTree(tree.ID); err != nil {
		return err
	}
	return e.JSON(200, "Tree has been removed.")
}

(m *FarmerHandler) AddAttend(e echo.Context) error {
	var tree domain.Tree
	if err := e.Bind(tree); err != nil {
		return err
	}
	if err := m.AUsecase.AddAttend(tree); err != nil {
		return err
	}
	return e.JSON(200, "Attend added.")
}