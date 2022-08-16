package domain

import (
	"garden/internal/domain/comment"
	"garden/internal/domain/garden"
	"garden/internal/domain/gardenLocation"
	"garden/internal/domain/gardenType"
	"garden/internal/domain/service"
	"garden/internal/domain/tag"
	"garden/internal/domain/tree"
	"garden/internal/domain/treeType"
	"garden/internal/domain/user"
	"garden/internal/domain/userType"
)

type UseCases struct {
	User       userDomain.UserUseCase
	UserType   userTypeDomain.UserTypeUseCase
	Tag        tagDomain.TagUseCase
	Garden     gardenDomain.GardenUseCase
	GardenLoc  gardenLocationDomain.GardenLocUseCase
	GardenType gardenTypeDomain.GardenTypeUseCase
	Tree       treeDomain.TreeUseCase
	TreeType   treeTypeDomain.TreeTypeUseCase
	Comment    commentDomain.CommentUseCase
	Service    serviceDomain.ServiceUseCase
}

type Repositories struct {
	User       userDomain.UserRepository
	UserType   userTypeDomain.UserTypeRepository
	Tag        tagDomain.TagRepository
	Garden     gardenDomain.GardenRepository
	GardenLoc  gardenLocationDomain.GardenLocRepository
	GardenType gardenTypeDomain.GardenTypeRepository
	Tree       treeDomain.TreeRepository
	TreeType   treeTypeDomain.TreeTypeRepository
	Comment    commentDomain.CommentRepository
	Service    serviceDomain.ServiceRepository
}
