package gardenTypeUsecase

import (
	"garden/internal/domain/gardenType"
	"strconv"
)

type usecase struct {
	GardenRepo gardenTypeDomain.GardenTypeRepository
}

func NewUseCase(r gardenTypeDomain.GardenTypeRepository) gardenTypeDomain.GardenTypeUseCase {
	return &usecase{
		GardenRepo: r,
	}
}

func (a *usecase) Create(gardenType *gardenTypeDomain.GardenType) error {
	if err := a.GardenRepo.Create(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(id string) ([]gardenTypeDomain.GardenType, error) {
	idInt, err := strconv.Atoi(id)
	b, err := a.GardenRepo.ReadID(uint(idInt))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *usecase) Update(gardenType *gardenTypeDomain.GardenTypeForm) error {
	if err := a.GardenRepo.Update(gardenType); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(gardenType *gardenTypeDomain.GardenType) error {
	if err := a.GardenRepo.Delete(gardenType.ID); err != nil {
		return err
	}
	return nil
}
