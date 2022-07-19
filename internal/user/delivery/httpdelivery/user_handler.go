package httpdelivery

import (
	"garden/internal/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UUsecase domain.UserUsecase
	AUsecase domain.AdminUsecase
	FUsecase domain.FarmerUsecase
}

func NewUserHandler(e *echo.Echo, us domain.UserUsecase, au domain.AdminUsecase, fu domain.FarmerUsecase) {
	handler := &UserHandler{
		UUsecase: us,
		AUsecase: au,
		FUsecase: fu,
	}
	e.POST("user/signup", handler.SignUp)
	e.POST("user/signin", handler.USignIn)
	e.GET("user/account/:username", handler.Account)
	e.POST("user/addcmt", handler.Comment)

	e.POST("admin/signin", handler.ASignIn)
	e.GET("admin/grd/shw", handler.ShowGarden)
	e.GET("admin/grd/rmv/:username/:id", handler.RemoveGarden)
	e.POST("admin/grd/add/:username", handler.AddGarden)
	e.POST("admin/frm/add/:username", handler.AddFarmer)

	e.POST("farmer/signin", handler.FSignIn)
	e.GET("farmer/tree/shw/:farmerid", handler.ShowTrees)
	e.GET("farmer/cmt/shw/:farmerid/:treeid", handler.ShowComments)
	e.POST("farmer/tree/add/", handler.AddTree)
	e.GET("farmer/tree/rmv/:farmerid/:treeid", handler.RemoveTree)
	e.POST("farmer/tree/attend", handler.AddAttend)

	e.Logger.Fatal(e.Start(":4000"))
}

//user

func (m *UserHandler) SignUp(e echo.Context) error {
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return err
	}
	if err := m.UUsecase.SignUp(user); err != nil {
		return err
	}
	u := domain.UserResponse{UserName: user.UserName, Email: user.Email}
	msg := &domain.SignUpMessage{
		Text:     "you signed up as, ",
		UserName: u.UserName,
		Email:    u.Email,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) USignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return err
	}
	u, err := m.UUsecase.USignIn(loginForm.Username, loginForm.Password)
	if err != nil {
		return err
	}
	msg := domain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: u,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) Account(e echo.Context) error {
	username := e.Param("username")
	user, err := m.UUsecase.Account(username)
	if err != nil {
		msg := &domain.AuthMessage{
			Text: "User not found",
		}
		return e.JSON(200, msg)
	}
	msg := &domain.AuthMessage{
		Text:     "User info: ",
		UserInfo: user,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) Comment(e echo.Context) error {
	form := new(domain.CommentForm)
	if err := e.Bind(form); err != nil {
		return err
	}
	if err := m.UUsecase.Comment(form.ID, form.TreeID, form.Text); err != nil {
		return err
	}
	return e.JSON(200, "Comment saved")
}

//admin

func (m *UserHandler) ASignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return err
	}
	u, err := m.AUsecase.ASignIn(loginForm.Username, loginForm.Password)
	if err != nil {
		return err
	}
	msg := domain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: u,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) ShowGarden(e echo.Context) error {
	g, err := m.AUsecase.ShowGarden()
	if err != nil {
		return err
	}
	return e.JSON(200, g)
}

func (m *UserHandler) RemoveGarden(e echo.Context) error {
	u := e.Param("username")
	id := e.Param("id")
	if err := m.AUsecase.RemoveGarden(id, u); err != nil {
		return err
	}
	return e.JSON(200, "Garden has been removed.")
}

func (m *UserHandler) AddGarden(e echo.Context) error {
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return err
	}
	garden.CreatedBy = e.Param("username")
	if err := m.AUsecase.AddGarden(garden); err != nil {
		return err
	}
	return e.JSON(200, "Garden added successfuly.")
}

func (m *UserHandler) AddFarmer(e echo.Context) error {
	farmer := new(domain.Farmer)
	if err := e.Bind(farmer); err != nil {
		return err
	}
	farmer.CreatedBy = e.Param("username")
	if err := m.AUsecase.AddFarmer(farmer); err != nil {
		return err
	}
	return e.JSON(200, "farmer added successfuly.")
}

//farmer

func (m *UserHandler) FSignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return err
	}
	u, err := m.FUsecase.FSignIn(loginForm.Username, loginForm.Password)
	if err != nil {
		return err
	}
	msg := domain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: u,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) ShowTrees(e echo.Context) error {
	id := e.Param("farmerid")
	t, err := m.FUsecase.ShowTrees(id)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) ShowComments(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("treeid"))
	farmerid, _ := strconv.Atoi(e.Param("farmerid"))
	comment, err := m.FUsecase.ShowComments(farmerid, id)
	if err != nil {
		return err
	}
	return e.JSON(200, comment)
}

func (m *UserHandler) AddTree(e echo.Context) error {
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return err
	}
	if err := m.FUsecase.AddTree(tree); err != nil {
		return err
	}
	return e.JSON(200, "Tree Added successfuly.")
}

func (m *UserHandler) RemoveTree(e echo.Context) error {
	farmerid, _ := strconv.Atoi(e.Param("farmerid"))
	treeid, _ := strconv.Atoi(e.Param("treeid"))
	if err := m.FUsecase.RemoveTree(treeid, farmerid); err != nil {
		return err
	}
	return e.JSON(200, "Tree has been removed.")
}

func (m *UserHandler) AddAttend(e echo.Context) error {
	form := new(domain.AttendForm)
	if err := e.Bind(form); err != nil {
		return err
	}
	if err := m.FUsecase.AddAttend(form); err != nil {
		return err
	}
	return e.JSON(200, "Attend added.")
}
