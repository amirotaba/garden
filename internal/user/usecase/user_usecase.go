package usecase

import (
	"errors"
	"garden/internal/domain"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type userUsecase struct {
	UserRepo domain.UserRepository

}

func NewUserUsecase(u domain.UserRepository, a domain.AdminRepository, f domain.FarmerRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepo: u,
	}
}

func (a *userUsecase) SignUp (user *domain.User) error {
	if _, err := a.UserRepo.Account(user.UserName); err == nil {
		return errors.New("this username is taken")
	}
	if _, err := a.UserRepo.SignIn("1", user.Email); err == nil {
		return errors.New("this email is used")
	}
	if err := a.UserRepo.SignUp(user); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) SignIn(password, email string) (domain.UserResponse, error) {
	user, err := a.UserRepo.SignIn(password, email)
	if err != nil {
		return domain.UserResponse{}, err
	}
	if user.PassWord != password {
		return domain.UserResponse{}, nil
	}
	jwtsig, errs := jwt.GenerateJWTSigned(user)
	if errs != nil {
		return domain.UserResponse{}, errs
	}
	u := domain.UserResponse{
		UserName: user.UserName,
		Email:    user.Email,
		Token:    jwtsig,
	}
	return u, nil
}

func (a *userUsecase) Account(username string) (domain.UserResponse, error) {
	user, err := a.UserRepo.Account(username)
	if err != nil {
		return domain.UserResponse{}, err
	}
	u := domain.UserResponse{
		UserName: user.UserName,
		Email: user.Email,
	}
	return u, nil
}

func (a *userUsecase) Comment(id int, treeID int, text string) error {
	t, err := a.UserRepo.SearchTree(id)
	if err != nil {
		return err
	}
	t.Comment[id] = text
	if err := a.UserRepo.Comment(t); err != nil {
		return err
	}
	return nil
}

//admin

func (a *userUsecase) ShowGarden() domain.Garden {
	g, err := a.UserRepo.ShowGarden()
	if err != nil {
		return err
	}
	return g
}
func (a *userUsecase) RemoveGarden(id int) error {
	if err := a.UserRepo.RemoveGarden(id); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddGarden(garden domain.Garden) error {
	if err := a.UserRepo.AddGarden(garden); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddFarmer(farmer domain.Farmer) error {
	if err := a.UserRepo.AddFarmer(farmer); err != nil {
		return err
	}
	return nil
}

//farmer

func (a *userUsecase) ShowTrees(id int) ([]string, err) {
	t, err := a.UserRepo.ShowTrees(id)
	if err != nil {
		return err
	}
	return t
}
func (a *userUsecase) ShowComment(id int) (map[int]string, error) {
	m := make(map[int]string)
	t, err := a.UserRepo.SearchTree(id)
	if err != nil {
		return m, err
	}
	return t.Comment, nil
}
func (a *userUsecase) AddTree(tree domain.Tree) error {
	if err := a.UserRepo.AddTree(tree); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) RemoveTree(id int) error {
	if err := a.UserRepo.Remove(tree, id); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddAttend(tree domain.Tree) error {
	if err := a.UserRepo.AddAttend(tree.ID, tree.Attend); err != nil {
		return err
	}
	return nil
}
