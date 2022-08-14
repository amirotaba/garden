package serviceUsecase

import (
	"errors"
	"strconv"
	"strings"

	"garden/internal/domain"
)

type ServiceUsecase struct {
	UserRepo    domain.UserRepository
	ServiceRepo domain.ServiceRepository
}

func NewUseCase(r domain.Repositories) domain.ServiceUseCase {
	return &ServiceUsecase{
		UserRepo:    r.User,
		ServiceRepo: r.Service,
	}
}

func (a *ServiceUsecase) Create(service *domain.Service) (int, error) {
	_, err := a.ServiceRepo.ReadURL(service.Url)
	if err == nil {
		return 201, nil
	}
	if err := a.ServiceRepo.Create(service); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *ServiceUsecase) Read(uid string) ([]domain.Service, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.Service{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/service/read")
	if err != nil {
		return []domain.Service{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.Service{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	b, err := a.ServiceRepo.Read()
	if err != nil {
		return []domain.Service{}, 400, err
	}
	return b, 200, nil
}

func (a *ServiceUsecase) Update(service *domain.ServiceForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/service/update")
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
	if err := a.ServiceRepo.Update(service); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *ServiceUsecase) Delete(service *domain.Service, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/service/create")
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
	if err := a.ServiceRepo.Delete(service.ID); err != nil {
		return 400, err
	}
	return 204, nil
}
