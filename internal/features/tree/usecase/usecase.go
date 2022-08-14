package treeUsecase

import (
	"errors"
	"strconv"
	"strings"

	"garden/internal/domain"
)

type TreeUsecase struct {
	UserRepo    domain.UserRepository
	ServiceRepo domain.ServiceRepository
	TreeRepo    domain.TreeRepository
}

func NewUseCase(r domain.Repositories) domain.TreeUseCase {
	return &TreeUsecase{
		TreeRepo:    r.Tree,
		ServiceRepo: r.Service,
		UserRepo:    r.User,
	}
}

func (a *TreeUsecase) Create(tree *domain.Tree, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tree/create")
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
	//tree.Qr = make a QRCode
	if err := a.TreeRepo.Create(tree); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) Read(form domain.ReadTreeForm) ([]domain.Tree, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(form.Uid)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tree/read")
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	span := nInt * 10
	if form.GardenID != "" {
		idInt, err := strconv.Atoi(form.GardenID)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.TreeRepo.ReadID(uint(idInt), "garden_id = ?")
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	} else if form.Tp != "" {
		idInt, err := strconv.Atoi(form.Tp)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.TreeRepo.ReadByType(uint(idInt), span)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	}
	b, err := a.TreeRepo.Read(span)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	return b, 200, nil
}

func (a *TreeUsecase) ReadUser(form domain.ReadTreeUserForm) ([]domain.Tree, int, error) {
	if form.GardenID != "" {
		idInt, err := strconv.Atoi(form.GardenID)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.TreeRepo.ReadID(uint(idInt), "garden_id = ?")
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(form.ID)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	t, err := a.TreeRepo.ReadID(uint(idInt), "id = ?")
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	return t, 200, nil
}

func (a *TreeUsecase) Update(tree *domain.TreeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tree/update")
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
	if err := a.TreeRepo.Update(tree); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) Delete(tree *domain.Tree, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tree/delete")
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
	if err := a.TreeRepo.Delete(tree.ID); err != nil {
		return 400, err
	}
	return 204, nil
}

func (a *TreeUsecase) CreateType(treeType *domain.TreeType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/treeType/create")
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
	if err := a.TreeRepo.CreateType(treeType); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) ReadType(id string, uid string) ([]domain.TreeType, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/treeType/read")
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
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
		t, err := a.TreeRepo.ReadType()
		if err != nil {
			return []domain.TreeType{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(id)
	b, err := a.TreeRepo.ReadTypeID(uint(idInt))
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	return b, 200, nil
}

func (a *TreeUsecase) UpdateType(treeType *domain.TreeTypeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/treeType/update")
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
	if err := a.TreeRepo.UpdateType(treeType); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) DeleteType(tree *domain.TreeType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/treeType/delete")
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
	if err := a.TreeRepo.DeleteType(tree.ID); err != nil {
		return 400, err
	}
	return 204, nil
}
