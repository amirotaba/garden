package main

import (
	"garden/internal/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	//connect to database
	Db := utils.Connection()

	//migrate tables
	utils.Migrate(Db)

	//start echo
	e := echo.New()

	//get repositories
	repos := utils.NewRepository(Db)

	//get UseCases
	useCases := utils.NewUseCase(repos)

	//register features
	utils.NewHandler(e, useCases)

	//route
	e.Logger.Fatal(e.Start(":4000"))
}
