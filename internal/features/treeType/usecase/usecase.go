package treeTypeUsecase

import (
	"garden/internal/domain/treeType"
	"strconv"
)

type Usecase struct {
	TreeRepo treeTypeDomain.TreeTypeRepository
}

func NewUseCase(r treeTypeDomain.TreeTypeRepository) treeTypeDomain.TreeTypeUseCase {
	return &Usecase{
		TreeRepo: r,
	}
}

func (a *Usecase) Create(treeType *treeTypeDomain.TreeType) error {
	if err := a.TreeRepo.Create(treeType); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(id string) ([]treeTypeDomain.TreeType, error) {
	if id == "" {
		t, err := a.TreeRepo.Read()
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	b, err := a.TreeRepo.ReadID(uint(idInt))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *Usecase) Update(treeType *treeTypeDomain.TreeTypeForm) error {
	if err := a.TreeRepo.Update(treeType); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(tree *treeTypeDomain.TreeType) error {
	if err := a.TreeRepo.Delete(tree.ID); err != nil {
		return err
	}
	return nil
}
