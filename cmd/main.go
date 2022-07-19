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
	_ = Db.AutoMigrate(&domain.Admin{})
	_ = Db.AutoMigrate(&domain.User{})
	_ = Db.AutoMigrate(&domain.Farmer{})
	_ = Db.AutoMigrate(&domain.Garden{})
	_ = Db.AutoMigrate(&domain.Tree{})
	r := echo.New()
	ur := mysqlhandler.NewMysqlUserRepository(Db)
	ar := mysqlhandler.NewMysqlAdminRepository(Db)
	fr := mysqlhandler.NewMysqlFarmerRepository(Db)
	uu := usecase.NewUserUsecase(ur)
	au := usecase.NewAdminUsecase(ar)
	fu := usecase.NewFarmerUsecase(fr)
	httpdelivery.NewUserHandler(r, uu, au, fu)
}
