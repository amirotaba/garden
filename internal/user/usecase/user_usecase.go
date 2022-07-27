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

func (a *userUsecase) SignUp(user *domain.User) (int, error) {
	if _, err := a.UserRepo.AccountUser(user.UserName); err == nil {
		return 403, errors.New("this username is taken")
	}
	if err := a.UserRepo.SignUp(user); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) Account(mp map[string]string) ([]domain.UserResponse, error) {
	var list []domain.UserResponse
	uidInt, err := strconv.Atoi(mp["uid"])
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.UserResponse{}, err
	}
	if int(u.Type) > 4 {
		return []domain.UserResponse{}, errors.New("you can't access to this page")
	}
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
	} else if mp["tp"] != "" {
		if int(u.Type) != 1 {
			return []domain.UserResponse{}, errors.New("you can't access to this page")
		}
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
	if int(u.Type) != 1 {
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

func (a *userUsecase) UpdateUser(user *domain.UserForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) == 1 || user.UserName == u.UserName {
		if err := a.UserRepo.UpdateUser(user); err != nil {
			return err
		}
		return nil

	}
	return errors.New("you can't access to this page")
}

func (a *userUsecase) DeleteUser(user *domain.User, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) == 1 || user.UserName == u.UserName {
		if err := a.UserRepo.DeleteUser(user.ID); err != nil {
			return err
		}
		return nil
	}
	return errors.New("you can't access to this page")
}

