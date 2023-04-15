package user_repository_impl

import (
	"context"
	"errors"
	"fmt"

	repository_intf "github.com/krisnasw/go-grst/examples/users-api/app/repository"
	pbProvince "github.com/krisnasw/go-grst/examples/users-api/clients/grst/province"
	"github.com/krisnasw/go-grst/examples/users-api/entity"
	"gorm.io/gorm"
)

type repository struct {
	db          *gorm.DB
	provinceCli pbProvince.ProvinceClient
	tableName   string
}

func New(db *gorm.DB, provinceCli pbProvince.ProvinceClient) repository_intf.UserRepository {
	return &repository{db, provinceCli, "tbl_users"}
}

func (r *repository) GetById(id int) (*entity.User, error) {
	data := &UserModel{}
	err := r.db.Table(r.tableName).Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrUserNotFound
		}
		return nil, err
	}
	return data.ToUserEntity(), nil
}

func (r *repository) GetProfileById(id int) (*entity.UserProfile, error) {
	data, err := r.GetById(id)
	if err != nil {
		return nil, err
	}
	provinceData, err := r.provinceCli.Get(context.Background(), &pbProvince.GetReq{Id: int32(data.ProvinceId)})
	if err != nil {
		return nil, fmt.Errorf("province-api error: %v", err)
	}
	result := entity.UserProfile{}.FromUser(*data, provinceData.Name)
	return &result, err
}
