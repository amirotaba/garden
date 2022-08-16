package treeUsecase

import (
	"garden/internal/domain/tree"
	"garden/internal/domain/treeType"
	"strconv"
)

type Usecase struct {
	TreeRepo treeDomain.TreeRepository
}

func NewUseCase(r treeDomain.TreeRepository) treeDomain.TreeUseCase {
	return &Usecase{
		TreeRepo: r,
	}
}

func (a *Usecase) Create(tree *treeDomain.Tree) error {
	//tree.Qr = make a QRCode
	if err := a.TreeRepo.Create(tree); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(form treeDomain.ReadTreeForm) ([]treeDomain.Tree, error) {
	var readFormType treeTypeDomain.ReadTreeType
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return nil, err
	}
	readFormType.Span = nInt * 10
	if form.GardenID != "" {
		idInt, err := strconv.Atoi(form.GardenID)
		if err != nil {
			return nil, err
		}
		readForm := treeDomain.ReadTreeID{
			Query: "garden_id = ?",
			ID:    uint(idInt),
		}
		t, err := a.TreeRepo.ReadID(readForm)
		if err != nil {
			return nil, err
		}
		return t, nil
	} else if form.Tp != "" {
		idInt, err := strconv.Atoi(form.Tp)
		if err != nil {
			return nil, err
		}
		readFormType.ID = uint(idInt)
		t, err := a.TreeRepo.ReadByType(readFormType)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	b, err := a.TreeRepo.Read(readFormType.Span)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *Usecase) ReadUser(form treeDomain.ReadTreeUserForm) ([]treeDomain.Tree, error) {
	var readForm treeDomain.ReadTreeID
	if form.GardenID != "" {
		idInt, err := strconv.Atoi(form.GardenID)
		if err != nil {
			return nil, err
		}
		readForm = treeDomain.ReadTreeID{
			Query: "garden_id = ?",
			ID:    uint(idInt),
		}
		t, err := a.TreeRepo.ReadID(readForm)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(form.ID)
	if err != nil {
		return nil, err
	}
	readForm = treeDomain.ReadTreeID{
		Query: "id = ?",
		ID:    uint(idInt),
	}
	t, err := a.TreeRepo.ReadID(readForm)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (a *Usecase) Update(tree *treeDomain.TreeForm) error {
	if err := a.TreeRepo.Update(tree); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(tree *treeDomain.Tree) error {
	if err := a.TreeRepo.Delete(tree.ID); err != nil {
		return err
	}
	return nil
}
