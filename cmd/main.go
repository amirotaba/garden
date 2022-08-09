package main

import (
	"garden/internal/domain"
	"garden/internal/user/delivery/httpdelivery"
	"garden/internal/user/repository/mysqlhandler"
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
	ar := mysqlhandler.NewMysqlUserRepository(Db)
	au := usecase.NewUserUseCase(ar)
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	httpdelivery.NewUserHandler(r, au)
}
