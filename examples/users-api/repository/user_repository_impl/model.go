package user_repository_impl

import (
	"time"

	"github.com/krisnasw/go-grst/examples/users-api/entity"
)

type UserModel struct {
	// table_name = "tbl_user"
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ProvinceId int    `json:"province_id"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (model *UserModel) ToUserEntity() *entity.User {
	return &entity.User{
		Id:         model.Id,
		ProvinceId: model.ProvinceId,
		Name:       model.Name,
	}
}
func (UserModel) FromUserEntity(in entity.User) *UserModel {
	return &UserModel{
		Id:         in.Id,
		ProvinceId: in.ProvinceId,
		Name:       in.Name,
	}
}
