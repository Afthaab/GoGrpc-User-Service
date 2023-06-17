package interfaces

import "github.com/profile/service/pkg/domain"

type UserRepo interface {
	ViewProfile(user domain.User) (domain.User, error)
}
