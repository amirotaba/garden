package userUsecase

import (
	"errors"
	"garden/internal/pkg/jwt"
	"strconv"
	"strings"

	"garden/internal/domain"
)

type usecase struct {
	UserRepo    domain.UserRepository
	ServiceRepo domain.ServiceRepository
}

func NewUseCase(r domain.Repositories) domain.UserUseCase {
	return &usecase{
		UserRepo:    r.User,
		ServiceRepo: r.Service,
	}
}

func (a *usecase) SignIn(form *domain.LoginForm) (domain.UserResponse, int, error) {
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
	t.Name, err = a.UserRepo.ReadTypeUser(t.ID)
	if err != nil {
		return domain.UserResponse{}, 400, err
	}

	u := domain.UserResponse{
		UserName: user.UserName,
		Type:     t,
		Token:    jwtsig,
	}

	// do not return status code in response
	return u, 200, nil
}

func (a *usecase) Create(user *domain.User) (int, error) {
	if _, err := a.UserRepo.ReadUsername(user.UserName); err == nil {
		return 400, errors.New("this username is taken")
	}

	if err := a.UserRepo.Create(user); err != nil {
		return 400, err
	}

	// do not return status code in response
	return 201, nil
}

func (a *usecase) Read(form domain.AccountForm) ([]domain.UserResponse, int, error) {
	var boolean bool
	var list []domain.UserResponse
	uidInt, err := strconv.Atoi(form.Uid)
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}

	SID, err := a.ServiceRepo.ReadURL("user/account")
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}

	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}

	t, err := a.UserRepo.ReadTypeID(u.Type)
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}

	List := strings.Split(t[0].AccessList, ",")
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

	if form.Tp != "" {
		tpInt, err := strconv.Atoi(form.Tp)
		if err != nil {
			return []domain.UserResponse{}, 400, err
		}

		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.UserResponse{}, 400, err
		}
		span := nInt * 10
		user, err := a.UserRepo.ReadByType(span, uint(tpInt))
		for i := range user {
			var t domain.TypeStruct
			t.ID = user[i].Type
			t.Name, err = a.UserRepo.ReadTypeUser(t.ID)
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
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return []domain.UserResponse{}, 400, err
	}
	span := nInt * 10
	user, err := a.UserRepo.Read(span)
	for i := range user {
		var t domain.TypeStruct
		t.ID = user[i].Type
		t.Name, err = a.UserRepo.ReadTypeUser(t.ID)
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

func (a *usecase) UserRead(form domain.UserAccountForm) (domain.UserResponse, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(form.Uid)
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/userAccount")
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	if form.Username != "" {
		user, err := a.UserRepo.ReadUsername(form.Username)
		if err != nil {
			return domain.UserResponse{}, 400, err
		}
		u := domain.UserResponse{
			UserName: user.UserName,
		}
		return u, 200, nil
	}
	idInt, err := strconv.Atoi(form.ID)
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	user, err := a.UserRepo.ReadID(uint(idInt))
	if err != nil {
		return domain.UserResponse{}, 400, err
	}
	o := domain.UserResponse{
		UserName: user.UserName,
	}
	return o, 200, nil
}

func (a *usecase) Update(user *domain.UserForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
		if err := a.UserRepo.Update(user); err != nil {
			return 400, err
		}
		return 201, nil
	}
	return 403, errors.New("you can't access to this page")
}

func (a *usecase) Delete(user *domain.User, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
		if err := a.UserRepo.Delete(user.ID); err != nil {
			return 400, err
		}
		return 204, nil
	}
	return 403, errors.New("you can't access to this page")
}

func (a *usecase) CreateType(usertype *domain.UserType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/userType/create")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	if err := a.UserRepo.CreateType(usertype); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *usecase) ReadType(id string, uid string) ([]domain.UserType, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/usertype/read")
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
		t, err := a.UserRepo.ReadType()
		if err != nil {
			return []domain.UserType{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(id)
	tt, err := a.UserRepo.ReadTypeID(uint(idInt))
	if err != nil {
		return []domain.UserType{}, 400, err
	}
	return tt, 200, nil
}

func (a *usecase) UpdateType(usertype *domain.UserTypeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/usertype/update")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	if err := a.UserRepo.UpdateType(usertype); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *usecase) UpdateAccess(access *domain.AccessForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/usertype/addAccess")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	AccList := strings.Split(access.TypeID, ",")
	List = append(List, AccList...)
	out := &domain.UserTypeForm{
		AccessList: strings.Join(List, ","),
		ID:         access.ID,
	}
	if err := a.UserRepo.UpdateType(out); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *usecase) DeleteType(usertype *domain.UserType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/usertype/delete")
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	if err := a.UserRepo.DeleteType(usertype.ID); err != nil {
		return 400, err
	}
	return 204, nil
}
