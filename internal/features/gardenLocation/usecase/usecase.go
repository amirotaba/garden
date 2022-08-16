package gardenLocUsecase

import (
	"garden/internal/domain/gardenLocation"
	"strconv"
)

type usecase struct {
	Repository gardenLocationDomain.GardenLocRepository
}

func NewUseCase(r gardenLocationDomain.GardenLocRepository) gardenLocationDomain.GardenLocUseCase {
	return &usecase{
		Repository: r,
	}
}

func (a *usecase) Create(location *gardenLocationDomain.GardenLocation) error {
	if err := a.Repository.Create(location); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read(form gardenLocationDomain.GardenLocRead) ([]gardenLocationDomain.GardenLocation, error) {
	if form.GardenID == "" {
		if form.PageNumber == "" {
			form.PageNumber = "1"
		}
		nInt, err := strconv.Atoi(form.PageNumber)
		if err != nil {
			return nil, err
		}
		span := nInt * 10
		t, err := a.Repository.Read(span)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	idInt, err := strconv.Atoi(form.GardenID)
	b, err := a.Repository.ReadID(uint(idInt))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *usecase) Update(loc *gardenLocationDomain.GardenLocationForm) error {
	if err := a.Repository.Update(loc); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(loc *gardenLocationDomain.GardenLocation) error {
	if err := a.Repository.Delete(loc.ID); err != nil {
		return err
	}
	return nil
}
