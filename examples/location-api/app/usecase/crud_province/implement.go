package crud_province_usecase

import (
	"errors"
	"fmt"

	"github.com/krisnasw/go-grst/examples/location-api/app/repository"
	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

type usecase struct {
	province_repo repository.ProvinceRepository
}

func New(province_repo repository.ProvinceRepository) UseCase {
	return &usecase{
		province_repo: province_repo,
	}
}
func (uc *usecase) GetByPrimaryKey(id int) (*entity.Province, error) {
	data, err := uc.province_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrProvinceNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) GetAll() ([]*entity.Province, error) {
	data, err := uc.province_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Create(in entity.Province) (*entity.Province, error) {
	data, err := uc.province_repo.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Update(in entity.Province) (*entity.Province, error) {
	data, err := uc.province_repo.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return data, nil
}
func (uc *usecase) Delete(id int) error {
	err := uc.province_repo.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}
	return nil
}
