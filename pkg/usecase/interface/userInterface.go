package interfaces

import "github.com/profile/service/pkg/domain"

type UserUseCase interface {
	ViewProfile(userid int64) (domain.User, error)
}
