package repository

import (
	"github.com/profile/service/pkg/domain"
	interfaces "github.com/profile/service/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserDataBase struct {
	DB *gorm.DB
}

func (r *UserDataBase) ViewProfile(id int64) (domain.User, error) {
	userData := domain.User{}
	result := r.DB.Raw("select * from users where id = ?", id).Scan(&userData).Error
	return userData, result
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserDataBase{
		DB: db,
	}
}
