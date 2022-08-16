package userUsecase

import (
	"errors"
	"garden/internal/domain"
	"garden/internal/domain/user"
	"garden/internal/domain/userType"
	"garden/internal/middleware/jwt"
	"strconv"
)

type Usecase struct {
	UserRepo     userDomain.UserRepository
	UserTypeRepo userTypeDomain.UserTypeRepository
}

func NewUseCase(r domain.Repositories) userDomain.UserUseCase {
	return &Usecase{
		UserRepo:     r.User,
		UserTypeRepo: r.UserType,
	}
}

func (a *Usecase) SignIn(form *userDomain.LoginForm) (userDomain.UserResponse, error) {
	user, err := a.UserRepo.ReadUsername(form.Username)
	if err != nil {
		return userDomain.UserResponse{}, err
	}

	if user.PassWord != form.Password {
		return userDomain.UserResponse{}, errors.New("incorrect password")
	}

	jwtsig, errs := jwt.GenerateToken(user)
	if errs != nil {
		return userDomain.UserResponse{}, errs
	}

	var t userTypeDomain.TypeStruct
	t.ID = user.Type
	tp, err := a.UserTypeRepo.ReadID(t.ID)
	if err != nil {
		return userDomain.UserResponse{}, err
	}
	t.Name = tp.Name

	u := userDomain.UserResponse{
		UserName: user.UserName,
		Type:     t,
		Token:    jwtsig,
	}

	return u, nil
}

func (a *Usecase) Create(user userDomain.User) (userDomain.UserResponse, error) {
	if _, err := a.UserRepo.ReadUsername(user.UserName); err == nil {
		return userDomain.UserResponse{}, errors.New("this username is taken")
	}

	if err := a.UserRepo.Create(user); err != nil {
		return userDomain.UserResponse{}, err
	}

	t, err := a.UserTypeRepo.ReadID(user.ID)
	if err != nil {
		return userDomain.UserResponse{}, err
	}

	tp := userTypeDomain.TypeStruct{
		ID:   t.ID,
		Name: t.Name,
	}

	return userDomain.UserResponse{
		UserName: user.UserName,
		Type:     tp,
	}, nil
}

func (a *Usecase) Read(form userDomain.AccountForm) ([]userDomain.UserResponse, error) {
	var list []userDomain.UserResponse
	if form.Tp != "" {
		tpInt, err := strconv.Atoi(form.Tp)
		if err != nil {
			return nil, err
		}

		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return nil, err
		}
		readForm := userDomain.UserReadForm{
			TypeID: uint(tpInt),
			Span:   nInt * 10,
		}
		user, err := a.UserRepo.ReadByType(readForm)
		for i := range user {
			var t userTypeDomain.TypeStruct
			t.ID = user[i].Type
			tp, err := a.UserTypeRepo.ReadID(t.ID)
			if err != nil {
				return nil, err
			}
			t.Name = tp.Name
			u := userDomain.UserResponse{
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
		return nil, err
	}
	span := nInt * 10
	user, err := a.UserRepo.Read(span)
	for i := range user {
		var t userTypeDomain.TypeStruct
		t.ID = user[i].Type
		tp, err := a.UserTypeRepo.ReadID(t.ID)
		if err != nil {
			return nil, err
		}
		t.Name = tp.Name
		u := userDomain.UserResponse{
			UserName: user[i].UserName,
			Type:     t,
		}
		list = append(list, u)
	}
	return list, nil
}

func (a *Usecase) UserRead(form userDomain.UserAccountForm) (userDomain.UserResponse, error) {
	if form.Username != "" {
		user, err := a.UserRepo.ReadUsername(form.Username)
		if err != nil {
			return userDomain.UserResponse{}, err
		}
		u := userDomain.UserResponse{
			UserName: user.UserName,
		}
		return u, nil
	}
	idInt, err := strconv.Atoi(form.ID)
	if err != nil {
		return userDomain.UserResponse{}, err
	}
	user, err := a.UserRepo.ReadID(uint(idInt))
	if err != nil {
		return userDomain.UserResponse{}, err
	}
	o := userDomain.UserResponse{
		UserName: user.UserName,
	}
	return o, nil
}

func (a *Usecase) Update(user *userDomain.UserForm) error {
	u, err := a.UserRepo.ReadUsername(user.UserName)
	if err != nil {
		return err
	}
	if user.UserName == u.UserName {
		if err := a.UserRepo.Update(user); err != nil {
			return err
		}
	}
	return nil
}

func (a *Usecase) Delete(user *userDomain.User) error {
	u, err := a.UserRepo.ReadUsername(user.UserName)
	if err != nil {
		return err
	}
	if user.UserName == u.UserName {
		if err := a.UserRepo.Delete(user.ID); err != nil {
			return err
		}
	}
	return nil
}
