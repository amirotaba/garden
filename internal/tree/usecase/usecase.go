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

func NewTreeUseCase(r domain.Repositories) domain.TreeUseCase {
	return &TreeUsecase{
		TreeRepo:    r.Tree,
		ServiceRepo: r.Service,
		UserRepo:    r.User,
	}
}

func (a *TreeUsecase) CreateTree(tree *domain.Tree, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tree/create")
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
	if err := a.TreeRepo.CreateTree(tree); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) ReadTree(form domain.ReadTreeForm) ([]domain.Tree, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(form.Uid)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tree/read")
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
		t, err := a.TreeRepo.ReadTreeID(uint(idInt), "garden_id = ?")
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	} else if form.Tp != "" {
		idInt, err := strconv.Atoi(form.Tp)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.TreeRepo.ReadTreeByType(uint(idInt), span)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	}
	b, err := a.TreeRepo.ReadTree(span)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	return b, 200, nil
}

func (a *TreeUsecase) ReadTreeUser(form domain.ReadTreeUserForm) ([]domain.Tree, int, error) {
	if form.GardenID != "" {
		idInt, err := strconv.Atoi(form.GardenID)
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		t, err := a.TreeRepo.ReadTreeID(uint(idInt), "garden_id = ?")
		if err != nil {
			return []domain.Tree{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(form.ID)
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	t, err := a.TreeRepo.ReadTreeID(uint(idInt), "id = ?")
	if err != nil {
		return []domain.Tree{}, 400, err
	}
	return t, 200, nil
}

func (a *TreeUsecase) UpdateTree(tree *domain.TreeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tree/update")
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
	if err := a.TreeRepo.UpdateTree(tree); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) DeleteTree(tree *domain.Tree, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tree/delete")
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
	if err := a.TreeRepo.DeleteTree(tree.ID); err != nil {
		return 400, err
	}
	return 204, nil
}

func (a *TreeUsecase) CreateTreeType(treeType *domain.TreeType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/treeType/create")
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
	if err := a.TreeRepo.CreateTreeType(treeType); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) ReadTreeType(id string, uid string) ([]domain.TreeType, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/treeType/read")
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
		t, err := a.TreeRepo.ReadTreeType()
		if err != nil {
			return []domain.TreeType{}, 400, err
		}
		return t, 200, nil
	}
	idInt, err := strconv.Atoi(id)
	b, err := a.TreeRepo.ReadTreeTypeID(uint(idInt))
	if err != nil {
		return []domain.TreeType{}, 400, err
	}
	return b, 200, nil
}

func (a *TreeUsecase) UpdateTreeType(treeType *domain.TreeTypeForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/treeType/update")
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
	if err := a.TreeRepo.UpdateTreeType(treeType); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TreeUsecase) DeleteTreeType(tree *domain.TreeType, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/treeType/delete")
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
	if err := a.TreeRepo.DeleteTreeType(tree.ID); err != nil {
		return 400, err
	}
	return 204, nil
}
