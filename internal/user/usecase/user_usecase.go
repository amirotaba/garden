package usecase

import (
	"errors"
	"garden/internal/domain"
	"garden/pkg/jwt"
	"strconv"
)

type userUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(a domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepo: a,
	}
}

func (a *userUsecase) SignIn(form *domain.LoginForm) (domain.UserResponse, error) {
	user, err := a.UserRepo.SignIn(form)
	if err != nil {
		return domain.UserResponse{}, err
	}
	if user.PassWord != form.Password {
		return domain.UserResponse{}, errors.New("incorrect password")
	}
	jwtsig, errs := jwt.GenerateToken(user)
	if errs != nil {
		return domain.UserResponse{}, errs
	}
	var t domain.TypeStruct
	t.ID = user.Type
	t.Name, err = a.UserRepo.UserType(t.ID)
	if err != nil {
		return domain.UserResponse{}, err
	}
	u := domain.UserResponse{
		UserName: user.UserName,
		Type:     t,
		Token:    jwtsig,
	}
	return u, nil
}

func (a *userUsecase) SignUp(user *domain.User) error {
	if _, err := a.UserRepo.AccountUser(user.UserName); err == nil {
		return errors.New("this username is taken")
	}
	if err := a.UserRepo.SignUp(user); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) Account(mp map[string]string) ([]domain.UserResponse, error) {
	var list []domain.UserResponse
	if mp["username"] != "" {
		user, err := a.UserRepo.AccountUser(mp["username"])
		if err != nil {
			return []domain.UserResponse{}, err
		}
		u := domain.UserResponse{
			UserName: user.UserName,
		}
		list = append(list, u)
		return list, nil
	} else if mp["id"] != "" {
		idInt, err := strconv.Atoi(mp["id"])
		if err != nil {
			return []domain.UserResponse{}, err
		}
		user, err := a.UserRepo.AccountID(uint(idInt))
		if err != nil {
			return []domain.UserResponse{}, err
		}
		u := domain.UserResponse{
			UserName: user.UserName,
		}
		list = append(list, u)
		return list, nil
	}
	uidInt, err := strconv.Atoi(mp["uid"])
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.UserResponse{}, err
	} else if mp["tp"] != "" {
		tpInt, err := strconv.Atoi(mp["tp"])
		if err != nil {
			return []domain.UserResponse{}, err
		}
		if mp["pageNumber"] == "" {
			mp["pageNumber"] = "1"
		}
		nInt, err := strconv.Atoi(mp["pageNumber"])
		if err != nil {
			return []domain.UserResponse{}, err
		}
		span := nInt * 10
		user, err := a.UserRepo.AccountType(span, uint(tpInt))
		for i := range user {
			var t domain.TypeStruct
			t.ID = user[i].Type
			t.Name, err = a.UserRepo.UserType(t.ID)
			if err != nil {
				return []domain.UserResponse{}, err
			}
			u := domain.UserResponse{
				UserName: user[i].UserName,
				Type:     t,
			}
			list = append(list, u)
		}
		return list, nil
	}
	if u.Type != uint(1) {
		return []domain.UserResponse{}, errors.New("you can't access to this page")
	}
	if mp["pageNumber"] == "" {
		mp["pageNumber"] = "1"
	}
	nInt, err := strconv.Atoi(mp["pageNumber"])
	if err != nil {
		return []domain.UserResponse{}, err
	}
	span := nInt * 10
	user, err := a.UserRepo.Account(span)
	if err != nil {
		return []domain.UserResponse{}, err
	}
	for i := range user {
		var t domain.TypeStruct
		t.ID = user[i].Type
		t.Name, err = a.UserRepo.UserType(t.ID)
		if err != nil {
			return []domain.UserResponse{}, err
		}
		u := domain.UserResponse{
			UserName: user[i].UserName,
			Type:     t,
		}
		list = append(list, u)
	}
	return list, nil
}

