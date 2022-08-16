package serviceUsecase

import (
	"garden/internal/domain"
)

type usecase struct {
	ServiceRepo domain.ServiceRepository
}

func NewUseCase(r domain.ServiceRepository) domain.ServiceUseCase {
	return &usecase{
		ServiceRepo: r,
	}
}

func (a *usecase) Create(service *domain.Service) error {
	_, err := a.ServiceRepo.ReadURL(service.Url)
	if err == nil {
		return nil
	}
	if err := a.ServiceRepo.Create(service); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read() ([]domain.Service, error) {
	b, err := a.ServiceRepo.Read()
	if err != nil {
		return []domain.Service{}, err
	}
	return b, nil
}

func (a *usecase) Update(service *domain.ServiceForm) error {
	if err := a.ServiceRepo.Update(service); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(service *domain.Service) error {
	if err := a.ServiceRepo.Delete(service.ID); err != nil {
		return err
	}
	return nil
}
