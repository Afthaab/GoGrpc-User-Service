package usecase

import (
	"github.com/profile/service/pkg/domain"
	interfaces "github.com/profile/service/pkg/repository/interface"
	useCase "github.com/profile/service/pkg/usecase/interface"
)

type UserRepo struct {
	Repo interfaces.UserRepo
}

func (u *UserRepo) ViewProfile(userid int64) (domain.User, error) {
	user := domain.User{}
	user, err := u.Repo.ViewProfile(userid)
	return user, err
}

func NewUserUseCase(repo interfaces.UserRepo) useCase.UserUseCase {
	return &UserRepo{
		Repo: repo,
	}
}
