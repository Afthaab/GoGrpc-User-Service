package interfaces

import "github.com/profile/service/pkg/domain"

type UserUseCase interface {
	ViewProfile(user domain.User) (domain.User, error)
}
