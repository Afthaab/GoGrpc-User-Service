package interfaces

import "github.com/profile/service/pkg/domain"

type UserRepo interface {
	FindProfile(user domain.User) (domain.User, error)
	EditProfile(user domain.User) int
	UpdatePassword(passwordData domain.Password) int
	CreateAddress(addressData domain.Address) (domain.Address, error)
}
