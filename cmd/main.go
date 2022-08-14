package main

import (
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

	e := echo.New()

	repos := utils.NewRepository(Db)

	useCases := utils.NewUseCase(repos)

	utils.NewHandler(e, useCases)

	e.Logger.Fatal(e.Start(":4000"))
}
