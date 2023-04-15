package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	crud_city_usecase "github.com/krisnasw/go-grst/examples/location-api/app/usecase/crud_city"
	search_usecase "github.com/krisnasw/go-grst/examples/location-api/app/usecase/search"
	pbCity "github.com/krisnasw/go-grst/examples/location-api/handler/grst/city"
	grst_errors "github.com/krisnasw/go-grst/grst/errors"
	"google.golang.org/grpc/codes"
)

type CityHandler struct {
	pbCity.UnimplementedCityServer
	crud_city_uc crud_city_usecase.UseCase
	search_uc    search_usecase.UseCase
}

func NewCityHandler(crud_city_uc crud_city_usecase.UseCase, search_uc search_usecase.UseCase) pbCity.CityServer {
	return &CityHandler{crud_city_uc: crud_city_uc, search_uc: search_uc}
}

func (h *CityHandler) Get(ctx context.Context, req *pbCity.GetReq) (*pbCity.City, error) {
	if err := pbCity.ValidateRequest(req); err != nil {
		return nil, err
	}
	data, err := h.crud_city_uc.GetByPrimaryKey(int(req.Id))
	if err != nil {
		if errors.Is(err, crud_city_usecase.ErrRecordNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 2101, err.Error(), &grst_errors.ErrorDetail{})
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 2102, err.Error(), &grst_errors.ErrorDetail{})
	}

	return &pbCity.City{
		Id:         int32(data.Id),
		ProvinceId: int32(data.ProvinceId),
		Name:       data.Name,
	}, nil
}
func (h *CityHandler) Search(ctx context.Context, req *pbCity.SearchReq) (*pbCity.CityProfiles, error) {
	if err := pbCity.ValidateRequest(req); err != nil {
		return nil, err
	}

	searchResult, err := h.search_uc.Search(req.Keyword)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 2202, err.Error(), &grst_errors.ErrorDetail{})
	}
	resp := &pbCity.CityProfiles{
		Cities: []*pbCity.CityProfile{},
	}
	for _, data := range searchResult {
		resp.Cities = append(resp.Cities, &pbCity.CityProfile{
			Id:           int32(data.Id),
			Name:         data.Name,
			ProvinceId:   int32(data.ProvinceId),
			ProvinceName: data.ProvinceName,
		})
	}
	return resp, nil
}

func (h *CityHandler) FileDownload(ctx context.Context, req *empty.Empty) (*pbCity.FileDownloadResp, error) {

	result := &bytes.Buffer{}
	writer := csv.NewWriter(result)

	header := []string{"NIP", "Amount"}
	content := [][]string{
		{"00001", "120000"},
		{"00002", "140000"},
		{"00003", "150000"},
	}
	writer.Write(header)
	for _, val := range content {
		writer.Write(val)
	}
	writer.Flush()

	r_enc := base64.StdEncoding.EncodeToString(result.Bytes())
	return &pbCity.FileDownloadResp{
		Filename:    "test.csv",
		FileContent: string(r_enc),
	}, nil
}
