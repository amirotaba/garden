package httpdelivery

import (
	"garden/internal/domain"
	"garden/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

type UserHandler struct {
	AUsecase domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, au domain.UserUsecase) {
	handler := &UserHandler{
		AUsecase: au,
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("admin/")
	api.Use(middleware.JWTWithConfig(jwt.Config))

	e.POST("user/signup", handler.SignUp)
	e.POST("user/signin", handler.SignIn)
	api.GET("user/account", handler.Account)
	e.POST("user/update", handler.UpdateUser)
	e.POST("user/delete", handler.DeleteUser)

	e.POST("admin/usertype/create", handler.CreateUserType)
	e.GET("admin/usertype/read", handler.ReadUserType)
	e.POST("admin/usertype/update", handler.UpdateUserType)
	e.POST("admin/usertype/delete", handler.DeleteUsertype)

	e.POST("user/tag/create", handler.CreateTag)
	e.GET("user/tag/read", handler.ReadTag)
	e.POST("user/tag/update", handler.UpdateTag)
	e.POST("user/tag/delete", handler.DeleteTag)

	e.POST("user/garden/create", handler.CreateGarden)
	e.GET("user/garden/read", handler.ReadGarden)
	e.POST("user/garden/update", handler.UpdateGarden)
	e.POST("user/garden/delete", handler.DeleteGarden)

	e.POST("user/loc/create", handler.CreateLocation)
	e.GET("user/loc/read", handler.ReadLocation)
	e.POST("user/loc/update", handler.UpdateLocation)
	e.POST("user/loc/delete", handler.DeleteLocation)

	e.POST("admin/gardentype/create", handler.CreateGardenType)
	e.GET("admin/gardentype/read", handler.ReadGardenType)
	e.POST("admin/gardentype/update", handler.UpdateGardenType)
	e.POST("admin/gardentype/delete", handler.DeleteGardentype)

	e.POST("user/tree/create", handler.CreateTree)
	e.GET("user/tree/read", handler.ReadTree)
	e.POST("user/tree/update", handler.UpdateTree)
	e.POST("user/tree/delete", handler.DeleteTree)

	e.POST("admin/treetype/create", handler.CreateTreeType)
	e.GET("admin/treetype/read", handler.ReadTreeType)
	e.POST("admin/treetype/update", handler.UpdateTreeType)
	e.POST("admin/treetype/delete", handler.DeleteTreetype)

	e.POST("user/comment/create", handler.CreateComment)
	e.GET("user/comment/read", handler.ReadComment)
	e.POST("user/comment/update", handler.UpdateTree)
	e.POST("user/comment/delete", handler.DeleteTree)

	e.Logger.Fatal(e.Start(":4000"))
}

func (m *UserHandler) SignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return err
	}
	u, err := m.AUsecase.SignIn(loginForm)
	if err != nil {
		return err
	}
	msg := domain.AuthMessage{
		Text:     "you logged in successfully",
		UserInfo: u,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) SignUp(e echo.Context) error {
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return err
	}
	if err := m.AUsecase.SignUp(user); err != nil {
		return err
	}
	u := domain.UserResponse{UserName: user.UserName}
	msg := &domain.SignUpMessage{
		Text:     "you signed up as, ",
		UserName: u.UserName,
	}
	return e.JSON(200, msg)
}

func (m *UserHandler) Account(e echo.Context) error {
	mp := make(map[string]string)
	mp["uid"] = strconv.Itoa(int(jwt.UserID(e)))
	mp["tp"] = e.QueryParam("type")
	mp["username"] = e.QueryParam("username")
	mp["id"] = e.QueryParam("id")
	mp["pageNumber"] = e.QueryParam("page")
	users, err := m.AUsecase.Account(mp)
	if err != nil {
		return e.JSON(200, err)
	}
	return e.JSON(200, users)
}

func (m *UserHandler) UpdateUser(e echo.Context) error {
	user := new(domain.UserForm)
	if err := e.Bind(user); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateUser(user); err != nil {
		return err
	}
	return e.JSON(200, "User updated successfully")
}

