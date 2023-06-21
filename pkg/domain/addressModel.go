package domain

type Address struct {
	Addressid uint   `JSON:"addressid" gorm:"primarykey;unique"`
	User      User   `gorm:"ForeignKey:uid"`
	Uid       uint   `JSON:"uid"`
	Houseno   string `JSON:"houseno" gorm:"not null"`
	Area      string `JSON:"area" gorm:"not null"`
	Landmark  string `JSON:"landmark" gorm:"not null"`
	City      string `JSON:"city" gorm:"not null"`
	Type      string `JSON:"type" gorm:"not null"`
}

//Uid       uint   `JSON:"uid, omitempty"`
