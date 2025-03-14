package domain

import "elearning/app/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
}

type UserService interface {
}

type UserAuthService interface {
}
