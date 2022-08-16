package gardenUsecase

import (
	"garden/internal/domain"
	"garden/internal/domain/garden"
	"garden/internal/domain/gardenLocation"
	"strconv"
)

type usecase struct {
	GardenRepo    gardenDomain.GardenRepository
	GardenLocRepo gardenLocationDomain.GardenLocRepository
}

func NewUseCase(r domain.Repositories) gardenDomain.GardenUseCase {
	return &usecase{
		GardenRepo:    r.Garden,
		GardenLocRepo: r.GardenLoc,
	}
}

func (a *usecase) Create(garden *gardenDomain.Garden) error {
	if err := a.GardenRepo.Create(garden); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(form gardenDomain.ReadGardenForm) ([]gardenDomain.Garden, error) {
	if form.ID != "" {
		idInt, err := strconv.Atoi(form.ID)
		t, err := a.GardenRepo.ReadID(uint(idInt))
		if err != nil {
			return nil, err
		}
		return t, nil
	} else if form.UserID != "" {
		idInt, err := strconv.Atoi(form.UserID)
		t, err := a.GardenRepo.ReadUID(uint(idInt))
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return nil, err
	}
	span := nInt * 10
	b, err := a.GardenRepo.Read(span)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *usecase) Update(garden *gardenDomain.GardenForm) error {
	if err := a.GardenRepo.Update(garden); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(garden *gardenDomain.Garden) error {
	if err := a.GardenRepo.Delete(garden.ID); err != nil {
		return err
	}
	if err := a.GardenLocRepo.Delete(garden.ID); err != nil {
		return err
	}
	return nil
}
