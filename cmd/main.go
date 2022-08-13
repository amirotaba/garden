package main

import (
	"garden/internal/domain"
	"garden/internal/user/delivery/http"
	mysql2 "garden/internal/user/repository/mysql"
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

	ur := mysql2.NewMysqlUserRepository(Db)
	tagr := mysql2.NewMysqlTagRepository(Db)
	gr := mysql2.NewMysqlGardenRepository(Db)
	treer := mysql2.NewMysqlTreeRepository(Db)
	cr := mysql2.NewMysqlCommentRepository(Db)
	sr := mysql2.NewMysqlSerivceRepository(Db)

	repo := domain.Repositories{
		User:    ur,
		Tag:     tagr,
		Garden:  gr,
		Tree:    treer,
		Comment: cr,
		Service: sr,
	}

	uu := usecase.NewUserUseCase(repo)
	tagu := usecase.NewTagUseCase(repo)
	gu := usecase.NewGardenUseCase(repo)
	treeu := usecase.NewTreeUseCase(repo)
	cu := usecase.NewCommentUseCase(repo)
	su := usecase.NewSerivceUseCase(repo)

	usecases := domain.UseCases{
		User:    uu,
		Tag:     tagu,
		Garden:  gu,
		Tree:    treeu,
		Comment: cu,
		Service: su,
	}

	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	http.NewHandler(r, usecases)
}
