package city_mysql

import (
	"errors"

	repository_intf "github.com/krisnasw/go-grst/examples/location-api/app/repository"
	"github.com/krisnasw/go-grst/examples/location-api/entity"
	"gorm.io/gorm"
)

type repository struct {
	db        *gorm.DB
	tableName string
}

func New(db *gorm.DB) repository_intf.CityRepository {
	return &repository{db, "tbl_city"}
}

func (r *repository) GetById(id int) (*entity.City, error) {
	data, err := CityModel{}.QueryGet(r.db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrCityNotFound
		}
		return nil, err
	}
	return data.ToCityEntity(), nil
}

func (r *repository) GetAll() ([]*entity.City, error) {
	datas, err := CityModel{}.QueryGetAll(r.db)
	if err != nil {
		return nil, err
	}

	result := []*entity.City{}
	for _, data := range datas {
		result = append(result, data.ToCityEntity())
	}
	return result, err
}

func (r *repository) Create(in entity.City) (*entity.City, error) {
	data, err := CityModel{}.QueryCreate(r.db, *CityModel{}.FromCityEntity(in))
	if err != nil {
		return nil, err
	}
	return data.ToCityEntity(), nil
}

func (r *repository) Update(in entity.City) (*entity.City, error) {
	_, err := CityModel{}.QueryUpdate(r.db, *CityModel{}.FromCityEntity(in))
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (r *repository) Delete(id int) error {
	err := CityModel{}.QueryDelete(r.db, id)
	if err != nil {
		return err
	}
	return err
}

func (r *repository) Search(keyword string) ([]*entity.City, error) {
	datas := []*CityModel{}
	err := r.db.Table(r.tableName).Where("name like ?", "%"+keyword+"%").Find(&datas).Error

	result := []*entity.City{}
	for _, data := range datas {
		result = append(result, data.ToCityEntity())
	}
	return result, err

}
