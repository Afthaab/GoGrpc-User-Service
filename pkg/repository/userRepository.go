package repository

import (
	"github.com/profile/service/pkg/domain"
	interfaces "github.com/profile/service/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserDataBase struct {
	DB *gorm.DB
}

func (r *UserDataBase) FindProfile(user domain.User) (domain.User, error) {
	result := r.DB.Raw("select * from users where id = ?", user.Id).Scan(&user).Error
	return user, result
}

func (r *UserDataBase) EditProfile(user domain.User) int {
	result := r.DB.Exec("UPDATE users SET username = ?, dateofbirth = ?, gender = ? where id = ?", user.Username, user.Dateofbirth, user.Gender, user.Id)
	return int(result.RowsAffected)
}

func (r *UserDataBase) UpdatePassword(passwordData domain.Password) int {
	result := r.DB.Exec("UPDATE users SET password = ? where id = ?", passwordData.Newpassword, passwordData.Id)
	return int(result.RowsAffected)
}

func (r *UserDataBase) CreateAddress(addressData domain.Address) (domain.Address, error) {
	result := r.DB.Create(&addressData)
	return addressData, result.Error
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserDataBase{
		DB: db,
	}
}