func (m *UserHandler) DeleteUser(e echo.Context) error {
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteUser(user); err != nil {
		return err
	}
	return e.JSON(200, "User deleted successfully")
}

func (m *UserHandler) CreateUserType(e echo.Context) error {
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return err
	}
	if err := m.AUsecase.CreateUserType(usertype); err != nil {
		return err
	}
	return e.JSON(200, "User type added successfully")
}

func (m *UserHandler) ReadUserType(e echo.Context) error {
	id := e.QueryParam("id")
	t, err := m.AUsecase.ReadUserType(id)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateUserType(e echo.Context) error {
	usertype := new(domain.UserTypeForm)
	if err := e.Bind(usertype); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateUserType(usertype); err != nil {
		return err
	}
	return e.JSON(200, "User type updated successfully")
}

func (m *UserHandler) DeleteUsertype(e echo.Context) error {
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteUserType(usertype); err != nil {
		return err
	}
	return e.JSON(200, "User type deleted successfully")
}

func (m *UserHandler) CreateTreeType(e echo.Context) error {
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return err
	}
	if err := m.AUsecase.CreateTreeType(treeType); err != nil {
		return err
	}
	return e.JSON(200, "Tree type added successfully")
}

func (m *UserHandler) ReadTreeType(e echo.Context) error {
	id := e.QueryParam("id")
	t, err := m.AUsecase.ReadTreeType(id)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateTreeType(e echo.Context) error {
	treeType := new(domain.TreeTypeForm)
	if err := e.Bind(treeType); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateTreeType(treeType); err != nil {
		return err
	}
	return e.JSON(200, "Tree type updated successfully")
}

func (m *UserHandler) DeleteTreetype(e echo.Context) error {
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteTreeType(treeType); err != nil {
		return err
	}
	return e.JSON(200, "Tree type deleted successfully")
}

func (m *UserHandler) CreateTag(e echo.Context) error {
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return err
	}
	if err := m.AUsecase.CreateTag(tag); err != nil {
		return err
	}
	return e.JSON(200, "Tag added successfully")
}

func (m *UserHandler) ReadTag(e echo.Context) error {
	id := e.QueryParam("id")
	pageNumber := e.QueryParam("page")
	t, err := m.AUsecase.ReadTag(id, pageNumber)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateTag(e echo.Context) error {
	tag := new(domain.TagForm)
	if err := e.Bind(tag); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateTag(tag); err != nil {
		return err
	}
	return e.JSON(200, "Tag updated successfully")
}

func (m *UserHandler) DeleteTag(e echo.Context) error {
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteTag(tag); err != nil {
		return err
	}
	return e.JSON(200, "Tag deleted successfully")
}

func (m *UserHandler) CreateGarden(e echo.Context) error {
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return err
	}
	if err := m.AUsecase.CreateGarden(garden); err != nil {
		return err
	}
	return e.JSON(200, "Garden added successfuly.")
}

func (m *UserHandler) ReadGarden(e echo.Context) error {
	id := e.QueryParam("user_id")
	pageNumber := e.QueryParam("page")
	g, err := m.AUsecase.ReadGarden(id, pageNumber)
	if err != nil {
		return err
	}
	return e.JSON(200, g)
}

func (m *UserHandler) UpdateGarden(e echo.Context) error {
	garden := new(domain.GardenForm)
	if err := e.Bind(garden); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateGarden(garden); err != nil {
		return err
	}
	return e.JSON(200, "Garden updated successfully")
}

func (m *UserHandler) DeleteGarden(e echo.Context) error {
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteGarden(garden); err != nil {
		return err
	}
	return e.JSON(200, "Garden has been removed.")
}

func (m *UserHandler) CreateLocation(e echo.Context) error {
	location := new(domain.GardenLocation)
	if err := e.Bind(location); err != nil {
		return err
	}
	if err := m.AUsecase.CreateLocation(location); err != nil {
		return err
	}
	return e.JSON(200, "Location added successfuly.")
}

