package userUsecase

import (
	"errors"
	"garden/internal/middleware/jwt"
	"strconv"

	"garden/internal/domain"
)

type Usecase struct {
	UserRepo 		domain.UserRepository
	UserTypeRepo	domain.UserTypeRepository
}

func NewUseCase(r domain.UserRepository) domain.UserUseCase {
	return &Usecase{
		UserRepo: r,
	}
}

func (a *Usecase) SignIn(form *domain.LoginForm) (domain.UserResponse, error) {
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
	t.Name, err = a.UserTypeRepo.ReadUser(t.ID)
	if err != nil {
		return domain.UserResponse{}, err
	}

	u := domain.UserResponse{
		UserName: user.UserName,
		Type:     t,
		Token:    jwtsig,
	}

	// do not return status code in response
	return u, nil
}

func (a *Usecase) Create(user domain.User) (domain.UserResponse, error) {
	if _, err := a.UserRepo.ReadUsername(user.UserName); err == nil {
		return domain.UserResponse{}, errors.New("this username is taken")
	}

	if err := a.UserRepo.Create(user); err != nil {
		return domain.UserResponse{}, err
	}

	t, err := a.UserTypeRepo.ReadID(user.ID)
	if err != nil {
		return domain.UserResponse{}, err
	}

	tp := domain.TypeStruct{
		ID:   t.ID,
		Name: t.Name,
	}

	return domain.UserResponse{
		UserName: user.UserName,
		Type:     tp,
	}, nil
}

func (a *Usecase) Read(form domain.AccountForm) ([]domain.UserResponse, error) {
	var list []domain.UserResponse
	if form.Tp != "" {
		tpInt, err := strconv.Atoi(form.Tp)
		if err != nil {
			return []domain.UserResponse{}, err
		}

		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return []domain.UserResponse{}, err
		}
		span := nInt * 10
		user, err := a.UserRepo.ReadByType(span, uint(tpInt))
		for i := range user {
			var t domain.TypeStruct
			t.ID = user[i].Type
			t.Name, err = a.UserTypeRepo.ReadUser(t.ID)
			if err != nil {
				return []domain.UserResponse{}, err
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
		return []domain.UserResponse{}, err
	}
	span := nInt * 10
	user, err := a.UserRepo.Read(span)
	for i := range user {
		var t domain.TypeStruct
		t.ID = user[i].Type
		t.Name, err = a.UserTypeRepo.ReadUser(t.ID)
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

func (a *Usecase) UserRead(form domain.UserAccountForm) (domain.UserResponse, error) {
	if form.Username != "" {
		user, err := a.UserRepo.ReadUsername(form.Username)
		if err != nil {
			return domain.UserResponse{}, err
		}
		u := domain.UserResponse{
			UserName: user.UserName,
		}
		return u, nil
	}
	idInt, err := strconv.Atoi(form.ID)
	if err != nil {
		return domain.UserResponse{}, err
	}
	user, err := a.UserRepo.ReadID(uint(idInt))
	if err != nil {
		return domain.UserResponse{}, err
	}
	o := domain.UserResponse{
		UserName: user.UserName,
	}
	return o, nil
}

func (a *Usecase) Update(user *domain.UserForm, uid uint) error {
	u, err := a.UserRepo.ReadID(uid)
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

func (a *Usecase) Delete(user *domain.User, uid uint) error {
	u, err := a.UserRepo.ReadID(uid)
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
