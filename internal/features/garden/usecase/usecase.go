package gardenUsecase

import (
	"errors"
	"strconv"
	"strings"

	"garden/internal/domain"
)

type gardenUsecase struct {
	UserRepo    domain.UserRepository
	GardenRepo  domain.GardenRepository
	ServiceRepo domain.ServiceRepository
}

func NewUseCase(a domain.Repositories) domain.GardenUseCase {
	return &gardenUsecase{
		UserRepo:    a.User,
		GardenRepo:  a.Garden,
		ServiceRepo: a.Service,
	}
}

func (a *gardenUsecase) Create(garden *domain.Garden, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/garden/create")
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
	if err := a.GardenRepo.Create(garden); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *gardenUsecase) Read(form domain.ReadGardenForm) ([]domain.Garden, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(form.Uid)
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/garden/read")
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	if form.ID != "" {
		idInt, err := strconv.Atoi(form.ID)
		t, err := a.GardenRepo.ReadID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, 400, err
		}
		return t, 200, nil
	} else if form.UserID != "" {
		idInt, err := strconv.Atoi(form.UserID)
		t, err := a.GardenRepo.ReadUID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, 400, err
		}
		return t, 200, nil
	}
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	span := nInt * 10
	b, err := a.GardenRepo.Read(span)
	if err != nil {
		return []domain.Garden{}, 400, err
	}
	return b, 200, nil
}

func (a *gardenUsecase) Update(garden *domain.GardenForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/garden/update")
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
	if err := a.GardenRepo.Update(garden); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *gardenUsecase) Delete(garden *domain.Garden, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/garden/delete")
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
	if err := a.GardenRepo.Delete(garden.ID); err != nil {
		return 400, err
	}
	if err := a.GardenRepo.DeleteLocation(garden.ID); err != nil {
		return 400, err
	}
	return 204, nil
}

func (a *gardenUsecase) CreateLocation(location *domain.GardenLocation, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/loc/create")
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
	if err := a.GardenRepo.CreateLocation(location); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *gardenUsecase) ReadLocation(gid string, pageNumber string, uid string) ([]domain.GardenLocation, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/loc/read")
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
		t, err := a.GardenRepo.ReadLocation(span)
		if err != nil {
			return []domain.GardenLocation{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(gid)
	b, err := a.GardenRepo.ReadLocationID(uint(idInt))
	if err != nil {
		return []domain.GardenLocation{}, 400, err
	}
	return b, 200, nil
}

func (a *gardenUsecase) UpdateLocation(loc *domain.GardenLocationForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/loc/update")
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
	if err := a.GardenRepo.UpdateLocation(loc); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *gardenUsecase) DeleteLocation(loc *domain.GardenLocation, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/loc/delete")
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
	if err := a.GardenRepo.DeleteLocation(loc.ID); err != nil {
		return 400, err
	}
	return 204, nil
}

func (a *gardenUsecase) CreateType(gardenType *domain.GardenType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/gardenType/create")
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
	if err := a.GardenRepo.CreateType(gardenType); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *gardenUsecase) ReadType(id string, uid string) ([]domain.GardenType, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/gardenType/read")
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	b, err := a.GardenRepo.ReadTypeID(uint(idInt))
	if err != nil {
		return []domain.GardenType{}, 400, err
	}
	return b, 200, nil
}

func (a *gardenUsecase) UpdateType(gardenType *domain.GardenTypeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/gardenType/update")
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
	if err := a.GardenRepo.UpdateType(gardenType); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *gardenUsecase) DeleteType(gardenType *domain.GardenType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/gardenType/delete")
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
	if err := a.GardenRepo.DeleteType(gardenType.ID); err != nil {
		return 400, err
	}
	return 204, nil
}