func (m *UserHandler) ReadLocation(e echo.Context) error {
	id := e.QueryParam("id")
	pageNumber := e.QueryParam("page")
	t, err := m.AUsecase.ReadLocation(id, pageNumber)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateLocation(e echo.Context) error {
	loc := new(domain.GardenLocationForm)
	if err := e.Bind(loc); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateLocation(loc); err != nil {
		return err
	}
	return e.JSON(200, "Location updated successfully")
}

func (m *UserHandler) DeleteLocation(e echo.Context) error {
	loc := new(domain.GardenLocation)
	if err := e.Bind(loc); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteLocation(loc); err != nil {
		return err
	}
	return e.JSON(200, "Location deleted successfully")
}

func (m *UserHandler) CreateGardenType(e echo.Context) error {
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return err
	}
	if err := m.AUsecase.CreateGardenType(gardenType); err != nil {
		return err
	}
	return e.JSON(200, "Garden type added successfully")
}

func (m *UserHandler) ReadGardenType(e echo.Context) error {
	id := e.QueryParam("id")
	t, err := m.AUsecase.ReadGardenType(id)
	if err != nil {
		return err
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateGardenType(e echo.Context) error {
	gardenType := new(domain.GardenTypeForm)
	if err := e.Bind(gardenType); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateGardenType(gardenType); err != nil {
		return err
	}
	return e.JSON(200, "Tree type updated successfully")
}

func (m *UserHandler) DeleteGardentype(e echo.Context) error {
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteGardenType(gardenType); err != nil {
		return err
	}
	return e.JSON(200, "Garden type deleted successfully")
}

func (m *UserHandler) CreateTree(e echo.Context) error {
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return err
	}
	if err := m.AUsecase.CreateTree(tree); err != nil {
		return err
	}
	return e.JSON(200, "Tree Added successfuly.")
}

func (m *UserHandler) ReadTree(e echo.Context) error {
	mp := make(map[string]string)
	pageNumber := e.QueryParam("page")
	mp["id"] = e.QueryParam("id")
	mp["type"] = e.QueryParam("type")
	mp["garden_ID"] = e.QueryParam("garden_id")

	tree, err := m.AUsecase.ReadTree(mp, pageNumber)
	if err != nil {
		return err
	}
	return e.JSON(200, tree)
}

func (m *UserHandler) UpdateTree(e echo.Context) error {
	tree := new(domain.TreeForm)
	if err := e.Bind(tree); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateTree(tree); err != nil {
		return err
	}
	return e.JSON(200, "Tree updated successfully")
}

func (m *UserHandler) DeleteTree(e echo.Context) error {
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteTree(tree); err != nil {
		return err
	}
	return e.JSON(200, "Tree has been removed.")
}

func (m *UserHandler) CreateComment(e echo.Context) error {
	form := new(domain.Comment)
	// user_id :=
	if err := e.Bind(form); err != nil {
		return err
	}
	if err := m.AUsecase.CreateComment(form); err != nil {
		return err
	}
	return e.JSON(200, "Comment added successfully")
}

func (m *UserHandler) ReadComment(e echo.Context) error {
	mp := make(map[string]string)
	pageNumber := e.QueryParam("page")
	mp["id"] = e.QueryParam("id")
	mp["tree_id"] = e.QueryParam("tree_id")
	mp["tag_id"] = e.QueryParam("tag_id")
	mp["user_id"] = e.QueryParam("user_id")
	comments, err := m.AUsecase.ReadComment(mp, pageNumber)
	if err != nil {
		return err
	}
	return e.JSON(200, comments)
}

func (m *UserHandler) UpdateComment(e echo.Context) error {
	comment := new(domain.CommentForm)
	if err := e.Bind(comment); err != nil {
		return err
	}
	if err := m.AUsecase.UpdateComment(comment); err != nil {
		return err
	}
	return e.JSON(200, "comment updated successfully")
}

func (m *UserHandler) DeleteComment(e echo.Context) error {
	comment := new(domain.Comment)
	if err := e.Bind(comment); err != nil {
		return err
	}
	if err := m.AUsecase.DeleteComment(comment); err != nil {
		return err
	}
	return e.JSON(200, "CreateComment has been removed.")
}
