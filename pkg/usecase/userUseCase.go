package usecase

import (
	"github.com/profile/service/pkg/domain"
	interfaces "github.com/profile/service/pkg/repository/interface"
	useCase "github.com/profile/service/pkg/usecase/interface"
)

type UserRepo struct {
	Repo interfaces.UserRepo
}

func (u *UserRepo) ViewProfile(user domain.User) (domain.User, error) {
	user, err := u.Repo.ViewProfile(user)
	return user, err
}

func NewUserUseCase(repo interfaces.UserRepo) useCase.UserUseCase {
	return &UserRepo{
		Repo: repo,
	}
}
