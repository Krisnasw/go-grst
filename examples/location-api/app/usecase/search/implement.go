package search_usecase

import (
	"fmt"

	"github.com/krisnasw/go-grst/examples/location-api/app/repository"
	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

type usecase struct {
	city_repo     repository.CityRepository
	province_repo repository.ProvinceRepository
}

func New(city_repo repository.CityRepository, province_repo repository.ProvinceRepository) UseCase {
	return &usecase{
		city_repo:     city_repo,
		province_repo: province_repo,
	}
}

func (uc *usecase) Search(keyword string) ([]entity.CityProfile, error) {
	datas, err := uc.city_repo.Search(keyword)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}

	mapOfProvinces, err := uc.province_repo.GetByIds(entity.Cities(datas).GetProvinceIds())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}

	resp := []entity.CityProfile{}
	for _, data := range datas {
		provinceName := ""
		if val, ok := mapOfProvinces[data.ProvinceId]; ok {
			provinceName = val.Name
		}
		resp = append(resp, entity.CityProfile{}.FromCity(*data, provinceName))
	}
	return resp, nil
}
