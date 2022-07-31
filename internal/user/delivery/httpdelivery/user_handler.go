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

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))

	e.POST("signUp", handler.SignUp)
	e.POST("signIn", handler.SignIn)
	res.GET("account", handler.Account)
	res.GET("userAccount", handler.UserAccount)
	res.PATCH("update", handler.UpdateUser)
	res.DELETE("delete", handler.DeleteUser)

	res.POST("usertype/create", handler.CreateUserType)
	res.GET("usertype/read", handler.ReadUserType)
	res.PATCH("usertype/update", handler.UpdateUserType)
	res.PATCH("usertype/addAccess", handler.AddAccess)
	res.DELETE("usertype/delete", handler.DeleteUsertype)

	res.POST("tag/create", handler.CreateTag)
	res.GET("tag/read", handler.ReadTag)
	res.GET("tag/readID", handler.ReadTagID)
	res.PATCH("tag/update", handler.UpdateTag)
	res.DELETE("tag/delete", handler.DeleteTag)

	res.POST("garden/create", handler.CreateGarden)
	res.GET("garden/read", handler.ReadGarden)
	res.PATCH("garden/update", handler.UpdateGarden)
	res.DELETE("garden/delete", handler.DeleteGarden)

	res.POST("loc/create", handler.CreateLocation)
	res.GET("loc/read", handler.ReadLocation)
	res.PATCH("loc/update", handler.UpdateLocation)
	res.DELETE("loc/delete", handler.DeleteLocation)

	res.POST("gardenType/create", handler.CreateGardenType)
	res.GET("gardenType/read", handler.ReadGardenType)
	res.PATCH("gardenType/update", handler.UpdateGardenType)
	res.DELETE("gardenType/delete", handler.DeleteGardenType)

	res.POST("tree/create", handler.CreateTree)
	res.GET("tree/read", handler.ReadTree)
	res.GET("tree/readUser", handler.ReadTreeUser)
	res.PATCH("tree/update", handler.UpdateTree)
	res.DELETE("tree/delete", handler.DeleteTree)

	res.POST("treeType/create", handler.CreateTreeType)
	res.GET("treeType/read", handler.ReadTreeType)
	res.PATCH("treeType/update", handler.UpdateTreeType)
	res.DELETE("treeType/delete", handler.DeleteTreeType)

	res.POST("comment/create", handler.CreateComment)
	res.GET("comment/read", handler.ReadComment)
	res.PATCH("comment/update", handler.UpdateComment)
	res.DELETE("comment/delete", handler.DeleteComment)

	res.POST("service/create", handler.CreateService)
	res.GET("service/read", handler.ReadService)
	res.PATCH("service/update", handler.UpdateService)
	res.DELETE("service/delete", handler.DeleteService)

	e.Logger.Fatal(e.Start(":4000"))
}

func (m *UserHandler) SignIn(e echo.Context) error {
	loginForm := new(domain.LoginForm)
	if err := e.Bind(loginForm); err != nil {
		return e.JSON(403, err.Error())
	}
	u, code, err := m.AUsecase.SignIn(loginForm)
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
	mp["pageNumber"] = e.QueryParam("page")
	users, code, err := m.AUsecase.Account(mp)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, users)
}

func (m *UserHandler) UserAccount(e echo.Context) error {
	mp := make(map[string]string)
	mp["uid"] = strconv.Itoa(int(jwt.UserID(e)))
	mp["username"] = e.QueryParam("username")
	mp["id"] = e.QueryParam("id")
	users, code, err := m.AUsecase.UserAccount(mp)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, users)
}

func (m *UserHandler) UpdateUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.UserForm)
	if err := e.Bind(user); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateUser(user, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User updated successfully")
}

func (m *UserHandler) DeleteUser(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	user := new(domain.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteUser(user, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User deleted successfully")
}

func (m *UserHandler) CreateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type added successfully")
}

func (m *UserHandler) ReadUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.AUsecase.ReadUserType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateUserType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserTypeForm)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *UserHandler) AddAccess(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	add := new(domain.AccessForm)
	if err := e.Bind(add); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateAccess(add, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *UserHandler) DeleteUsertype(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	usertype := new(domain.UserType)
	if err := e.Bind(usertype); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteUserType(usertype, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type deleted successfully")
}

func (m *UserHandler) CreateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag added successfully")
}

func (m *UserHandler) ReadTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	pageNumber := e.QueryParam("page")
	t, code, err := m.AUsecase.ReadTag(pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) ReadTagID(e echo.Context) error {
	id := e.QueryParam("id")
	t, code, err := m.AUsecase.ReadTagID(id)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.TagForm)
	if err := e.Bind(tag); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag updated successfully")
}

func (m *UserHandler) DeleteTag(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tag := new(domain.Tag)
	if err := e.Bind(tag); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteTag(tag, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tag deleted successfully")
}

func (m *UserHandler) CreateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateGarden(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden added successfully.")
}

func (m *UserHandler) ReadGarden(e echo.Context) error {
	mp := make(map[string]string)
	mp["uid"] = strconv.Itoa(int(jwt.UserID(e)))
	mp["userID"] = e.QueryParam("user_id")
	mp["pageNumber"] = e.QueryParam("page")
	mp["id"] = e.QueryParam("id")
	g, code, err := m.AUsecase.ReadGarden(mp)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, g)
}

