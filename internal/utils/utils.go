package utils

import (
	"garden/internal/domain"
	comment "garden/internal/features/comment/handler/http"
	"garden/internal/features/comment/repository/mysql"
	"garden/internal/features/comment/usecase"
	garden "garden/internal/features/garden/handler/http"
	"garden/internal/features/garden/repository/mysql"
	"garden/internal/features/garden/usecase"
	service "garden/internal/features/service/handler/http"
	"garden/internal/features/service/repository/mysql"
	"garden/internal/features/service/usecase"
	tag "garden/internal/features/tag/handler/http"
	"garden/internal/features/tag/repository/mysql"
	"garden/internal/features/tag/usecase"
	tree "garden/internal/features/tree/handler/http"
	"garden/internal/features/tree/repository/mysql"
	"garden/internal/features/tree/usecase"
	user "garden/internal/features/user/handler/http"
	"garden/internal/features/user/repository/mysql"
	"garden/internal/features/user/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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
		User:    userRepo.NewMysqlRepository(Db),
		Tag:     tagRepo.NewMysqlRepository(Db),
		Garden:  gardenRepo.NewMysqlRepository(Db),
		Tree:    treeRepo.NewMysqlRepository(Db),
		Comment: commentRepo.NewMysqlRepository(Db),
		Service: serviceRepo.NewMysqlRepository(Db),
	}
	return repository
}

func NewUseCase(repo domain.Repositories) domain.UseCases {
	usecase := domain.UseCases{
		User:    userUsecase.NewUseCase(repo),
		Tag:     tagUsecase.NewUseCase(repo),
		Garden:  gardenUsecase.NewUseCase(repo),
		Tree:    treeUsecase.NewUseCase(repo),
		Comment: commentUsecase.NewUseCase(repo),
		Service: serviceUsecase.NewUseCase(repo),
	}
	return usecase
}

func NewHandler(e *echo.Echo, useCase domain.UseCases) {
	user.NewHandler(e, useCase.User)
	tag.NewHandler(e, useCase.Tag)
	garden.NewHandler(e, useCase.Garden)
	tree.NewHandler(e, useCase.Tree)
	comment.NewHandler(e, useCase.Comment)
	service.NewHandler(e, useCase.Service)
}
