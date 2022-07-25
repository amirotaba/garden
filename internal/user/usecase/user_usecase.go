package usecase

import (
	"errors"
	"garden/internal/domain"
	"strconv"

	"github.com/majidzarephysics/go-jwt/pkg/jwt"
)

type userUsecase struct {
	//UserRepo  domain.UserRepository
	AdminRepo domain.AdminRepository
}

//func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
//	return &userUsecase{
//		UserRepo: u,
//	}
//}

func NewAdminUsecase(a domain.AdminRepository) domain.AdminUsecase {
	return &userUsecase{
		AdminRepo: a,
	}
}

//func NewFarmerUsecase(f domain.FarmerRepository) domain.FarmerUsecase {
//	return &userUsecase{
//		AdminRepo: f,
//	}
//}

func (a *userUsecase) SignIn(form *domain.LoginForm) (domain.UserResponse, error) {
	user, err := a.AdminRepo.SignIn(form)
	if err != nil {
		return domain.UserResponse{}, err
	}
	if user.PassWord != form.Password {
		return domain.UserResponse{}, errors.New("incorrect password")
	}
	t, err := a.AdminRepo.UserType(form.Type)
	if err != nil {
		return domain.UserResponse{}, err
	}
	jwtsig, errs := jwt.GenerateJWTSigned(user)
	if errs != nil {
		return domain.UserResponse{}, errs
	}
	u := domain.UserResponse{
		UserName: user.UserName,
		Type:     t,
		Token:    jwtsig,
	}
	return u, nil
}

func (a *userUsecase) SignUp(user *domain.User) error {
	if _, err := a.AdminRepo.AccountUser(user.UserName); err == nil {
		return errors.New("this username is taken")
	}
	if err := a.AdminRepo.SignUp(user); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) Account(username string) ([]domain.UserResponse, error) {
	var list []domain.UserResponse
	if username == "" {
		user, err := a.AdminRepo.Account()
		if err != nil {
			return []domain.UserResponse{}, err
		}
		for i := range user {
			u := domain.UserResponse{
				UserName: user[i].UserName,
			}
			list = append(list, u)
		}
		return list, nil
	}
	user, err := a.AdminRepo.AccountUser(username)
	if err != nil {
		return []domain.UserResponse{}, err
	}
	for i := range user {
		u := domain.UserResponse{
			UserName: user[i].UserName,
		}
		list = append(list, u)
	}
	return list, nil
}

