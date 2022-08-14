package utils

import (
	cRepo "garden/internal/apps/comment/repository/mysql"
	cUsecase "garden/internal/apps/comment/usecase"
	gRepo "garden/internal/apps/garden/repository/mysql"
	gUsecase "garden/internal/apps/garden/usecase"
	sRepo "garden/internal/apps/service/repository/mysql"
	sUsecase "garden/internal/apps/service/usecase"
	tagRepo "garden/internal/apps/tag/repository/mysql"
	tagUsecase "garden/internal/apps/tag/usecase"
	treeRepo "garden/internal/apps/tree/repository/mysql"
	treeUsecase "garden/internal/apps/tree/usecase"
	uRepo "garden/internal/apps/user/repository/mysql"
	uUsecase "garden/internal/apps/user/usecase"
	"garden/internal/domain"
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
	return repo
}

func NewUseCase(repo domain.Repositories) domain.UseCases {
	uu := uUsecase.NewUserUseCase(repo)
	tagu := tagUsecase.NewTagUseCase(repo)
	gu := gUsecase.NewGardenUseCase(repo)
	treeu := treeUsecase.NewTreeUseCase(repo)
	cu := cUsecase.NewCommentUseCase(repo)
	su := sUsecase.NewSerivceUseCase(repo)

	u := domain.UseCases{
		User:    uu,
		Tag:     tagu,
		Garden:  gu,
		Tree:    treeu,
		Comment: cu,
		Service: su,
	}
	return u
}
