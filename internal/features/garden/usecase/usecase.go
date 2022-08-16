package gardenUsecase

import (
	"garden/internal/domain"
	"strconv"
)

type usecase struct {
	GardenRepo    domain.GardenRepository
	GardenLocRepo domain.GardenLocRepository
}

func NewUseCase(r domain.Repositories) domain.GardenUseCase {
	return &usecase{
		GardenRepo:    r.Garden,
		GardenLocRepo: r.GardenLoc,
	}
}

func (a *usecase) Create(garden *domain.Garden) error {
	if err := a.GardenRepo.Create(garden); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(form domain.ReadGardenForm) ([]domain.Garden, error) {
	if form.ID != "" {
		idInt, err := strconv.Atoi(form.ID)
		t, err := a.GardenRepo.ReadID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, err
		}
		return t, nil
	} else if form.UserID != "" {
		idInt, err := strconv.Atoi(form.UserID)
		t, err := a.GardenRepo.ReadUID(uint(idInt))
		if err != nil {
			return []domain.Garden{}, err
		}
		return t, nil
	}
	if form.PageNumber == "" {
		form.PageNumber = "1"
	}
	nInt, err := strconv.Atoi(form.PageNumber)
	if err != nil {
		return []domain.Garden{}, err
	}
	span := nInt * 10
	b, err := a.GardenRepo.Read(span)
	if err != nil {
		return []domain.Garden{}, err
	}
	return b, nil
}

func (a *usecase) Update(garden *domain.GardenForm) error {
	if err := a.GardenRepo.Update(garden); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(garden *domain.Garden) error {
	if err := a.GardenRepo.Delete(garden.ID); err != nil {
		return err
	}
	if err := a.GardenLocRepo.Delete(garden.ID); err != nil {
		return err
	}
	return nil
}
