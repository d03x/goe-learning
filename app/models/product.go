package models

type Products struct {
	Id    uint `gorm:"primarykey;default:auto_random()"`
	Code  string
	Price uint
	Title string
	Size  uint
}