func (a *userUsecase) CreateUserType(usertype *domain.UserType, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateUserType(usertype); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadUserType(id string, uid string) ([]domain.UserType, error) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.UserType{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return []domain.UserType{}, errors.New("you can't access to this page")
	}
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

func (a *userUsecase) UpdateUserType(usertype *domain.UserTypeForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateUserType(usertype); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteUserType(usertype *domain.UserType, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteUserType(usertype.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTag(tag *domain.Tag, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateTag(tag); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTag(id string, pageNumber string, uid string) ([]domain.Tag, error) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.Tag{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if id == "" {
		if int(u.Type) > 4 {
			return []domain.Tag{}, errors.New("you can't access to this page")
		}
		if pageNumber == "" {
			pageNumber = "1"
		}
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
	if int(u.Type) != 1 {
		return []domain.Tag{}, errors.New("you can't access to this page")
	}
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadTagID(uint(idInt))
	if err != nil {
		return []domain.Tag{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateTag(tag *domain.TagForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateTag(tag); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTag(tag *domain.Tag, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteTag(tag.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateGarden(garden *domain.Garden, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateGarden(garden); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadGarden(mp map[string]string) ([]domain.Garden, error) {
	uidInt, err := strconv.Atoi(mp["uid"])
	if err != nil {
		return []domain.Garden{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 4 {
		return []domain.Garden{}, errors.New("you can't access to this page")
	}
	if mp["id"] != "" {
		idInt, err := strconv.Atoi(mp["id"])
		t, err := a.UserRepo.ReadGardenID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, err
		}
		return t, nil
	} else if mp["userID"] != "" {
		if int(u.Type) > 3 {
			return []domain.Garden{}, errors.New("you can't access to this page")
		}
		idInt, err := strconv.Atoi(mp["userID"])
		t, err := a.UserRepo.ReadGardenUID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, err
		}
		return t, nil
	}
	if int(u.Type) != 1 {
		return []domain.Garden{}, errors.New("you can't access to this page")
	}
	if mp["pageNumber"] == "" {
		mp["pageNumber"] = "1"
	}
	nInt, err := strconv.Atoi(mp["pageNumber"])
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

func (a *userUsecase) UpdateGarden(garden *domain.GardenForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateGarden(garden); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteGarden(garden *domain.Garden, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteGarden(garden.ID); err != nil {
		return err
	}
	if err := a.UserRepo.DeleteLocation(garden.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateLocation(location *domain.GardenLocation, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateLocation(location); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadLocation(gid string, pageNumber string, uid string) ([]domain.GardenLocation, error) {
	if gid == "" {
		uidInt, err := strconv.Atoi(uid)
		if err != nil {
			return []domain.GardenLocation{}, err
		}
		u, err := a.UserRepo.AccountID(uint(uidInt))
		if int(u.Type) > 3 {
			return []domain.GardenLocation{}, errors.New("you can't access to this page")
		}
		if pageNumber == "" {
			pageNumber = "1"
		}
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
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.GardenLocation{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return []domain.GardenLocation{}, errors.New("you can't access to this page")
	}
	idInt, err := strconv.Atoi(gid)
	t, err := a.UserRepo.ReadLocationID(uint(idInt))
	if err != nil {
		return []domain.GardenLocation{}, err
	}
	return t, nil
}

func (a *userUsecase) UpdateLocation(loc *domain.GardenLocationForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateLocation(loc); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteLocation(loc *domain.GardenLocation, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteLocation(loc.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateGardenType(gardenType *domain.GardenType, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateGardenType(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadGardenType(id string, uid string) ([]domain.GardenType, error) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.GardenType{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return []domain.GardenType{}, errors.New("you can't access to this page")
	}
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

func (a *userUsecase) UpdateGardenType(gardenType *domain.GardenTypeForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateGardenType(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteGardenType(gardenType *domain.GardenType, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteGardenType(gardenType.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTree(tree *domain.Tree, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	//tree.Qr = make a QRCode
	if err := a.UserRepo.CreateTree(tree); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTree(m map[string]string, pageNumber, uid string) ([]domain.Tree, error) {
	uidInt, err := strconv.Atoi(m["uid"])
	if err != nil {
		return []domain.Tree{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 4 {
		return []domain.Tree{}, errors.New("you can't access to this page")
	}
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
			if pageNumber == "" {
				pageNumber = "1"
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
	if int(u.Type) != 1 {
		return []domain.Tree{}, errors.New("you can't access to this page")
	}
	if pageNumber == "" {
		pageNumber = "1"
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

func (a *userUsecase) UpdateTree(tree *domain.TreeForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateTree(tree); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTree(tree *domain.Tree, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 3 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteTree(tree.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateTreeType(treeType *domain.TreeType, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateTreeType(treeType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadTreeType(id string, uid string) ([]domain.TreeType, error) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.TreeType{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return []domain.TreeType{}, errors.New("you can't access to this page")
	}
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

func (a *userUsecase) UpdateTreeType(treeType *domain.TreeTypeForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateTreeType(treeType); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) DeleteTreeType(tree *domain.TreeType, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) != 1 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteTreeType(tree.ID); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) CreateComment(comment *domain.Comment, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 4 {
		return errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateComment(comment); err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) ReadComment(m map[string]string, pageNumber, uid string) ([]domain.Comment, error) {
	uidInt, err := strconv.Atoi(m["uid"])
	if err != nil {
		return []domain.Comment{}, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if int(u.Type) > 4 {
		return []domain.Comment{}, errors.New("you can't access to this page")
	}
	for i := range m {
		if m[i] != "" {
			if pageNumber == "" {
				pageNumber = "1"
			}
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
	if int(u.Type) != 1 {
		return []domain.Comment{}, errors.New("you can't access to this page")
	}
	if pageNumber == "" {
		pageNumber = "1"
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

func (a *userUsecase) UpdateComment(comment *domain.CommentForm, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	c, err := a.UserRepo.ReadCommentID(comment.ID, "id", 1)
	if int(u.Type) == 1 || int(c[0].ID) == uidInt {
		if err := a.UserRepo.UpdateComment(comment); err != nil {
			return err
		}
		return nil
	}
	return errors.New("you can't access to this page")
}

func (a *userUsecase) DeleteComment(comment *domain.Comment, uid string) error {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	c, err := a.UserRepo.ReadCommentID(comment.ID, "id", 1)
	if int(u.Type) == 1 || int(c[0].ID) == uidInt {
		if err := a.UserRepo.DeleteComment(comment.ID); err != nil {
			return err
		}
		return nil
	}
	return errors.New("you can't access to this page")
}
