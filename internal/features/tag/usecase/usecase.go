package tagUsecase

import (
	"errors"
	"strconv"
	"strings"

	"garden/internal/domain"
)

type TagUsecase struct {
	UserRepo    domain.UserRepository
	ServiceRepo domain.ServiceRepository
	TagRepo     domain.TagRepository
	TreeRepo    domain.TreeRepository
}

func NewUseCase(r domain.Repositories) domain.TagUseCase {
	return &TagUsecase{
		UserRepo:    r.User,
		ServiceRepo: r.Service,
		TagRepo:     r.Tag,
		TreeRepo:    r.Tree,
	}
}

func (a *TagUsecase) Create(tag *domain.Tag, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tag/create")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t.AccessList, ",")
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
	if err := a.TagRepo.Create(tag); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TagUsecase) Read(pageNumber string, uid string) ([]domain.Tag, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tag/read")
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	List := strings.Split(t.AccessList, ",")
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
	b, err := a.TagRepo.Read(span)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	return b, 200, nil
}

func (a *TagUsecase) ReadID(id string) ([]domain.Tag, int, error) {
	idInt, err := strconv.Atoi(id)
	t, err := a.TagRepo.ReadID(uint(idInt))
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	return t, 200, nil
}

func (a *TagUsecase) Update(tag *domain.TagForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tag/update")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t.AccessList, ",")
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
	if err := a.TagRepo.Update(tag); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TagUsecase) Delete(tag *domain.Tag, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.ReadID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadURL("user/tag/delete")
	if err != nil {
		return 400, err
	}
	t, err := a.UserRepo.ReadTypeID(u.Type)
	if err != nil {
		return 400, err
	}
	List := strings.Split(t.AccessList, ",")
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
	if err := a.TagRepo.Delete(tag.ID); err != nil {
		return 400, err
	}
	return 204, nil
}

func (a *TagUsecase) CreateTree(tree *domain.Tree, uid string) (int, error) {
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
	List := strings.Split(t.AccessList, ",")
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
