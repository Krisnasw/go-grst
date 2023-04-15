package province_mysql

import (
	"time"

	"github.com/krisnasw/go-grst/examples/location-api/entity"
)

type ProvinceModel struct {
	// table_name = "tbl_province"
	Id   int    `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (model *ProvinceModel) ToProvinceEntity() *entity.Province {
	return &entity.Province{
		Id:   model.Id,
		Name: model.Name,
	}
}
func (ProvinceModel) FromProvinceEntity(in entity.Province) *ProvinceModel {
	return &ProvinceModel{
		Id:   in.Id,
		Name: in.Name,
	}
}
