package repository

import (
	"github.com/profile/service/pkg/domain"
	interfaces "github.com/profile/service/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserDataBase struct {
	DB *gorm.DB
}

func (r *UserDataBase) ViewProfile(user domain.User) (domain.User, error) {
	result := r.DB.Raw("select * from users where id = ?", user.Id).Scan(&user).Error
	return user, result
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserDataBase{
		DB: db,
	}
}
