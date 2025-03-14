package domain

import "elearning/app/models"

type UserRepository interface {
	FindAll() []models.Users
}

type UserService interface {
}

type UserAuthService interface {
}
