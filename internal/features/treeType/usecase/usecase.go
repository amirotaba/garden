package treeTypeUsecase

import (
	"garden/internal/domain"
	"strconv"
)

type Usecase struct {
	TreeRepo domain.TreeTypeRepository
}

func NewUseCase(r domain.TreeTypeRepository) domain.TreeTypeUseCase {
	return &Usecase{
		TreeRepo: r,
	}
}

func (a *Usecase) Create(treeType *domain.TreeType) error {
	if err := a.TreeRepo.Create(treeType); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(id string) ([]domain.TreeType, error) {
	if id == "" {
		t, err := a.TreeRepo.Read()
		if err != nil {
			return []domain.TreeType{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(id)
	b, err := a.TreeRepo.ReadID(uint(idInt))
	if err != nil {
		return []domain.TreeType{}, err
	}
	return b, nil
}

func (a *Usecase) Update(treeType *domain.TreeTypeForm) error {
	if err := a.TreeRepo.Update(treeType); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(tree *domain.TreeType) error {
	if err := a.TreeRepo.Delete(tree.ID); err != nil {
		return err
	}
	return nil
}
