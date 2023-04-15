package crud_province_usecase

import (
	"github.com/herryg91/cdd/examples/location-api/entity"
)

type Repository interface {
	GetByPrimaryKey(id int) (*entity.Province, error)
	GetAll() ([]*entity.Province, error)
	Create(in entity.Province) (*entity.Province, error)
	Update(in entity.Province) (*entity.Province, error)
	Delete(id int) error
}

type UseCase interface {
	GetByPrimaryKey(id int) (*entity.Province, error)
	GetAll() ([]*entity.Province, error)
	Create(in entity.Province) (*entity.Province, error)
	Update(in entity.Province) (*entity.Province, error)
	Delete(id int) error
}
