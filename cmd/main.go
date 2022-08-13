package main

import (
	cDel "garden/internal/comment/handler"
	"garden/internal/comment/repository/mysql"
	"garden/internal/comment/usecase"
	"garden/internal/domain"
	gDel "garden/internal/garden/handler"
	"garden/internal/garden/repository/mysql"
	"garden/internal/garden/usecase"
	sDel "garden/internal/service/handler"
	"garden/internal/service/repository/mysql"
	"garden/internal/service/usecase"
	tagDel "garden/internal/tag/handler"
	"garden/internal/tag/repository/mysql"
	"garden/internal/tag/usecase"
	treeDel "garden/internal/tree/handler"
	"garden/internal/tree/repository/mysql"
	"garden/internal/tree/usecase"
	"garden/internal/user/delivery/http"
	"garden/internal/user/repository/mysql"
	"garden/internal/user/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dbUser := "root"
	dbPassword := "97216017"
	dbName := "garden"
	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connecting to database failed")
	}
	_ = Db.AutoMigrate(&domain.User{})
	_ = Db.AutoMigrate(&domain.UserType{})
	_ = Db.AutoMigrate(&domain.Garden{})
	_ = Db.AutoMigrate(&domain.GardenType{})
	_ = Db.AutoMigrate(&domain.Tree{})
	_ = Db.AutoMigrate(&domain.TreeType{})
	_ = Db.AutoMigrate(&domain.GardenLocation{})
	_ = Db.AutoMigrate(&domain.Comment{})
	_ = Db.AutoMigrate(&domain.Tag{})
	_ = Db.AutoMigrate(&domain.Service{})

	r := echo.New()

	ur := uRepo.NewMysqlUserRepository(Db)
	tagr := tagRepo.NewMysqlTagRepository(Db)
	gr := gRepo.NewMysqlRepository(Db)
	treer := treeRepo.NewMysqlTreeRepository(Db)
	cr := cRepo.NewMysqlCommentRepository(Db)
	sr := sRepo.NewMysqlSerivceRepository(Db)

	repo := domain.Repositories{
		User:    ur,
		Tag:     tagr,
		Garden:  gr,
		Tree:    treer,
		Comment: cr,
		Service: sr,
	}

	uu := uUsecase.NewUserUseCase(repo)
	tagu := tagUsecase.NewTagUseCase(repo)
	gu := gUsecase.NewGardenUseCase(repo)
	treeu := treeUsecase.NewTreeUseCase(repo)
	cu := cUsecase.NewCommentUseCase(repo)
	su := sUsecase.NewSerivceUseCase(repo)


	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	
	uDel.NewHandler(r, uu)
	tagDel.NewHandler(r, tagu)
	gDel.NewHandler(r, gu)
	treeDel.NewHandler(r, treeu)
	cDel.NewHandler(r, cu)
	sDel.NewHandler(r, su)
}
