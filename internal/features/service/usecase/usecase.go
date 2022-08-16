package serviceUsecase

import (
	"garden/internal/domain/service"
)

type usecase struct {
	ServiceRepo serviceDomain.ServiceRepository
}

func NewUseCase(r serviceDomain.ServiceRepository) serviceDomain.ServiceUseCase {
	return &usecase{
		ServiceRepo: r,
	}
}

func (a *usecase) Create(service *serviceDomain.Service) error {
	_, err := a.ServiceRepo.ReadURL(service.Url)
	if err == nil {
		return nil
	}
	if err := a.ServiceRepo.Create(service); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Read() ([]serviceDomain.Service, error) {
	b, err := a.ServiceRepo.Read()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *usecase) Update(service *serviceDomain.ServiceForm) error {
	if err := a.ServiceRepo.Update(service); err != nil {
		return err
	}
	return nil
}

func (a *usecase) Delete(service *serviceDomain.Service) error {
	if err := a.ServiceRepo.Delete(service.ID); err != nil {
		return err
	}
	return nil
}
