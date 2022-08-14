package deliver

import (
	"garden/internal/domain"
	"garden/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	User    domain.UserUseCase
	Tag     domain.TagUseCase
	Garden  domain.GardenUseCase
	Tree    domain.TreeUseCase
	Comment domain.CommentUseCase
	Service domain.ServiceUseCase
}

func NewHandler(e *echo.Echo, u domain.UseCases) {
	handler := &Handler{
		User:    u.User,
		Tag:     u.Tag,
		Garden:  u.Garden,
		Tree:    u.Tree,
		Comment: u.Comment,
		Service: u.Service,
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

	res.POST("garden/type/create", handler.CreateGardenType)
	res.GET("garden/type/read", handler.ReadGardenType)
	res.PATCH("garden/type/update", handler.UpdateGardenType)
	res.DELETE("garden/type/delete", handler.DeleteGardenType)

	res.POST("tree/create", handler.CreateTree)
	res.GET("tree/read", handler.ReadTree)
	res.GET("tree/readUser", handler.ReadTreeUser)
	res.PATCH("tree/update", handler.UpdateTree)
	res.DELETE("tree/delete", handler.DeleteTree)

	res.POST("tree/type/create", handler.CreateTreeType)
	res.GET("tree/type/read", handler.ReadTreeType)
	res.PATCH("tree/type/update", handler.UpdateTreeType)
	res.DELETE("tree/type/delete", handler.DeleteTreeType)

	res.POST("comment/create", handler.CreateComment)
	res.GET("comment/read", handler.ReadComment)
	res.PATCH("comment/update", handler.UpdateComment)
	res.DELETE("comment/delete", handler.DeleteComment)

	res.POST("service/create", handler.CreateService)
	res.GET("service/read", handler.ReadService)
	res.PATCH("service/update", handler.UpdateService)
	res.DELETE("service/delete", handler.DeleteService)

	handler.addRoutes(e)

	e.Logger.Fatal(e.Start(":4000"))
}