func (a *userUsecase) UpdateUser(user *domain.User) error {
	if err := a.AdminRepo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteUser(user *domain.User) error {
	if err := a.AdminRepo.DeleteUser(user.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateUserType(usertype *domain.UserType) error {
	if err := a.AdminRepo.CreateUserType(usertype); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadUserType(id string) ([]domain.UserType, error) {
	if id == "" {
		t, err := a.AdminRepo.ReadUserType()
		if err != nil {
			return []domain.UserType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.AdminRepo.ReadUserTypeID(uint(idInt))
	if err != nil {
		return []domain.UserType{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateUserType(usertype *domain.UserType) error {
	if err := a.AdminRepo.UpdateUserType(usertype); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteUserType(usertype *domain.UserType) error {
	if err := a.AdminRepo.DeleteUserType(usertype.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTag(tag *domain.Tag) error {
	if err := a.AdminRepo.CreateTag(tag); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTag(id string) ([]domain.Tag, error) {
	if id == "" {
		t, err := a.AdminRepo.ReadTag()
		if err != nil {
			return []domain.Tag{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.AdminRepo.ReadTagID(uint(idInt))
	if err != nil {
		return []domain.Tag{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateTag(tag *domain.Tag) error {
	if err := a.AdminRepo.UpdateTag(tag); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTag(tag *domain.Tag) error {
	if err := a.AdminRepo.DeleteTag(tag.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateGarden(garden *domain.Garden) error {
	if err := a.AdminRepo.CreateGarden(garden); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadGarden(id string) ([]domain.Garden, error) {
	if id == "" {
		t, err := a.AdminRepo.ReadGarden()
		if err != nil {
			return []domain.Garden{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.AdminRepo.ReadGardenID(uint(idInt))
	if err != nil {
		return []domain.Garden{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateGarden(garden *domain.Garden) error {
	if err := a.AdminRepo.UpdateGarden(garden); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteGarden(garden *domain.Garden) error {
	if err := a.AdminRepo.DeleteGarden(garden.ID); err != nil {
		return err
	}
	if err := a.AdminRepo.DeleteLocation(garden.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateLocation(location *domain.GardenLocation) error {
	if err := a.AdminRepo.CreateLocation(location); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadLocation(id string) ([]domain.GardenLocation, error) {
	if id == "" {
		t, err := a.AdminRepo.ReadLocation()
		if err != nil {
			return []domain.GardenLocation{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.AdminRepo.ReadLocationID(uint(idInt))
	if err != nil {
		return []domain.GardenLocation{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateLocation(loc *domain.GardenLocation) error {
	if err := a.AdminRepo.UpdateLocation(loc); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteLocation(loc *domain.GardenLocation) error {
	if err := a.AdminRepo.DeleteLocation(loc.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateGardenType(gardenType *domain.GardenType) error {
	if err := a.AdminRepo.CreateGardenType(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadGardenType(id string) ([]domain.GardenType, error) {
	if id == "" {
		t, err := a.AdminRepo.ReadGardenType()
		if err != nil {
			return []domain.GardenType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.AdminRepo.ReadGardenTypeID(uint(idInt))
	if err != nil {
		return []domain.GardenType{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateGardenType(gardenType *domain.GardenType) error {
	if err := a.AdminRepo.UpdateGardenType(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteGardenType(gardenType *domain.GardenType) error {
	if err := a.AdminRepo.DeleteGardenType(gardenType.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTreeType(treeType *domain.TreeType) error {
	if err := a.AdminRepo.CreateTreeType(treeType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTreeType(id string) ([]domain.TreeType, error) {
	if id == "" {
		t, err := a.AdminRepo.ReadTreeType()
		if err != nil {
			return []domain.TreeType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.AdminRepo.ReadTreeTypeID(uint(idInt))
	if err != nil {
		return []domain.TreeType{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateTreeType(treeType *domain.TreeType) error {
	if err := a.AdminRepo.UpdateTreeType(treeType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTreeType(tree *domain.TreeType) error {
	if err := a.AdminRepo.DeleteTreeType(tree.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTree(tree *domain.Tree) error {
	//tree.Qr = make a QRCode
	if err := a.AdminRepo.CreateTree(tree); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTree(m map[string]string) ([]domain.Tree, error) {
	for i := range m {
		if m[i] != "" {
			if i != "type" {
				idInt, err := strconv.Atoi(m[i])
				if err != nil {
					return []domain.Tree{}, err
				}
				q := i + " = ?"
				t, err := a.AdminRepo.ReadTreeID(uint(idInt), q)
				if err != nil {
					return []domain.Tree{}, err
				}
				return t, nil
			}
			t, err := a.AdminRepo.ReadTreeByType(m[i])
			if err != nil {
				return []domain.Tree{}, err
			}
			return t, nil
		}
	}
	tree, err := a.AdminRepo.ReadTree()
	if err != nil {
		return []domain.Tree{}, err
	}
	return tree, nil
}

func (a *userUsecase) UpdateTree(tree *domain.Tree) error {
	if err := a.AdminRepo.UpdateTree(tree); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTree(tree *domain.Tree) error {
	//tIdInt, err := strconv.Atoi(treeid)
	//if err != nil {
	//	return err
	//}
	//fIdInt, err := strconv.Atoi(farmerid)
	//if err != nil {
	//	return err
	//}
	//tIdUint, fIdUint := uint(tIdInt), uint(fIdInt)
	//t, err := a.AdminRepo.SearchTree(tIdUint)
	//if err != nil {
	//	return err
	//}
	//if t.FarmerID != fIdUint {
	//	return errors.New("this tree isn't yours")
	//}
	if err := a.AdminRepo.DeleteTree(tree.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateComment(comment *domain.Comment) error {
	if err := a.AdminRepo.CreateComment(comment); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadComment(m map[string]string) ([]domain.Comment, error) {
	for i := range m {
		if m[i] != "" {
			idInt, err := strconv.Atoi(m[i])
			if err != nil {
				return []domain.Comment{}, err
			}
			q := i + " = ?"
			t, err := a.AdminRepo.ReadCommentID(uint(idInt), q)
			if err != nil {
				return []domain.Comment{}, err
			}
			return t, nil
		}
	}
	c, err := a.AdminRepo.ReadComment()
	if err != nil {
		return []domain.Comment{}, err
	}
	return c, nil
}

func (a *userUsecase) UpdateComment(comment *domain.Comment) error {
	if err := a.AdminRepo.UpdateComment(comment); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteComment(comment *domain.Comment) error {
	if err := a.AdminRepo.DeleteComment(comment.ID); err != nil {
		return err
	}
	return nil
}
