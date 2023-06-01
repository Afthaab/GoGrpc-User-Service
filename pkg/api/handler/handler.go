package handler

import (
	"context"
	"net/http"

	"github.com/profile/service/pkg/domain"
	"github.com/profile/service/pkg/pb"
	interfaces "github.com/profile/service/pkg/usecase/interface"
)

type UserHandler struct {
	useCase interfaces.UserUseCase
	pb.ProfileManagementServer
}

func NewUserHandler(usecase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: usecase,
	}
}

func (h *UserHandler) ViewProfile(ctx context.Context, req *pb.ViewProfileRequest) (*pb.ViewProfileResponse, error) {
	user := domain.User{}
	user, err := h.useCase.ViewProfile(req.Id)
	if err != nil {
		return &pb.ViewProfileResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}
	return &pb.ViewProfileResponse{
		Status:   http.StatusOK,
		Id:       int64(user.Id),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
