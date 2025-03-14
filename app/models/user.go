package models

type Users struct {
	Id             uint `gorm:"primarykey;default:auto_random()"`
	Name           string
	Age            uint
	Email          string `gorm:"unique"`
	Address        string
	Password       string
	ProfilePicture string
}
