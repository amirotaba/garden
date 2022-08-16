package utils

import (
	"garden/internal/domain"
	"garden/internal/features/comment/handler/http"
	"garden/internal/features/comment/repository/mysql"
	"garden/internal/features/comment/usecase"
	"garden/internal/features/garden/handler/http"
	"garden/internal/features/garden/repository/mysql"
	"garden/internal/features/garden/usecase"
	"garden/internal/features/gardenLocation/handler/http"
	"garden/internal/features/gardenLocation/repository/mysql"
	"garden/internal/features/gardenLocation/usecase"
	"garden/internal/features/gardenType/handler/http"
	"garden/internal/features/gardenType/repository/mysql"
	"garden/internal/features/gardenType/usecase"
	"garden/internal/features/service/handler/http"
	"garden/internal/features/service/repository/mysql"
	"garden/internal/features/service/usecase"
	"garden/internal/features/tag/handler/http"
	"garden/internal/features/tag/repository/mysql"
	"garden/internal/features/tag/usecase"
	"garden/internal/features/tree/handler/http"
	"garden/internal/features/tree/repository/mysql"
	"garden/internal/features/tree/usecase"
	"garden/internal/features/treeType/handler/http"
	"garden/internal/features/treeType/repository/mysql"
	"garden/internal/features/treeType/usecase"
	"garden/internal/features/user/handler/http"
	"garden/internal/features/user/repository/mysql"
	"garden/internal/features/user/usecase"
	"garden/internal/features/userType/handler/http"
	"garden/internal/features/userType/repository/mysql"
	"garden/internal/features/userType/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connection() *gorm.DB {
	dbUser := "root"
	dbPassword := "97216017"
	dbName := "garden"
	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connecting to database failed")
	}
	return Db
}

func Migrate(Db *gorm.DB) {
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
}

func NewRepository(Db *gorm.DB) domain.Repositories {
	repository := domain.Repositories{
		User:       userRepo.NewMysqlRepository(Db),
		UserType:   userTypeRepo.NewMysqlRepository(Db),
		Tag:        tagRepo.NewMysqlRepository(Db),
		Garden:     gardenRepo.NewMysqlRepository(Db),
		GardenLoc:  gardenLocRepo.NewMysqlRepository(Db),
		GardenType: gardenTypeRepo.NewMysqlRepository(Db),
		Tree:       treeRepo.NewMysqlRepository(Db),
		TreeType:   treeTypeRepo.NewMysqlRepository(Db),
		Comment:    commentRepo.NewMysqlRepository(Db),
		Service:    serviceRepo.NewMysqlRepository(Db),
	}
	return repository
}

func NewUseCase(repo domain.Repositories) domain.UseCases {
	usecase := domain.UseCases{
		User:       userUsecase.NewUseCase(repo),
		UserType:   userTypeUsecase.NewUseCase(repo.UserType),
		Tag:        tagUsecase.NewUseCase(repo.Tag),
		Garden:     gardenUsecase.NewUseCase(repo),
		GardenLoc:  gardenLocUsecase.NewUseCase(repo.GardenLoc),
		GardenType: gardenTypeUsecase.NewUseCase(repo.GardenType),
		Tree:       treeUsecase.NewUseCase(repo.Tree),
		TreeType:   treeTypeUsecase.NewUseCase(repo.TreeType),
		Comment:    commentUsecase.NewUseCase(repo.Comment),
		Service:    serviceUsecase.NewUseCase(repo.Service),
	}
	return usecase
}

func NewHandler(e *echo.Echo, useCase domain.UseCases) {
	user.NewHandler(e, useCase.User)
	userType.NewHandler(e, useCase.UserType)
	tag.NewHandler(e, useCase.Tag)
	garden.NewHandler(e, useCase.Garden)
	gardenLoc.NewHandler(e, useCase.GardenLoc)
	gardenType.NewHandler(e, useCase.GardenType)
	tree.NewHandler(e, useCase.Tree)
	treeType.NewHandler(e, useCase.TreeType)
	comment.NewHandler(e, useCase.Comment)
	service.NewHandler(e, useCase.Service)
}
