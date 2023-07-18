package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	crud_province_usecase "github.com/krisnasw/go-grst/examples/location-api/app/usecase/crud_province"
	pbProvince "github.com/krisnasw/go-grst/examples/location-api/handler/grst/province"
	grst_errors "github.com/krisnasw/go-grst/grst/errors"
	"google.golang.org/grpc/codes"
)

type ProvinceHandler struct {
	pbProvince.UnimplementedProvinceServer
	crud_province_uc crud_province_usecase.UseCase
}

func NewProvinceHandler(crud_province_uc crud_province_usecase.UseCase) pbProvince.ProvinceServer {
	return &ProvinceHandler{crud_province_uc: crud_province_uc}
}

func (h *ProvinceHandler) Get(ctx context.Context, req *pbProvince.GetReq) (*pbProvince.Province, error) {
	if err := pbProvince.ValidateRequest(req); err != nil {
		return nil, err
	}

	data, err := h.crud_province_uc.GetByPrimaryKey(int(req.Id))
	if err != nil {
		if errors.Is(err, crud_province_usecase.ErrRecordNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 1101, err.Error(), &grst_errors.ErrorDetail{})
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 1102, err.Error(), &grst_errors.ErrorDetail{})
	}

	return &pbProvince.Province{
		Id:   int32(data.Id),
		Name: data.Name,
	}, nil
}

func (h *ProvinceHandler) GetAll(ctx context.Context, req *empty.Empty) (*pbProvince.Provinces, error) {
	datas, err := h.crud_province_uc.GetAll()
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 1201, err.Error(), &grst_errors.ErrorDetail{})
	}
	result := &pbProvince.Provinces{
		Provinces: []*pbProvince.Province{},
	}
	for _, data := range datas {
		result.Provinces = append(result.Provinces, &pbProvince.Province{
			Id:   int32(data.Id),
			Name: data.Name,
		})
	}
	return result, nil
}