func (a *userUsecase) UpdateUser(user *domain.UserForm) error {
	if err := a.UserRepo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteUser(user *domain.User) error {
	if err := a.UserRepo.DeleteUser(user.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateUserType(usertype *domain.UserType) error {
	if err := a.UserRepo.CreateUserType(usertype); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadUserType(id string) ([]domain.UserType, error) {
	if id == "" {
		t, err := a.UserRepo.ReadUserType()
		if err != nil {
			return []domain.UserType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadUserTypeID(uint(idInt))
	if err != nil {
		return []domain.UserType{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateUserType(usertype *domain.UserTypeForm) error {
	if err := a.UserRepo.UpdateUserType(usertype); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteUserType(usertype *domain.UserType) error {
	if err := a.UserRepo.DeleteUserType(usertype.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTag(tag *domain.Tag) error {
	if err := a.UserRepo.CreateTag(tag); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTag(id string, pageNumber string) ([]domain.Tag, error) {
	if id == "" {
		nInt, err := strconv.Atoi(pageNumber)
		if err != nil {
			return []domain.Tag{}, err
		}
		span := nInt * 10
		t, err := a.UserRepo.ReadTag(span)
		if err != nil {
			return []domain.Tag{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadTagID(uint(idInt))
	if err != nil {
		return []domain.Tag{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateTag(tag *domain.TagForm) error {
	if err := a.UserRepo.UpdateTag(tag); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTag(tag *domain.Tag) error {
	if err := a.UserRepo.DeleteTag(tag.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateGarden(garden *domain.Garden) error {
	if err := a.UserRepo.CreateGarden(garden); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadGarden(id string, pageNumber string) ([]domain.Garden, error) {
	if id == "" {
		nInt, err := strconv.Atoi(pageNumber)
		if err != nil {
			return []domain.Garden{}, err
		}
		span := nInt * 10
		t, err := a.UserRepo.ReadGarden(span)
		if err != nil {
			return []domain.Garden{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadGardenID(uint(idInt))
	if err != nil {
		return []domain.Garden{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateGarden(garden *domain.GardenForm) error {
	if err := a.UserRepo.UpdateGarden(garden); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteGarden(garden *domain.Garden) error {
	if err := a.UserRepo.DeleteGarden(garden.ID); err != nil {
		return err
	}
	if err := a.UserRepo.DeleteLocation(garden.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateLocation(location *domain.GardenLocation) error {
	if err := a.UserRepo.CreateLocation(location); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadLocation(id string, pageNumber string) ([]domain.GardenLocation, error) {
	if id == "" {
		nInt, err := strconv.Atoi(pageNumber)
		if err != nil {
			return []domain.GardenLocation{}, err
		}
		span := nInt * 10
		t, err := a.UserRepo.ReadLocation(span)
		if err != nil {
			return []domain.GardenLocation{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadLocationID(uint(idInt))
	if err != nil {
		return []domain.GardenLocation{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateLocation(loc *domain.GardenLocationForm) error {
	if err := a.UserRepo.UpdateLocation(loc); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteLocation(loc *domain.GardenLocation) error {
	if err := a.UserRepo.DeleteLocation(loc.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateGardenType(gardenType *domain.GardenType) error {
	if err := a.UserRepo.CreateGardenType(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadGardenType(id string) ([]domain.GardenType, error) {
	if id == "" {
		t, err := a.UserRepo.ReadGardenType()
		if err != nil {
			return []domain.GardenType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadGardenTypeID(uint(idInt))
	if err != nil {
		return []domain.GardenType{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateGardenType(gardenType *domain.GardenTypeForm) error {
	if err := a.UserRepo.UpdateGardenType(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteGardenType(gardenType *domain.GardenType) error {
	if err := a.UserRepo.DeleteGardenType(gardenType.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTree(tree *domain.Tree) error {
	//tree.Qr = make a QRCode
	if err := a.UserRepo.CreateTree(tree); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTree(m map[string]string, pageNumber string) ([]domain.Tree, error) {
	for i := range m {
		if m[i] != "" {
			if i != "type" {
				idInt, err := strconv.Atoi(m[i])
				if err != nil {
					return []domain.Tree{}, err
				}
				q := i + " = ?"
				t, err := a.UserRepo.ReadTreeID(uint(idInt), q)
				if err != nil {
					return []domain.Tree{}, err
				}
				return t, nil
			}
			nInt, err := strconv.Atoi(pageNumber)
			if err != nil {
				return []domain.Tree{}, err
			}
			span := nInt * 10
			t, err := a.UserRepo.ReadTreeByType(m[i], span)
			if err != nil {
				return []domain.Tree{}, err
			}
			return t, nil
		}
	}
	nInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return []domain.Tree{}, err
	}
	span := nInt * 10
	tree, err := a.UserRepo.ReadTree(span)
	if err != nil {
		return []domain.Tree{}, err
	}
	return tree, nil
}

func (a *userUsecase) UpdateTree(tree *domain.TreeForm) error {
	if err := a.UserRepo.UpdateTree(tree); err != nil {
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
	//t, err := a.UserRepo.SearchTree(tIdUint)
	//if err != nil {
	//	return err
	//}
	//if t.FarmerID != fIdUint {
	//	return errors.New("this tree isn't yours")
	//}
	if err := a.UserRepo.DeleteTree(tree.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTreeType(treeType *domain.TreeType) error {
	if err := a.UserRepo.CreateTreeType(treeType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTreeType(id string) ([]domain.TreeType, error) {
	if id == "" {
		t, err := a.UserRepo.ReadTreeType()
		if err != nil {
			return []domain.TreeType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadTreeTypeID(uint(idInt))
	if err != nil {
		return []domain.TreeType{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateTreeType(treeType *domain.TreeTypeForm) error {
	if err := a.UserRepo.UpdateTreeType(treeType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTreeType(tree *domain.TreeType) error {
	if err := a.UserRepo.DeleteTreeType(tree.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateComment(comment *domain.Comment) error {
	if err := a.UserRepo.CreateComment(comment); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadComment(m map[string]string, pageNumber string) ([]domain.Comment, error) {
	for i := range m {
		if m[i] != "" {
			nInt, err := strconv.Atoi(pageNumber)
			if err != nil {
				return []domain.Comment{}, err
			}
			span := nInt * 10
			idInt, err := strconv.Atoi(m[i])
			if err != nil {
				return []domain.Comment{}, err
			}
			q := i + " = ?"
			t, err := a.UserRepo.ReadCommentID(uint(idInt), q, span)
			if err != nil {
				return []domain.Comment{}, err
			}
			return t, nil
		}
	}
	nInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return []domain.Comment{}, err
	}
	span := nInt * 10
	c, err := a.UserRepo.ReadComment(span)
	if err != nil {
		return []domain.Comment{}, err
	}
	return c, nil
}

func (a *userUsecase) UpdateComment(comment *domain.CommentForm) error {
	if err := a.UserRepo.UpdateComment(comment); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteComment(comment *domain.Comment) error {
	if err := a.UserRepo.DeleteComment(comment.ID); err != nil {
		return err
	}
	return nil
}
