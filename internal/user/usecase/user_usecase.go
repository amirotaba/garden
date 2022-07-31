package usecase

import (
	"errors"
	"garden/internal/domain"
	"garden/pkg/jwt"
	"strconv"
	"strings"
)

type userUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(a domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepo: a,
	}
}

func (a *userUsecase) SignIn(form *domain.LoginForm) (domain.UserResponse, int, error) {
	user, err := a.UserRepo.SignIn(form)
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	if user.PassWord != form.Password {
		return domain.UserResponse{}, 403, errors.New("incorrect password")
	}
	jwtsig, errs := jwt.GenerateToken(user)
	if errs != nil {
		return domain.UserResponse{}, 400, errs
	}
	var t domain.TypeStruct
	t.ID = user.Type
	t.Name, err = a.UserRepo.UserType(t.ID)
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	u := domain.UserResponse{
		UserName: user.UserName,
		Type:     t,
		Token:    jwtsig,
	}
	return u, 200, nil
}

func (a *userUsecase) SignUp(user *domain.User) (int, error) {
	if _, err := a.UserRepo.AccountUsername(user.UserName); err == nil {
		return 400, errors.New("this username is taken")
	}
	if err := a.UserRepo.SignUp(user); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) Account(mp map[string]string) ([]domain.UserResponse, int, error) {
	var boolean bool
	var list []domain.UserResponse
	uidInt, err := strconv.Atoi(mp["uid"])
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/account")
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}
	List := strings.Split(t[0].AccessList, " ")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.UserResponse{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.UserResponse{}, 403, errors.New("you can't access to this page")
	}
	if mp["tp"] != "" {
		tpInt, err := strconv.Atoi(mp["tp"])
		if err != nil {
			return []domain.UserResponse{}, 400, err
		}
		if mp["pageNumber"] == "" {
			mp["pageNumber"] = "1"
		}
		nInt, err := strconv.Atoi(mp["pageNumber"])
		if err != nil {
			return []domain.UserResponse{}, 400, err
		}
		span := nInt * 10
		user, err := a.UserRepo.AccountType(span, uint(tpInt))
		for i := range user {
			var t domain.TypeStruct
			t.ID = user[i].Type
			t.Name, err = a.UserRepo.UserType(t.ID)
			if err != nil {
				return []domain.UserResponse{}, 400, err
			}
			u := domain.UserResponse{
				UserName: user[i].UserName,
				Type:     t,
			}
			list = append(list, u)
		}
	}
	if mp["pageNumber"] == "" {
		mp["pageNumber"] = "1"
	}
	nInt, err := strconv.Atoi(mp["pageNumber"])
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}
	span := nInt * 10
	user, err := a.UserRepo.Account(span)
	for i := range user {
		var t domain.TypeStruct
		t.ID = user[i].Type
		t.Name, err = a.UserRepo.UserType(t.ID)
		if err != nil {
			return []domain.UserResponse{}, 400, err
		}
		u := domain.UserResponse{
			UserName: user[i].UserName,
			Type:     t,
		}
		list = append(list, u)
	}
	return list, 200, nil
}

func (a *userUsecase) UserAccount(mp map[string]string) (domain.UserResponse, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(mp["uid"])
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/useraccount")
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return domain.UserResponse{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return domain.UserResponse{}, 403, errors.New("you can't access to this page")
	}
	if mp["username"] != "" {
		user, err := a.UserRepo.AccountUsername(mp["username"])
		if err != nil {
			return domain.UserResponse{}, 400, err
		}
		u := domain.UserResponse{
			UserName: user.UserName,
		}
		return u, 200, nil
	}
	idInt, err := strconv.Atoi(mp["id"])
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	user, err := a.UserRepo.AccountID(uint(idInt))
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	o := domain.UserResponse{
		UserName: user.UserName,
	}
	return o, 200, nil
}

func (a *userUsecase) UpdateUser(user *domain.UserForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if boolean || user.UserName == u.UserName {
		if err := a.UserRepo.UpdateUser(user); err != nil {
			return 400, err
		}
		return 200, nil
	}
	return 403, errors.New("you can't access to this page")
}

func (a *userUsecase) DeleteUser(user *domain.User, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if boolean || user.UserName == u.UserName {
		if err := a.UserRepo.DeleteUser(user.ID); err != nil {
			return 400, err
		}
		return 200, nil
	}
	return 403, errors.New("you can't access to this page")
}

