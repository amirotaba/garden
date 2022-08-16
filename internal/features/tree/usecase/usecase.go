package treeUsecase

import (
	"garden/internal/domain"
	"strconv"
)

type Usecase struct {
	TreeRepo domain.TreeRepository
}

func NewUseCase(r domain.TreeRepository) domain.TreeUseCase {
	return &Usecase{
		TreeRepo: r,
	}
}

func (a *Usecase) Create(tree *domain.Tree) error {
	//tree.Qr = make a QRCode
	if err := a.TreeRepo.Create(tree); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(form domain.ReadTreeForm) ([]domain.Tree, error) {
	var readFormType domain.ReadTreeType
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return []domain.Tree{}, err
	}
	readFormType.Span = nInt * 10
	if form.GardenID != "" {
		idInt, err := strconv.Atoi(form.GardenID)
		if err != nil {
			return []domain.Tree{}, err
		}
		readForm := domain.ReadTreeID{
			Query: "garden_id = ?",
			ID:    uint(idInt),
		}
		t, err := a.TreeRepo.ReadID(readForm)
		if err != nil {
			return []domain.Tree{}, err
		}
		return t, nil
	} else if form.Tp != "" {
		idInt, err := strconv.Atoi(form.Tp)
		if err != nil {
			return []domain.Tree{}, err
		}
		readFormType.ID = uint(idInt)
		t, err := a.TreeRepo.ReadByType(readFormType)
		if err != nil {
			return []domain.Tree{}, err
		}
		return t, nil
	}
	b, err := a.TreeRepo.Read(readFormType.Span)
	if err != nil {
		return []domain.Tree{}, err
	}
	return b, nil
}

func (a *Usecase) ReadUser(form domain.ReadTreeUserForm) ([]domain.Tree, error) {
	var readForm domain.ReadTreeID
	if form.GardenID != "" {
		idInt, err := strconv.Atoi(form.GardenID)
		if err != nil {
			return []domain.Tree{}, err
		}
		readForm = domain.ReadTreeID{
			Query: "garden_id = ?",
			ID:    uint(idInt),
		}
		t, err := a.TreeRepo.ReadID(readForm)
		if err != nil {
			return []domain.Tree{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(form.ID)
	if err != nil {
		return []domain.Tree{}, err
	}
	readForm = domain.ReadTreeID{
		Query: "id = ?",
		ID:    uint(idInt),
	}
	t, err := a.TreeRepo.ReadID(readForm)
	if err != nil {
		return []domain.Tree{}, err
	}
	return t, nil
}

func (a *Usecase) Update(tree *domain.TreeForm) error {
	if err := a.TreeRepo.Update(tree); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(tree *domain.Tree) error {
	if err := a.TreeRepo.Delete(tree.ID); err != nil {
		return err
	}
	return nil
}
