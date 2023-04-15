package crud_city_usecase

import (
	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

type Repository interface {
	GetByPrimaryKey(id int) (*entity.City, error)
	GetAll() ([]*entity.City, error)
	Create(in entity.City) (*entity.City, error)
	Update(in entity.City) (*entity.City, error)
	Delete(id int) error
}

type UseCase interface {
	GetByPrimaryKey(id int) (*entity.City, error)
	GetAll() ([]*entity.City, error)
	Create(in entity.City) (*entity.City, error)
	Update(in entity.City) (*entity.City, error)
	Delete(id int) error
}
