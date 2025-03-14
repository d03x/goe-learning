package repositories

import (
	"elearning/app/domain"
	"elearning/app/models"
	"github.com/doug-martin/goqu/v9"
)

type user struct {
	db *goqu.Database
}

func NewUser(db *goqu.Database) domain.UserRepository {
	return &user{
		db: db,
	}
}

func (a *user) FindAll() ([]models.User, error) {
	var user []models.User
	err := a.db.From("users").ScanStructs(&user)
	if err != nil {
		return []models.User{}, err
	}
	return user, err
}
