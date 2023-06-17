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
	user := domain.User{
		Id: uint(req.Id),
	}
	user, err := h.useCase.ViewProfile(user)
	if err != nil {
		return &pb.ViewProfileResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}
	return &pb.ViewProfileResponse{
		Status:   http.StatusOK,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Profile:  user.Profile,
	}, nil
}
