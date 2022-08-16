package tagUsecase

import (
	"garden/internal/domain"
	"strconv"
)

type Usecase struct {
	TagRepo domain.TagRepository
}

func NewUseCase(r domain.TagRepository) domain.TagUseCase {
	return &Usecase{
		TagRepo: r,
	}
}

func (a *Usecase) Create(tag *domain.Tag) error {
	if err := a.TagRepo.Create(tag); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(pageNumber string) ([]domain.Tag, error) {
	if pageNumber == "" {
		pageNumber = "1"
	}
	nInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return []domain.Tag{}, err
	}
	span := nInt * 10
	b, err := a.TagRepo.Read(span)
	if err != nil {
		return []domain.Tag{}, err
	}
	return b, nil
}

func (a *Usecase) ReadID(id string) ([]domain.Tag, error) {
	idInt, err := strconv.Atoi(id)
	t, err := a.TagRepo.ReadID(uint(idInt))
	if err != nil {
		return []domain.Tag{}, err
	}
	return t, nil
}

func (a *Usecase) Update(tag *domain.TagForm) error {
	if err := a.TagRepo.Update(tag); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(tag *domain.Tag) error {
	if err := a.TagRepo.Delete(tag.ID); err != nil {
		return err
	}
	return nil
}
