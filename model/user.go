package model

type User struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Name string `gorm:"varchar(255)" json:"name"`
	Email string `gorm:"unique;varchar(255)" json:"email"`
	Password string `gorm:"varchar(255)" json:"password"`
}