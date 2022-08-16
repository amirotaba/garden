package tagUsecase

import (
	"garden/internal/domain/tag"
	"strconv"
)

type Usecase struct {
	TagRepo tagDomain.TagRepository
}

func NewUseCase(r tagDomain.TagRepository) tagDomain.TagUseCase {
	return &Usecase{
		TagRepo: r,
	}
}

func (a *Usecase) Create(tag *tagDomain.Tag) error {
	if err := a.TagRepo.Create(tag); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Read(pageNumber string) ([]tagDomain.Tag, error) {
	if pageNumber == "" {
		pageNumber = "1"
	}
	nInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return nil, err
	}
	span := nInt * 10
	b, err := a.TagRepo.Read(span)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *Usecase) ReadID(id string) ([]tagDomain.Tag, error) {
	idInt, err := strconv.Atoi(id)
	t, err := a.TagRepo.ReadID(uint(idInt))
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (a *Usecase) Update(tag *tagDomain.TagForm) error {
	if err := a.TagRepo.Update(tag); err != nil {
		return err
	}
	return nil
}

func (a *Usecase) Delete(tag *tagDomain.Tag) error {
	if err := a.TagRepo.Delete(tag.ID); err != nil {
		return err
	}
	return nil
}
