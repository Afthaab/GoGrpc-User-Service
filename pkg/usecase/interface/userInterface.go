package interfaces

import "github.com/profile/service/pkg/domain"

type UserUseCase interface {
	ViewProfile(user domain.User) (domain.User, error)
	EditProfile(user domain.User) error
	ChangePassword(passwordData domain.Password) error
	AddAddress(addressData domain.Address) (domain.Address, error)
	ViewAddress(addressData domain.Address) ([]domain.Address, error)
	EditAddress(addressData domain.Address) (domain.Address, error)
}
