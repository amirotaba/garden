package userTypeUsecase

import (
	"garden/internal/domain/user"
	"garden/internal/domain/userType"
	"strconv"
	"strings"
)

type Usecase struct {
	UserTypeRepo userTypeDomain.UserTypeRepository
	UserRepo     userDomain.UserRepository
}

func NewUseCase(r userTypeDomain.UserTypeRepository) userTypeDomain.UserTypeUseCase {
	return &Usecase{
		UserTypeRepo: r,
	}
}

func (a *Usecase) Create(usertype *userTypeDomain.UserType) error {
	if err := a.UserTypeRepo.Create(usertype); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(id string) ([]userTypeDomain.UserType, error) {
	list := make([]userTypeDomain.UserType, 0)
	if id == "" {
		t, err := a.UserTypeRepo.Read()
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	tt, err := a.UserTypeRepo.ReadID(uint(idInt))
	if err != nil {
		return nil, err
	}
	list = append(list, tt)
	return list, nil
}

func (a *Usecase) Update(usertype *userTypeDomain.UserTypeForm) error {
	if err := a.UserTypeRepo.Update(usertype); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) UpdateAccess(access *userTypeDomain.AccessForm) error {
	u, err := a.UserRepo.ReadID(access.ID)
	if err != nil {
		return err
	}
	t, err := a.UserTypeRepo.ReadID(u.Type)
	if err != nil {
		return err
	}
	List := strings.Split(t.AccessList, ",")
	AccList := strings.Split(access.TypeID, ",")
	List = append(List, AccList...)
	out := &userTypeDomain.UserTypeForm{
		AccessList: strings.Join(List, ","),
		ID:         access.ID,
	}
	if err := a.UserTypeRepo.Update(out); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(usertype *userTypeDomain.UserType) error {
	if err := a.UserTypeRepo.Delete(usertype.ID); err != nil {
		return err
	}
	return nil
}
