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

func NewTagUseCase(r domain.Repositories) domain.TagUseCase {
	return &TagUsecase{
		UserRepo:    r.User,
		ServiceRepo: r.Service,
		TagRepo:     r.Tag,
		TreeRepo:    r.Tree,
	}
}

func (a *TagUsecase) CreateTag(tag *domain.Tag, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tag/create")
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
	if err := a.TagRepo.CreateTag(tag); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TagUsecase) ReadTag(pageNumber string, uid string) ([]domain.Tag, int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tag/read")
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	t, err := a.UserRepo.ReadUserTypeID(u.Type)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	List := strings.Split(t[0].AccessList, ",")
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
	b, err := a.TagRepo.ReadTag(span)
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	return b, 200, nil
}

func (a *TagUsecase) ReadTagID(id string) ([]domain.Tag, int, error) {
	idInt, err := strconv.Atoi(id)
	t, err := a.TagRepo.ReadTagID(uint(idInt))
	if err != nil {
		return []domain.Tag{}, 400, err
	}
	return t, 200, nil
}

func (a *TagUsecase) UpdateTag(tag *domain.TagForm, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tag/update")
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
	if err := a.TagRepo.UpdateTag(tag); err != nil {
		return 400, err
	}
	return 201, nil
}

func (a *TagUsecase) DeleteTag(tag *domain.Tag, uid string) (int, error) {
	var boolean bool
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return 400, err
	}
	u, err := a.UserRepo.AccountID(uint(uidInt))
	if err != nil {
		return 400, err
	}
	SID, err := a.ServiceRepo.ReadServiceUrl("user/tag/delete")
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
	if err := a.TagRepo.DeleteTag(tag.ID); err != nil {
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
