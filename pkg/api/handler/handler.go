package handler

import (
	"context"
	"net/http"

	"github.com/profile/service/pkg/domain"
	"github.com/profile/service/pkg/pb"
	interfaces "github.com/profile/service/pkg/usecase/interface"
)

// ------------------------------------ User Profile ------------------------------------------
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
		Error:    "No Error",
		Dob:      user.Dateofbirth,
		Gender:   user.Gender,
	}, nil
}

func (h *UserHandler) EditProfile(ctx context.Context, req *pb.EditProfileRequest) (*pb.EditProfileResponse, error) {
	user := domain.User{
		Id:          uint(req.Id),
		Username:    req.Username,
		Dateofbirth: req.Dob,
		Gender:      req.Gender,
	}
	err := h.useCase.EditProfile(user)
	if err != nil {
		return &pb.EditProfileResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Editing the user profile",
		}, err
	} else {
		return &pb.EditProfileResponse{
			Status: http.StatusOK,
			Error:  "nil",
		}, nil
	}
}

func (h *UserHandler) ChangePassword(ctx context.Context, req *pb.ChangeRequest) (*pb.ChangeResponse, error) {
	passwordData := domain.Password{
		Id:          uint(req.Id),
		Oldpassword: req.Oldpassword,
		Newpassword: req.Newpassword,
	}
	err := h.useCase.ChangePassword(passwordData)
	if err != nil {
		return &pb.ChangeResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Changing the Password",
		}, err
	} else {
		return &pb.ChangeResponse{
			Status: http.StatusOK,
			Error:  "nil",
		}, nil
	}
}

// ------------------------------------ Address Management ------------------------------------------

func (h *UserHandler) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	addressData := domain.Address{
		Uid:      uint(req.Id),
		Houseno:  req.Houseno,
		Area:     req.Area,
		Landmark: req.Landmark,
		City:     req.City,
		Type:     req.Type,
	}
	addressData, err := h.useCase.AddAddress(addressData)
	if err != nil {
		return &pb.AddAddressResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Adding the Address",
		}, err
	} else {
		return &pb.AddAddressResponse{
			Status: http.StatusOK,
			Error:  "nil",
			Addid:  int64(addressData.Addressid),
		}, nil
	}

}
