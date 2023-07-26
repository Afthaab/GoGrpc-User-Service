package domain

type User struct {
	Id          uint   `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username    string `json:"username" gorm:"unique" validate:"required,min=2,max=50"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	Phone       string `json:"phone" validate:"required,len=10"`
	Profile     string `json:"profile"`
	Dateofbirth string `json:"dateofbirth" gorm:"default:null"`
	Gender      string `json:"gender" gorm:"default:null" `
	Isblocked   bool   `json:"isblocked" gorm:"default:false"`
}

type Password struct {
	Id          uint   `json:"id"`
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}
