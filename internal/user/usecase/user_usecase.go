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

func (a *userUsecase) Comment(comment *domain.Comment, user_id string) error {
	if err := a.UserRepo.Comment(comment); err != nil {
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

func (a *userUsecase) AddGarden(garden *domain.Garden, user_id string) error {
	idInt, err := strconv.Atoi(user_id)
	if err != nil {
		return err
	}
	garden.UserId = uint(idInt)
	if err := a.AdminRepo.AddGarden(garden); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) AddLocation(location *domain.GardenLocation, user_id string) error {
	idInt, err := strconv.Atoi(user_id)
	if err != nil {
		return err
	}
	location.UserId = uint(idInt)
	if err := a.AdminRepo.AddLocation(location); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) AddFarmer(farmer *domain.Farmer, user_id string) error {
	idInt, err := strconv.Atoi(user_id)
	if err != nil {
		return err
	}
	farmer.UserId = uint(idInt)
	if err := a.AdminRepo.AddFarmer(farmer); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ShowGarden() ([]domain.Garden, error) {
	g, err := a.AdminRepo.ShowGarden()
	if err != nil {
		return []domain.Garden{}, err
	}
	return g, nil
}

func (a *userUsecase) RemoveGarden(id string, user_id string) error {
	garId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	garIdUint := uint(garId)
	idInt, err := strconv.Atoi(user_id)
	if err != nil {
		return err
	}
	idUint := uint(idInt)
	if err := a.AdminRepo.DeletedBy(garIdUint, idUint); err != nil {
		return err
	}
	if err := a.AdminRepo.RemoveGarden(garIdUint); err != nil {
		return err
	}
	if err := a.AdminRepo.RemoveGardenLocation(garIdUint); err != nil {
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

func (a *userUsecase) AddTree(tree *domain.Tree, user_id string) error {
	//tree.Qr = make a QRCode
	idInt, err := strconv.Atoi(user_id)
	if err != nil {
		return err
	}
	tree.FarmerID = uint(idInt)
	if err := a.FarmerRepo.AddTree(tree); err != nil {
		return err
	}
	f, err := a.FarmerRepo.SearchFarmer(tree.FarmerID)
	if err != nil {
		return err
	}
	g, err := a.FarmerRepo.SearchGarden(f.GardenID)
	if err != nil {
		return err
	}
	g.TreesCount += 1
	if err := a.FarmerRepo.UpdateGarden(f.GardenID, g.TreesCount); err != nil {
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

func (a *userUsecase) ShowTrees(id string) ([]domain.Tree, error) {
	idInt, _ := strconv.Atoi(id)
	t, err := a.FarmerRepo.ShowTrees(uint(idInt))
	if err != nil {
		return []domain.Tree{}, err
	}
	return t, nil
}

func (a *userUsecase) ShowComments(farmerid, id string) ([]domain.Comment, error) {
	fidInt, err := strconv.Atoi(farmerid)
	if err != nil {
		return []domain.Comment{}, err
	}
	tidInt, err := strconv.Atoi(id)
	if err != nil {
		return []domain.Comment{}, err
	}
	t, err := a.FarmerRepo.SearchTree(uint(tidInt))
	if err != nil {
		return []domain.Comment{}, err
	}
	if t.FarmerID != uint(fidInt) {
		return []domain.Comment{}, errors.New("this tree isn't yours")
	}
	c, err := a.FarmerRepo.SearchComment(uint(tidInt))
	return c, nil
}

func (a *userUsecase) RemoveTree(treeid, farmerid string) error {
	tIdInt, err := strconv.Atoi(treeid)
	if err != nil {
		return err
	}
	fIdInt, err := strconv.Atoi(farmerid)
	if err != nil {
		return err
	}
	tIdUint, fIdUint := uint(tIdInt), uint(fIdInt)
	t, err := a.FarmerRepo.SearchTree(tIdUint)
	if err != nil {
		return err
	}
	if t.FarmerID != fIdUint {
		return errors.New("this tree isn't yours")
	}
	if err := a.FarmerRepo.RemoveTree(tIdUint); err != nil {
		return err
	}
	return nil
	g, err := a.FarmerRepo.SearchGarden(t.GardenId)
	if err != nil {
		return err
	}
	g.TreesCount -= 1
	if err := a.FarmerRepo.UpdateGarden(t.GardenId, g.TreesCount); err != nil {
		return err
	}
	return nil
}
