package repository

import (
	"errors"

	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

var ErrProvinceNotFound = errors.New("Province not found")
var ErrProvinceAlreadyExist = errors.New("Province already exist")

type ProvinceRepository interface {
	Get(id int) (*entity.Province, error)
	GetByIds(ids []int) (map[int]entity.Province, error)
	GetAll() ([]*entity.Province, error)
	Create(in entity.Province) (*entity.Province, error)
	Update(in entity.Province) (*entity.Province, error)
	Delete(id int) error
}
