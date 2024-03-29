package interfaces

import (
	"github.com/profile/service/pkg/domain"
)

type UserRepo interface {
	FindProfile(user domain.User) (domain.User, int64)
	EditProfile(user domain.User) int
	UpdatePassword(passwordData domain.Password) int64
	CreateAddress(addressData domain.Address) (domain.Address, error)
	ViewAllAddress(addressData domain.Address) ([]domain.Address, int64)
	ViewAddressByID(addressData domain.Address) (domain.Address, int64)
	EditAddress(addressData domain.Address) (domain.Address, int64)
	FindByUserName(user domain.User) (domain.User, int64)
	ViewAllUsers(userData domain.User) ([]domain.User, int64)
	BlockOrUnblockUser(userData domain.User) int64
}
