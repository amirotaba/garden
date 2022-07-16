package usecase

import (
	"errors"
	"garden/internal/domain"
	"strconv"

	"github.com/majidzarephysics/go-jwt/pkg/jwt"
)

type userUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepo: u,
	}
}

func (a *userUsecase) SignUp(user *domain.User) error {
	if _, err := a.UserRepo.Account(user.UserName); err == nil {
		return errors.New("this username is taken")
	}
	if _, err := a.UserRepo.SignInUser("1", user.Email); err == nil {
		return errors.New("this email is used")
	}
	if err := a.UserRepo.SignUp(user); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) SignIn(tp string, password, email string) (domain.UserResponse, error) {
	switch tp {
	case "admin":
		user, err := a.UserRepo.SignInAdmin(password, email)
		if err != nil {
			return domain.UserResponse{}, err
		}
		if user.PassWord != password {
			return domain.UserResponse{}, errors.New("incorrect password")
		}
		jwtsig, errs := jwt.GenerateJWTSigned(user)
		if errs != nil {
			return domain.UserResponse{}, errs
		}
		u := domain.UserResponse{
			UserName: user.UserName,
			Token:    jwtsig,
		}
		return u, nil
	case "user":
		user, err := a.UserRepo.SignInUser(password, email)
		if err != nil {
			return domain.UserResponse{}, err
		}
		if user.PassWord != password {
			return domain.UserResponse{}, errors.New("incorrect password")
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
	case "farmer":
		user, err := a.UserRepo.SignInFarmer(password, email)
		if err != nil {
			return domain.UserResponse{}, err
		}
		if user.PassWord != password {
			return domain.UserResponse{}, errors.New("incorrect password")
		}
		jwtsig, errs := jwt.GenerateJWTSigned(user)
		if errs != nil {
			return domain.UserResponse{}, errs
		}
		u := domain.UserResponse{
			UserName: user.UserName,
			Token:    jwtsig,
		}
		return u, nil
	}
	return domain.UserResponse{}, nil
}

func (a *userUsecase) Account(username string) (domain.UserResponse, error) {
	user, err := a.UserRepo.Account(username)
	if err != nil {
		return domain.UserResponse{}, err
	}
	u := domain.UserResponse{
		UserName: user.UserName,
		Email:    user.Email,
	}
	return u, nil
}

func (a *userUsecase) Comment(id int, treeID int, text string) error {
	t, err := a.UserRepo.SearchTree(treeID)
	if err != nil {
		return err
	}
	t.Comment = t.Comment + strconv.Itoa(id) + text
	if err := a.UserRepo.Comment(t); err != nil {
		return err
	}
	return nil
}

//admin

func (a *userUsecase) ShowGarden() ([]domain.Garden, error) {
	g, err := a.UserRepo.ShowGarden()
	if err != nil {
		return []domain.Garden{}, err
	}
	return g, nil
}
func (a *userUsecase) RemoveGarden(id int) error {
	if err := a.UserRepo.RemoveGarden(id); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddGarden(garden *domain.Garden) error {
	if err := a.UserRepo.AddGarden(garden); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddFarmer(farmer *domain.Farmer) error {
	if err := a.UserRepo.AddFarmer(farmer); err != nil {
		return err
	}
	return nil
}

//farmer

func (a *userUsecase) ShowTrees(id string) ([]domain.Tree, error) {
	idInt, _ := strconv.Atoi(id)
	t, err := a.UserRepo.ShowTrees(idInt)
	if err != nil {
		return []domain.Tree{}, err
	}
	return t, nil
}
func (a *userUsecase) ShowComments(farmerid, id int) (string, error) {
	t, err := a.UserRepo.SearchTree(id)
	if err != nil {
		return "", err
	}
	if t.FarmerID != farmerid {
		return "", errors.New("this tree isn't yours")
	}
	return t.Comment, nil
}
func (a *userUsecase) AddTree(tree *domain.Tree) error {
	if err := a.UserRepo.AddTree(tree); err != nil {
		return err
	}
	t, err := a.UserRepo.LastTree()
	if err != nil {
		return err
	}
	f, err := a.UserRepo.SearchFarmer(tree.FarmerID)
	if err != nil {
		return err
	}
	f.Trees = f.Trees + ", " + strconv.Itoa(int(t.ID))
	if err := a.UserRepo.UpdateFarmer(tree.FarmerID, f.Trees); err != nil {
		return err
	}
	g, err := a.UserRepo.SearchGarden(f.GardenID)
	if err != nil {
		return err
	}
	g.Trees = g.Trees + ", " + strconv.Itoa(int(t.ID))
	if err := a.UserRepo.UpdateGarden(f.GardenID, g.Trees); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) RemoveTree(treeid, farmerid int) error {
	t, err := a.UserRepo.SearchTree(treeid)
	if err != nil {
		return err
	}
	if t.FarmerID != farmerid {
		return errors.New("this tree isn't yours")
	}
	if err := a.UserRepo.RemoveTree(treeid); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddAttend(form *domain.AttendForm) error {
	t, err := a.UserRepo.SearchTree(form.ID)
	if err != nil {
		return err
	}
	t.Attend = t.Attend + ", " + form.Text
	if err := a.UserRepo.AddAttend(t); err != nil {
		return err
	}
	return nil
}
