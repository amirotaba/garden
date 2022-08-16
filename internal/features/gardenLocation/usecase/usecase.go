package gardenLocUsecase

import (
	"garden/internal/domain"
	"strconv"
)

type usecase struct {
	Repository domain.GardenLocRepository
}

func NewUseCase(r domain.GardenLocRepository) domain.GardenLocUseCase {
	return &usecase{
		Repository: r,
	}
}

func (a *usecase) Create(location *domain.GardenLocation) error {
	if err := a.Repository.Create(location); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(gid string, pageNumber string) ([]domain.GardenLocation, error) {
	if gid == "" {
		if pageNumber == "" {
			pageNumber = "1"
		}
		nInt, err := strconv.Atoi(pageNumber)
		if err != nil {
			return []domain.GardenLocation{}, err
		}
		span := nInt * 10
		t, err := a.Repository.Read(span)
		if err != nil {
			return []domain.GardenLocation{}, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(gid)
	b, err := a.Repository.ReadID(uint(idInt))
	if err != nil {
		return []domain.GardenLocation{}, err
	}
	return b, nil
}

func (a *usecase) Update(loc *domain.GardenLocationForm) error {
	if err := a.Repository.Update(loc); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(loc *domain.GardenLocation) error {
	if err := a.Repository.Delete(loc.ID); err != nil {
		return err
	}
	return nil
}
