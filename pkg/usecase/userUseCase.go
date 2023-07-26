package usecase

import (
	"errors"

	"github.com/profile/service/pkg/domain"
	interfaces "github.com/profile/service/pkg/repository/interface"
	useCase "github.com/profile/service/pkg/usecase/interface"
	utility "github.com/profile/service/pkg/utils"
)

type UserRepo struct {
	Repo interfaces.UserRepo
}

func (u *UserRepo) BlockOrUnblockUser(userData domain.User) error {
	err := u.Repo.BlockOrUnblockUser(userData)
	if err == 0 {
		return errors.New("could not block/unblock the user")
	}
	return nil
}

func (u *UserRepo) ViewAllUser(userData domain.User) ([]domain.User, error) {
	userDatas, result := u.Repo.ViewAllUsers(userData)
	if result == 0 {
		return userDatas, errors.New("Could do not fetch the data")
	}
	return userDatas, nil
}

func (u *UserRepo) ViewProfile(user domain.User) (domain.User, error) {
	user, result := u.Repo.FindProfile(user)
	if result == 0 {
		return user, errors.New("User details not found")
	}
	return user, nil
}

func (u *UserRepo) EditProfile(user domain.User) error {
	user, err := u.Repo.FindByUserName(user)
	if err != 0 {
		return errors.New("Username already exist")
	}

	result := u.Repo.EditProfile(user)
	if result == 0 {
		return errors.New("Could not update the user details")
	}
	return nil
}
func (u *UserRepo) ChangePassword(passwordData domain.Password) error {
	user := domain.User{
		Id: passwordData.Id,
	}
	// finding the userDetails through user id from middleware
	user, result := u.Repo.FindProfile(user)
	if result == 0 {
		return errors.New("User details not found")
	}

	// checking the entered old passwords
	if !utility.VerifyPassword(passwordData.Oldpassword, user.Password) {
		return errors.New("Current Password did not match")
	}

	// Hash the new password
	passwordData.Newpassword = utility.HashPassword(passwordData.Newpassword)

	//updating the password
	result = u.Repo.UpdatePassword(passwordData)
	if result == 0 {
		return errors.New("Could not update the new Password")
	}

	return nil
}

func (u *UserRepo) AddAddress(addressData domain.Address) (domain.Address, error) {
	addressData, err := u.Repo.CreateAddress(addressData)
	if err != nil {
		return addressData, errors.New("Could not create the Address")
	}
	return addressData, nil
}

func (u *UserRepo) ViewAddress(addressData domain.Address) ([]domain.Address, error) {
	var address []domain.Address
	address, result := u.Repo.ViewAllAddress(addressData)
	if result == 0 {
		return address, errors.New("Could not view the User Address")
	}
	return address, nil
}
func (u *UserRepo) EditAddress(addressData domain.Address) (domain.Address, error) {
	addressData, err := u.Repo.EditAddress(addressData)
	if err == 0 {
		return addressData, errors.New("Could not edit the address")
	}
	return addressData, nil
}
func (u *UserRepo) ViewAddressByID(addressData domain.Address) (domain.Address, error) {
	addressData, result := u.Repo.ViewAddressByID(addressData)
	if result == 0 {
		return addressData, errors.New("Could not view the address")
	}
	return addressData, nil
}
func NewUserUseCase(repo interfaces.UserRepo) useCase.UserUseCase {
	return &UserRepo{
		Repo: repo,
	}
}
