package usecase

import (
	"errors"
	"garden/internal/domain"
	"strconv"

	"github.com/majidzarephysics/go-jwt/pkg/jwt"
)

type userUsecase struct {
	UserRepo   domain.UserRepository
	AdminRepo  domain.AdminRepository
	FarmerRepo domain.FarmerRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepo: u,
	}
}

func NewAdminUsecase(a domain.AdminRepository) domain.AdminUsecase {
	return &userUsecase{
		AdminRepo: a,
	}
}

func NewFarmerUsecase(f domain.FarmerRepository) domain.FarmerUsecase {
	return &userUsecase{
		FarmerRepo: f,
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

func (a *userUsecase) USignIn(email, password string) (domain.UserResponse, error) {
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
	t.Comment = t.Comment + strconv.Itoa(id) + ": " + text + "\n"
	if err := a.UserRepo.Comment(t); err != nil {
		return err
	}
	return nil
}

//admin

func (a *userUsecase) ASignIn(email, password string) (domain.UserResponse, error) {
	user, err := a.AdminRepo.SignInAdmin(email, password)
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

func (a *userUsecase) ShowGarden() ([]domain.Garden, error) {
	g, err := a.AdminRepo.ShowGarden()
	if err != nil {
		return []domain.Garden{}, err
	}
	return g, nil
}
func (a *userUsecase) RemoveGarden(id string, u string) error {
	idInt, _ := strconv.Atoi(id)
	if err := a.AdminRepo.DeletedBy(idInt, u); err != nil {
		return err
	}
	if err := a.AdminRepo.RemoveGarden(idInt); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddGarden(garden *domain.Garden) error {
	if err := a.AdminRepo.AddGarden(garden); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddFarmer(farmer *domain.Farmer) error {
	if err := a.AdminRepo.AddFarmer(farmer); err != nil {
		return err
	}
	return nil
}

//farmer

func (a *userUsecase) FSignIn(email, password string) (domain.UserResponse, error) {
	user, err := a.FarmerRepo.SignInFarmer(email, password)
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

func (a *userUsecase) ShowTrees(id string) ([]domain.Tree, error) {
	idInt, _ := strconv.Atoi(id)
	t, err := a.FarmerRepo.ShowTrees(idInt)
	if err != nil {
		return []domain.Tree{}, err
	}
	return t, nil
}
func (a *userUsecase) ShowComments(farmerid, id int) (string, error) {
	t, err := a.FarmerRepo.SearchTree(id)
	if err != nil {
		return "", err
	}
	if t.FarmerID != farmerid {
		return "", errors.New("this tree isn't yours")
	}
	return t.Comment, nil
}
func (a *userUsecase) AddTree(tree *domain.Tree) error {
	if err := a.FarmerRepo.AddTree(tree); err != nil {
		return err
	}
	t, err := a.FarmerRepo.LastTree()
	if err != nil {
		return err
	}
	f, err := a.FarmerRepo.SearchFarmer(tree.FarmerID)
	if err != nil {
		return err
	}
	f.Trees = f.Trees + ", " + strconv.Itoa(int(t.ID))
	if err := a.FarmerRepo.UpdateFarmer(tree.FarmerID, f.Trees); err != nil {
		return err
	}
	g, err := a.FarmerRepo.SearchGarden(f.GardenID)
	if err != nil {
		return err
	}
	g.Trees = g.Trees + ", " + strconv.Itoa(int(t.ID))
	if err := a.FarmerRepo.UpdateGarden(f.GardenID, g.Trees); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) RemoveTree(treeid, farmerid int) error {
	t, err := a.FarmerRepo.SearchTree(treeid)
	if err != nil {
		return err
	}
	if t.FarmerID != farmerid {
		return errors.New("this tree isn't yours")
	}
	if err := a.FarmerRepo.RemoveTree(treeid); err != nil {
		return err
	}
	return nil
}
func (a *userUsecase) AddAttend(form *domain.AttendForm) error {
	t, err := a.FarmerRepo.SearchTree(form.ID)
	if err != nil {
		return err
	}
	t.Attend = t.Attend + ", " + form.Text
	if err := a.FarmerRepo.AddAttend(t); err != nil {
		return err
	}
	return nil
}
