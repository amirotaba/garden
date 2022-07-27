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

	res := e.Group("res/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	e.POST("user/signup", handler.SignUp)
	e.POST("user/signin", handler.SignIn)
	res.GET("user/account", handler.Account)
	res.POST("user/update", handler.UpdateUser)
	res.POST("user/delete", handler.DeleteUser)

	res.POST("admin/usertype/create", handler.CreateUserType)
	res.GET("admin/usertype/read", handler.ReadUserType)
	res.POST("admin/usertype/update", handler.UpdateUserType)
	res.POST("admin/usertype/delete", handler.DeleteUsertype)

	res.POST("user/tag/create", handler.CreateTag)
	res.GET("user/tag/read", handler.ReadTag)
	res.POST("user/tag/update", handler.UpdateTag)
	res.POST("user/tag/delete", handler.DeleteTag)

	res.POST("user/garden/create", handler.CreateGarden)
	res.GET("user/garden/read", handler.ReadGarden)
	res.POST("user/garden/update", handler.UpdateGarden)
	res.POST("user/garden/delete", handler.DeleteGarden)

	res.POST("user/loc/create", handler.CreateLocation)
	res.GET("user/loc/read", handler.ReadLocation)
	res.POST("user/loc/update", handler.UpdateLocation)
	res.POST("user/loc/delete", handler.DeleteLocation)

	res.POST("admin/gardentype/create", handler.CreateGardenType)
	res.GET("admin/gardentype/read", handler.ReadGardenType)
	res.POST("admin/gardentype/update", handler.UpdateGardenType)
	res.POST("admin/gardentype/delete", handler.DeleteGardentype)

	res.POST("user/tree/create", handler.CreateTree)
	res.GET("user/tree/read", handler.ReadTree)
	res.POST("user/tree/update", handler.UpdateTree)
	res.POST("user/tree/delete", handler.DeleteTree)

	res.POST("admin/treetype/create", handler.CreateTreeType)
	res.GET("admin/treetype/read", handler.ReadTreeType)
	res.POST("admin/treetype/update", handler.UpdateTreeType)
	res.POST("admin/treetype/delete", handler.DeleteTreetype)

	res.POST("user/comment/create", handler.CreateComment)
	res.GET("user/comment/read", handler.ReadComment)
	res.POST("user/comment/update", handler.UpdateComment)
	res.POST("user/comment/delete", handler.DeleteComment)

	e.Logger.Fatal(e.Start(":4000"))
}

func (m *UserHandler) SignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return e.JSON(403, err.Error())
	}
	u, err := m.AUsecase.SignIn(loginForm)
	if err != nil {
		return e.JSON(403, err.Error())
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
		return e.JSON(403, err.Error())
	}
	code, err := m.AUsecase.SignUp(user)
	if err != nil {
		return e.JSON(code, err.Error())
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
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, users)
}

func (m *UserHandler) UpdateUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.UserForm)
	if err := e.Bind(user); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateUser(user, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "User updated successfully")
}

func (m *UserHandler) DeleteUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteUser(user, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "User deleted successfully")
}

func (m *UserHandler) CreateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return err
	}
	if err := m.AUsecase.CreateUserType(usertype, uid); err != nil {
		return err
	}
	return e.JSON(200, "User type added successfully")
}

func (m *UserHandler) ReadUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, err := m.AUsecase.ReadUserType(id, uid)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserTypeForm)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateUserType(usertype, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "User type updated successfully")
}

func (m *UserHandler) DeleteUsertype(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteUserType(usertype, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "User type deleted successfully")
}

func (m *UserHandler) CreateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.CreateTag(tag, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tag added successfully")
}

func (m *UserHandler) ReadTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	pageNumber := e.QueryParam("page")
	t, err := m.AUsecase.ReadTag(id, pageNumber, uid)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.TagForm)
	if err := e.Bind(tag); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateTag(tag, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tag updated successfully")
}

func (m *UserHandler) DeleteTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteTag(tag, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tag deleted successfully")
}

func (m *UserHandler) CreateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.CreateGarden(garden, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Garden added successfuly.")
}

