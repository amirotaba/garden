package userTypeUsecase

import (
	"strconv"
	"strings"

	"garden/internal/domain"
)

type Usecase struct {
	UserTypeRepo 	domain.UserTypeRepository
	UserRepo 		domain.UserRepository

}

func NewUseCase(r domain.UserTypeRepository) domain.UserTypeUseCase {
	return &Usecase{
		UserTypeRepo: r,
	}
}

func (a *Usecase) Create(usertype *domain.UserType) error {
	if err := a.UserTypeRepo.Create(usertype); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(id string) ([]domain.UserType, error) {
	list := make([]domain.UserType, 0)
	if id == "" {
		t, err := a.UserTypeRepo.Read()
		if err != nil {
			return []domain.UserType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	tt, err := a.UserTypeRepo.ReadID(uint(idInt))
	if err != nil {
		return []domain.UserType{}, err
	}
	list = append(list, tt)
	return list, nil
}

func (a *Usecase) Update(usertype *domain.UserTypeForm) error {
	if err := a.UserTypeRepo.Update(usertype); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) UpdateAccess(access *domain.AccessForm, uid uint) error {
	u, err := a.UserRepo.ReadID(uid)
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
	out := &domain.UserTypeForm{
		AccessList: strings.Join(List, ","),
		ID:         access.ID,
	}
	if err := a.UserTypeRepo.Update(out); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(usertype *domain.UserType) error {
	if err := a.UserTypeRepo.Delete(usertype.ID); err != nil {
		return err
	}
	return nil
}
