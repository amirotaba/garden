package gardenTypeUsecase

import (
	"garden/internal/domain"
	"strconv"
)

type usecase struct {
	GardenRepo domain.GardenTypeRepository
}

func NewUseCase(r domain.GardenTypeRepository) domain.GardenTypeUseCase {
	return &usecase{
		GardenRepo: r,
	}
}

func (a *usecase) Create(gardenType *domain.GardenType) error {
	if err := a.GardenRepo.Create(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(id string) ([]domain.GardenType, error) {
	idInt, err := strconv.Atoi(id)
	b, err := a.GardenRepo.ReadID(uint(idInt))
	if err != nil {
		return []domain.GardenType{}, err
	}
	return b, nil
}

func (a *usecase) Update(gardenType *domain.GardenTypeForm) error {
	if err := a.GardenRepo.Update(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(gardenType *domain.GardenType) error {
	if err := a.GardenRepo.Delete(gardenType.ID); err != nil {
		return err
	}
	return nil
}
