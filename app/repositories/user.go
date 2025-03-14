package repositories

import (
	"elearning/app/domain"
	"elearning/app/models"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.UserRepository {
	return &user{
		db: db,
	}
}

func (a *user) FindAll() []models.Users {
	a.db.FirstOrCreate(&models.Users{
		Email: "dadan@gmail.com",
	})
	var user []models.Users
	a.db.Take(&user)
	return user
}
