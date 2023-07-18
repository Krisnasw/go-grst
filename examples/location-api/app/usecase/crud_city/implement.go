package crud_city_usecase

import (
	"errors"
	"fmt"

	"github.com/krisnasw/go-grst/examples/location-api/app/repository"
	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

type usecase struct {
	city_repo repository.CityRepository
}

func New(city_repo repository.CityRepository) UseCase {
	return &usecase{
		city_repo: city_repo,
	}
}
func (uc *usecase) GetByPrimaryKey(id int) (*entity.City, error) {
	data, err := uc.city_repo.GetById(id)
	if err != nil {
		if errors.Is(err, repository.ErrCityNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) GetAll() ([]*entity.City, error) {
	data, err := uc.city_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Create(in entity.City) (*entity.City, error) {
	data, err := uc.city_repo.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Update(in entity.City) (*entity.City, error) {
	data, err := uc.city_repo.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Delete(id int) error {
	err := uc.city_repo.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return nil
}
