package httpdelivery

import (
	"garden/internal/domain"
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
	e.POST("user/addcmt/:id", handler.Comment)

	e.POST("usertype/create", handler.CreateUserType)
	e.GET("usertype/read", handler.ReadUserType)
	e.POST("usertype/update", handler.UpdateUserType)
	e.POST("usertype/delete", handler.Deleteusertype)

	e.POST("admin/signin", handler.ASignIn)
	e.POST("admin/grd/add/:id", handler.AddGarden)
	e.POST("admin/grd/add/loc/:id", handler.AddLocation)
	e.POST("admin/frm/add/:id", handler.AddFarmer)
	e.GET("admin/grd/shw", handler.ShowGarden)
	e.GET("admin/grd/rmv/:aid/:id", handler.RemoveGarden)

	e.POST("farmer/signin", handler.FSignIn)
	e.GET("farmer/tree/shw/:farmerid", handler.ShowTrees)
	e.GET("farmer/cmt/shw/:farmerid/:treeid", handler.ShowComments)
	e.POST("farmer/tree/add/:farmerid", handler.AddTree)
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
	form := new(domain.Comment)
	user_id := e.Param("id")
	if err := e.Bind(form); err != nil {
		return err
	}
	if err := m.UUsecase.Comment(form, user_id); err != nil {
		return err
	}
	return e.JSON(200, "Comment saved")
}

func (m *UserHandler) CreateUserType(e echo.Context) error {
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return err
	}
	if err := m.AUsecase.CreateUserType(usertype); err != nil {
		return err
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) ReadUserType(e echo.Context) error {
	t, err := m.AUsecase.ReadUserType()
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateUserType(e echo.Context) error {
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateUserType(usertype); err != nil {
		return err
	}
	return e.JSON(200, "user updated successfuly")
}

func (m *UserHandler) Deleteusertype(e echo.Context) error {
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return err
	}
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
	aid := e.Param("aid")
	id := e.Param("id")
	if err := m.AUsecase.RemoveGarden(id, aid); err != nil {
		return err
	}
	return e.JSON(200, "Garden has been removed.")
}

func (m *UserHandler) AddGarden(e echo.Context) error {
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return err
	}
	user_id := e.Param("id")
	if err := m.AUsecase.AddGarden(garden, user_id); err != nil {
		return err
	}
	return e.JSON(200, "Garden added successfuly.")
}

func (m *UserHandler) AddFarmer(e echo.Context) error {
	farmer := new(domain.Farmer)
	if err := e.Bind(farmer); err != nil {
		return err
	}
	user_id := e.Param("id")
	if err := m.AUsecase.AddFarmer(farmer, user_id); err != nil {
		return err
	}
	return e.JSON(200, "farmer added successfuly.")
}

func (m *UserHandler) AddLocation(e echo.Context) error {
	user_id := e.Param("id")
	location := new(domain.GardenLocation)
	if err := e.Bind(location); err != nil {
		return err
	}
	if err := m.AUsecase.AddLocation(location, user_id); err != nil {
		return err
	}
	return e.JSON(200, "Location added successfuly.")
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

func (m *UserHandler) AddTree(e echo.Context) error {
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return err
	}
	user_id := e.Param("farmerid")
	if err := m.FUsecase.AddTree(tree, user_id); err != nil {
		return err
	}
	return e.JSON(200, "Tree Added successfuly.")
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

func (m *UserHandler) ShowTrees(e echo.Context) error {
	id := e.Param("farmerid")
	t, err := m.FUsecase.ShowTrees(id)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) ShowComments(e echo.Context) error {
	id := e.Param("treeid")
	farmerid := e.Param("farmerid")
	comment, err := m.FUsecase.ShowComments(farmerid, id)
	if err != nil {
		return err
	}
	return e.JSON(200, comment)
}

func (m *UserHandler) RemoveTree(e echo.Context) error {
	farmerid := e.Param("farmerid")
	treeid := e.Param("treeid")
	if err := m.FUsecase.RemoveTree(treeid, farmerid); err != nil {
		return err
	}
	return e.JSON(200, "Tree has been removed.")
}
