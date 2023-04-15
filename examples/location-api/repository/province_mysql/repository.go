package province_mysql

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

func New(db *gorm.DB) repository_intf.ProvinceRepository {
	return &repository{db, "tbl_province"}
}

func (r *repository) Get(id int) (*entity.Province, error) {
	data, err := ProvinceModel{}.QueryGet(r.db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrProvinceNotFound
		}
		return nil, err
	}
	return data.ToProvinceEntity(), nil
}

func (r *repository) GetByIds(ids []int) (map[int]entity.Province, error) {
	datas := []*ProvinceModel{}
	err := r.db.Table(r.tableName).Where("id in (?)", ids).Find(&datas).Error
	if err != nil {
		return map[int]entity.Province{}, err
	}
	result := map[int]entity.Province{}
	for _, data := range datas {
		result[data.Id] = *data.ToProvinceEntity()
	}
	return result, err
}

func (r *repository) GetAll() ([]*entity.Province, error) {
	datas, err := ProvinceModel{}.QueryGetAll(r.db)
	if err != nil {
		return nil, err
	}

	result := []*entity.Province{}
	for _, data := range datas {
		result = append(result, data.ToProvinceEntity())
	}
	return result, err
}

func (r *repository) Create(in entity.Province) (*entity.Province, error) {
	data, err := ProvinceModel{}.QueryCreate(r.db, *ProvinceModel{}.FromProvinceEntity(in))
	if err != nil {
		return nil, err
	}
	return data.ToProvinceEntity(), nil
}

func (r *repository) Update(in entity.Province) (*entity.Province, error) {
	_, err := ProvinceModel{}.QueryUpdate(r.db, *ProvinceModel{}.FromProvinceEntity(in))
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (r *repository) Delete(id int) error {
	err := ProvinceModel{}.QueryDelete(r.db, id)
	if err != nil {
		return err
	}
	return err
}
