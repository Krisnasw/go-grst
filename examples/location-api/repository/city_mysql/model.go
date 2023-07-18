package city_mysql

import (
	"time"

	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

type CityModel struct {
	// table_name = "tbl_city"
	Id         int    `gorm:"primary_key;column:id"`
	ProvinceId int    `gorm:"column:province_id"`
	Name       string `gorm:"column:name"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (model *CityModel) ToCityEntity() *entity.City {
	return &entity.City{
		Id:         model.Id,
		ProvinceId: model.ProvinceId,
		Name:       model.Name,
	}
}
func (CityModel) FromCityEntity(in entity.City) *CityModel {
	return &CityModel{
		Id:         in.Id,
		ProvinceId: in.ProvinceId,
		Name:       in.Name,
	}
}
