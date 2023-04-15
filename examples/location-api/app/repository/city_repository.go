package repository

import (
	"errors"

	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

var ErrCityNotFound = errors.New("City not found")
var ErrCityAlreadyExist = errors.New("City already exist")

type CityRepository interface {
	GetById(id int) (*entity.City, error)
	GetAll() ([]*entity.City, error)
	Create(in entity.City) (*entity.City, error)
	Update(in entity.City) (*entity.City, error)
	Delete(id int) error
	Search(keyword string) ([]*entity.City, error)
}
