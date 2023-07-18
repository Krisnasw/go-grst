package handler

import (
	"context"
	"errors"
	"net/http"

	profile_usecase "github.com/krisnasw/go-grst/examples/users-api/app/usecase/profile"
	pbUsers "github.com/krisnasw/go-grst/examples/users-api/handler/grst/users"
	grst_errors "github.com/krisnasw/go-grst/grst/errors"
	"google.golang.org/grpc/codes"
)

type UsersHandler struct {
	pbUsers.UnimplementedUsersServer
	profileUsecase profile_usecase.UseCase
}

func NewHandler(profileUsecase profile_usecase.UseCase) pbUsers.UsersServer {
	return &UsersHandler{pbUsers.UnimplementedUsersServer{}, profileUsecase}
}

func (h *UsersHandler) GetProfile(ctx context.Context, req *pbUsers.GetProfileReq) (*pbUsers.UserProfile, error) {
	if err := pbUsers.ValidateRequest(req); err != nil {
		return nil, err
	}

	res, err := h.profileUsecase.GetProfile(int(req.Id))
	if err != nil {
		if errors.Is(err, profile_usecase.ErrRecordNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 1101, err.Error(), &grst_errors.ErrorDetail{})
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 1102, err.Error(), &grst_errors.ErrorDetail{})
	}
	return res.ToPbUserProfile(), nil
}
