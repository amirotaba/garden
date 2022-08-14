package main

import (
	"garden/internal/handler/http"
	"garden/internal/utils"
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
	utils.Migrate(Db)

	r := echo.New()

	repo := utils.NewRepository(Db)

	usecase := utils.NewUseCase(repo)

	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	deliver.NewHandler(r, usecase)
}