func (a *userUsecase) CreateUserType(usertype *domain.UserType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/usertype/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateUserType(usertype); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadUserType(id string, uid string) ([]domain.UserType, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/usertype/read")
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.UserType{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.UserType{}, 403, errors.New("you can't access to this page")
	}
	if id == "" {
		t, err := a.UserRepo.ReadUserType()
		if err != nil {
			return []domain.UserType{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(id)
	tt, err := a.UserRepo.ReadUserTypeID(uint(idInt))
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	return tt, 200, nil
}

func (a *userUsecase) UpdateUserType(usertype *domain.UserTypeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/usertype/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateUserType(usertype); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) UpdateAccess(access *domain.AccessForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/usertype/addaccess")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	List = append(List, strconv.Itoa(int(access.TypeID)))
	out := &domain.UserTypeForm{
		AccessList: strings.Join(List, ""),
		ID:         access.ID,
	}
	if err := a.UserRepo.UpdateUserType(out); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteUserType(usertype *domain.UserType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/usertype/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteUserType(usertype.ID); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) CreateTag(tag *domain.Tag, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tag/create")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateTag(tag); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadTag(pageNumber string, uid string) ([]domain.Tag, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tag/read")
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.Tag{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.Tag{}, 403, errors.New("you can't access to this page")
	}
	if pageNumber == "" {
		pageNumber = "1"
	}
	nInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	span := nInt * 10
	b, err := a.UserRepo.ReadTag(span)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	return b, 200, nil
}

func (a *userUsecase) ReadTagID(id string) ([]domain.Tag, int, error) {
	idInt, err := strconv.Atoi(id)
	t, err := a.UserRepo.ReadTagID(uint(idInt))
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	return t, 200, nil
}

func (a *userUsecase) UpdateTag(tag *domain.TagForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tag/update")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateTag(tag); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteTag(tag *domain.Tag, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tag/delete")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteTag(tag.ID); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) CreateGarden(garden *domain.Garden, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/garden/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateGarden(garden); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadGarden(mp map[string]string) ([]domain.Garden, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(mp["uid"])
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/garden/read")
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.Garden{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.Garden{}, 403, errors.New("you can't access to this page")
	}
	if mp["id"] != "" {
		idInt, err := strconv.Atoi(mp["id"])
		t, err := a.UserRepo.ReadGardenID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, 400, err
		}
		return t, 200, nil
	} else if mp["userID"] != "" {
		idInt, err := strconv.Atoi(mp["userID"])
		t, err := a.UserRepo.ReadGardenUID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, 400, err
		}
		return t, 200, nil
	}
	if mp["pageNumber"] == "" {
		mp["pageNumber"] = "1"
	}
	nInt, err := strconv.Atoi(mp["pageNumber"])
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	span := nInt * 10
	b, err := a.UserRepo.ReadGarden(span)
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	return b, 200, nil
}

func (a *userUsecase) UpdateGarden(garden *domain.GardenForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/garden/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateGarden(garden); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteGarden(garden *domain.Garden, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/garden/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteGarden(garden.ID); err != nil {
		return 400, err
	}
	if err := a.UserRepo.DeleteLocation(garden.ID); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) CreateLocation(location *domain.GardenLocation, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/loc/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateLocation(location); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadLocation(gid string, pageNumber string, uid string) ([]domain.GardenLocation, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/loc/read")
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.GardenLocation{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.GardenLocation{}, 403, errors.New("you can't access to this page")
	}
	if gid == "" {
		if pageNumber == "" {
			pageNumber = "1"
		}
		nInt, err := strconv.Atoi(pageNumber)
		if err != nil {
			return []domain.GardenLocation{}, 400, err
		}
		span := nInt * 10
		t, err := a.UserRepo.ReadLocation(span)
		if err != nil {
			return []domain.GardenLocation{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(gid)
	b, err := a.UserRepo.ReadLocationID(uint(idInt))
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	return b, 200, nil
}

func (a *userUsecase) UpdateLocation(loc *domain.GardenLocationForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/loc/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateLocation(loc); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteLocation(loc *domain.GardenLocation, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/loc/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteLocation(loc.ID); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) CreateGardenType(gardenType *domain.GardenType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/gardentype/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateGardenType(gardenType); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadGardenType(id string, uid string) ([]domain.GardenType, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/gardentype/read")
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.GardenType{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.GardenType{}, 403, errors.New("you can't access to this page")
	}
	idInt, err := strconv.Atoi(id)
	b, err := a.UserRepo.ReadGardenTypeID(uint(idInt))
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	return b, 200, nil
}

func (a *userUsecase) UpdateGardenType(gardenType *domain.GardenTypeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/gardentype/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateGardenType(gardenType); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteGardenType(gardenType *domain.GardenType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/gardentype/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteGardenType(gardenType.ID); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) CreateTree(tree *domain.Tree, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tree/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	//tree.Qr = make a QRCode
	if err := a.UserRepo.CreateTree(tree); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadTreeUser(m map[string]string) ([]domain.Tree, int, error) {
	if m["garden_ID"] != "" {
		idInt, err := strconv.Atoi(m["garden_ID"])
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.UserRepo.ReadTreeID(uint(idInt), "garden_id = ?")
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(m["id"])
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	t, err := a.UserRepo.ReadTreeID(uint(idInt), "id = ?")
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	return t, 200, nil
}

func (a *userUsecase) ReadTree(m map[string]string) ([]domain.Tree, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(m["uid"])
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tree/read")
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.Tree{}, 403, errors.New("you can't access to this page")
	}
	if m["pageNumber"] == "" {
		m["pageNumber"] = "1"
	}
	nInt, err := strconv.Atoi(m["pageNumber"])
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	span := nInt * 10
	if m["garden_ID"] != "" {
		idInt, err := strconv.Atoi(m["garden_ID"])
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.UserRepo.ReadTreeID(uint(idInt), "garden_id = ?")
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	} else if m["type"] != "" {
		idInt, err := strconv.Atoi(m["type"])
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.UserRepo.ReadTreeByType(uint(idInt), span)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	}
	b, err := a.UserRepo.ReadTree(span)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	return b, 200, nil
}

func (a *userUsecase) UpdateTree(tree *domain.TreeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tree/update")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateTree(tree); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteTree(tree *domain.Tree, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tree/delete")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteTree(tree.ID); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) CreateTreeType(treeType *domain.TreeType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/treetype/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateTreeType(treeType); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadTreeType(id string, uid string) ([]domain.TreeType, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/treetype/read")
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.TreeType{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.TreeType{}, 403, errors.New("you can't access to this page")
	}
	if id == "" {
		t, err := a.UserRepo.ReadTreeType()
		if err != nil {
			return []domain.TreeType{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(id)
	b, err := a.UserRepo.ReadTreeTypeID(uint(idInt))
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	return b, 200, nil
}

func (a *userUsecase) UpdateTreeType(treeType *domain.TreeTypeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/treetype/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateTreeType(treeType); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteTreeType(tree *domain.TreeType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/treetype/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteTreeType(tree.ID); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) CreateComment(comment *domain.Comment) (int, error) {
	if err := a.UserRepo.CreateComment(comment); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadComment(m map[string]string, pageNumber, uid string) ([]domain.Comment, int, error) {
	for i := range m {
		if m[i] != "" {
			if pageNumber == "" {
				pageNumber = "1"
			}
			nInt, err := strconv.Atoi(pageNumber)
			if err != nil {
				return []domain.Comment{}, 400, err
			}
			span := nInt * 10
			idInt, err := strconv.Atoi(m[i])
			if err != nil {
				return []domain.Comment{}, 400, err
			}
			q := i + " = ?"
			t, err := a.UserRepo.ReadCommentID(uint(idInt), q, span)
			if err != nil {
				return []domain.Comment{}, 400, err
			}
			return t, 200, nil
		}
	}
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/comment/read")
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.Comment{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.Comment{}, 403, errors.New("you can't access to this page")
	}
	if pageNumber == "" {
		pageNumber = "1"
	}
	nInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	span := nInt * 10
	c, err := a.UserRepo.ReadComment(span)
	if err != nil {
		return []domain.Comment{}, 400, err
	}
	return c, 200, nil
}

func (a *userUsecase) UpdateComment(comment *domain.CommentForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tree/update")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	c, err := a.UserRepo.ReadCommentID(comment.ID, "id", 1)
	if boolean || int(c[0].ID) == uidInt {
		if err := a.UserRepo.UpdateComment(comment); err != nil {
			return 400, err
		}
		return 200, nil
	}
	return 403, errors.New("you can't access to this page")
}

func (a *userUsecase) DeleteComment(comment *domain.Comment, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/tree/update")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	c, err := a.UserRepo.ReadCommentID(comment.ID, "id", 1)
	if boolean || int(c[0].ID) == uidInt {
		if err := a.UserRepo.DeleteComment(comment.ID); err != nil {
			return 400, err
		}
		return 200, nil
	}
	return 403, errors.New("you can't access to this page")
}

func (a *userUsecase) CreateService(service *domain.Service, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/service/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.CreateService(service); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) ReadService(uid string) ([]domain.Service, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.Service{}, 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/service/read")
	if err != nil {
		return []domain.Service{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Service{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.Service{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []domain.Service{}, 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return []domain.Service{}, 403, errors.New("you can't access to this page")
	}
	b, err := a.UserRepo.ReadService()
	if err != nil {
		return []domain.Service{}, 400, err
	}
	return b, 200, nil
}

func (a *userUsecase) UpdateService(service *domain.ServiceForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/service/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.UpdateService(service); err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *userUsecase) DeleteService(service *domain.Service, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.UserRepo.ReadServiceUrl("/user/service/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
	for _, v := range List {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 400, err
		}
		if uint(i) == SID.ID {
			boolean = true
		}
	}
	if !boolean {
		return 403, errors.New("you can't access to this page")
	}
	if err := a.UserRepo.DeleteService(service.ID); err != nil {
		return 400, err
	}
	return 200, nil
}