func (m *UserHandler) UpdateGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.GardenForm)
	if err := e.Bind(garden); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateGarden(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden updated successfully")
}

func (m *UserHandler) DeleteGarden(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	garden := new(domain.Garden)
	if err := e.Bind(garden); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteGarden(garden, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden has been removed.")
}

func (m *UserHandler) CreateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	location := new(domain.GardenLocation)
	if err := e.Bind(location); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateLocation(location, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location added successfully.")
}

func (m *UserHandler) ReadLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gid := e.QueryParam("garden_id")
	pageNumber := e.QueryParam("page")
	t, code, err := m.AUsecase.ReadLocation(gid, pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocationForm)
	if err := e.Bind(loc); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateLocation(loc, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location updated successfully")
}

func (m *UserHandler) DeleteLocation(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	loc := new(domain.GardenLocation)
	if err := e.Bind(loc); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteLocation(loc, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Location deleted successfully")
}

func (m *UserHandler) CreateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type added successfully")
}

func (m *UserHandler) ReadGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.AUsecase.ReadGardenType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenTypeForm)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type updated successfully")
}

func (m *UserHandler) DeleteGardenType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	gardenType := new(domain.GardenType)
	if err := e.Bind(gardenType); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteGardenType(gardenType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Garden type deleted successfully")
}

func (m *UserHandler) CreateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateTree(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree Added successfuly.")
}

func (m *UserHandler) ReadTree(e echo.Context) error {
	mp := make(map[string]string)
	mp["uid"] = strconv.Itoa(int(jwt.UserID(e)))
	mp["garden_ID"] = e.QueryParam("garden_id")
	mp["type"] = e.QueryParam("type")
	mp["pageNumber"] = e.QueryParam("page")
	tree, code, err := m.AUsecase.ReadTree(mp)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, tree)
}

func (m *UserHandler) ReadTreeUser(e echo.Context) error {
	mp := make(map[string]string)
	mp["id"] = e.QueryParam("id")
	mp["garden_ID"] = e.QueryParam("garden_id")
	tree, code, err := m.AUsecase.ReadTreeUser(mp)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, tree)
}

func (m *UserHandler) UpdateTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.TreeForm)
	if err := e.Bind(tree); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateTree(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree updated successfully")
}

func (m *UserHandler) DeleteTree(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	tree := new(domain.Tree)
	if err := e.Bind(tree); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteTree(tree, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree has been removed.")
}

func (m *UserHandler) CreateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateTreeType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type added successfully")
}

func (m *UserHandler) ReadTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	id := e.QueryParam("id")
	t, code, err := m.AUsecase.ReadTreeType(id, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeTypeForm)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateTreeType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type updated successfully")
}

func (m *UserHandler) DeleteTreeType(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	treeType := new(domain.TreeType)
	if err := e.Bind(treeType); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteTreeType(treeType, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Tree type deleted successfully")
}

func (m *UserHandler) CreateComment(e echo.Context) error {
	form := new(domain.Comment)
	if err := e.Bind(form); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateComment(form)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Comment added successfully")
}

func (m *UserHandler) ReadComment(e echo.Context) error {
	mp := make(map[string]string)
	pageNumber := e.QueryParam("page")
	uid := strconv.Itoa(int(jwt.UserID(e)))
	mp["id"] = e.QueryParam("id")
	mp["tree_id"] = e.QueryParam("tree_id")
	mp["tag_id"] = e.QueryParam("tag_id")
	mp["user_id"] = e.QueryParam("user_id")
	comments, code, err := m.AUsecase.ReadComment(mp, pageNumber, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, comments)
}

func (m *UserHandler) UpdateComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.CommentForm)
	if err := e.Bind(comment); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateComment(comment, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "comment updated successfully")
}

func (m *UserHandler) DeleteComment(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	comment := new(domain.Comment)
	if err := e.Bind(comment); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteComment(comment, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Comment has been removed.")
}

func (m *UserHandler) CreateService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	service := new(domain.Service)
	if err := e.Bind(service); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.CreateService(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "Service added successfully")
}

func (m *UserHandler) ReadService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	t, code, err := m.AUsecase.ReadService(uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, t)
}

func (m *UserHandler) UpdateService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	service := new(domain.ServiceForm)
	if err := e.Bind(service); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.UpdateService(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type updated successfully")
}

func (m *UserHandler) DeleteService(e echo.Context) error {
	uid := strconv.Itoa(int(jwt.UserID(e)))
	service := new(domain.Service)
	if err := e.Bind(service); err != nil {
		return e.JSON(400, err.Error())
	}
	code, err := m.AUsecase.DeleteService(service, uid)
	if err != nil {
		return e.JSON(code, err.Error())
	}
	return e.JSON(code, "User type deleted successfully")
}
