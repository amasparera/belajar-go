package teslayer

import "errors"

type CategoryService struct {
	CategoryRepository CategoryRepository
}

func (service CategoryService) Get(Id string) (*Category, error) {
	category, _ := service.CategoryRepository.FindById(Id)
	if category == nil {
		return category, errors.New("Data null")
	} else {
		return category, nil
	}

}
