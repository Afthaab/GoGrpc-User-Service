package domain

type Address struct {
	Addressid       uint   `JSON:"addressid" gorm:"primarykey;unique"`
	User            User   `gorm:"ForeignKey:uid"`
	Uid             uint   `JSON:"uid"`
	Type            string `JSON:"type" gorm:"not null"`
	Locationaddress string `JSON:"locationaddress" gorm:"not null"`
	CompleteAddress string `JSON:"completeAddress" gorm:"not null"`
	Landmark        string `JSON:"landmark" gorm:"not null"`
	Floorno         string `JSON:"floorno" gorm:"not null"`
}

//Uid       uint   `JSON:"uid, omitempty"`