func (m *UserHandler) ReadGarden(e echo.Context) error {
	mp := make(map[string]string)
	mp["uid"] = strconv.Itoa(int(jwt.UserID(e)))
	mp["userID"] = e.QueryParam("user_id")
	mp["pageNumber"] = e.QueryParam("page")
	mp["id"] = e.QueryParam("id")
	g, err := m.AUsecase.ReadGarden(mp)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, g)
}

func (m *UserHandler) UpdateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.GardenForm)
	if err := e.Bind(garden); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateGarden(garden, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Garden updated successfully")
}

func (m *UserHandler) DeleteGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteGarden(garden, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Garden has been removed.")
}

func (m *UserHandler) CreateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	location := new(domain.GardenLocation)
	if err := e.Bind(location); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.CreateLocation(location, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Location added successfuly.")
}

func (m *UserHandler) ReadLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gid := e.QueryParam("garden_id")
	pageNumber := e.QueryParam("page")
	t, err := m.AUsecase.ReadLocation(gid, pageNumber, uid)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocationForm)
	if err := e.Bind(loc); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateLocation(loc, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Location updated successfully")
}

func (m *UserHandler) DeleteLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocation)
	if err := e.Bind(loc); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteLocation(loc, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Location deleted successfully")
}

func (m *UserHandler) CreateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.CreateGardenType(gardenType, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Garden type added successfully")
}

func (m *UserHandler) ReadGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, err := m.AUsecase.ReadGardenType(id, uid)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenTypeForm)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateGardenType(gardenType, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tree type updated successfully")
}

func (m *UserHandler) DeleteGardentype(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteGardenType(gardenType, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Garden type deleted successfully")
}

func (m *UserHandler) CreateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.CreateTree(tree, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tree Added successfuly.")
}

func (m *UserHandler) ReadTree(e echo.Context) error {
	mp := make(map[string]string)
	pageNumber := e.QueryParam("page")
	uid := strconv.Itoa(int(jwt.UserID(e)))
	mp["id"] = e.QueryParam("id")
	mp["type"] = e.QueryParam("type")
	mp["garden_ID"] = e.QueryParam("garden_id")

	tree, err := m.AUsecase.ReadTree(mp, pageNumber, uid)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, tree)
}

func (m *UserHandler) UpdateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.TreeForm)
	if err := e.Bind(tree); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateTree(tree, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tree updated successfully")
}

func (m *UserHandler) DeleteTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteTree(tree, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tree has been removed.")
}

func (m *UserHandler) CreateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.CreateTreeType(treeType, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tree type added successfully")
}

func (m *UserHandler) ReadTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, err := m.AUsecase.ReadTreeType(id, uid)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, t)
}

func (m *UserHandler) UpdateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeTypeForm)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateTreeType(treeType, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tree type updated successfully")
}

func (m *UserHandler) DeleteTreetype(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteTreeType(treeType, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Tree type deleted successfully")
}

func (m *UserHandler) CreateComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	form := new(domain.Comment)
	if err := e.Bind(form); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.CreateComment(form, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Comment added successfully")
}

func (m *UserHandler) ReadComment(e echo.Context) error {
	mp := make(map[string]string)
	pageNumber := e.QueryParam("page")
	uid := strconv.Itoa(int(jwt.UserID(e)))
	mp["id"] = e.QueryParam("id")
	mp["tree_id"] = e.QueryParam("tree_id")
	mp["tag_id"] = e.QueryParam("tag_id")
	mp["user_id"] = e.QueryParam("user_id")
	comments, err := m.AUsecase.ReadComment(mp, pageNumber, uid)
	if err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, comments)
}

func (m *UserHandler) UpdateComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.CommentForm)
	if err := e.Bind(comment); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.UpdateComment(comment, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "comment updated successfully")
}

func (m *UserHandler) DeleteComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.Comment)
	if err := e.Bind(comment); err != nil {
		return e.JSON(403, err.Error())
	}
	if err := m.AUsecase.DeleteComment(comment, uid); err != nil {
		return e.JSON(403, err.Error())
	}
	return e.JSON(200, "Comment has been removed.")
}
