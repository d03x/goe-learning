package api

import (
	"elearning/app/domain"
	"elearning/app/dto"
	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	userRepo domain.UserRepository
}

func NewAuth(app *fiber.App, userRepo domain.UserRepository) {
	api := Auth{
		userRepo: userRepo,
	}
	app.Get("/login", api.login)

}
func (ap Auth) login(ctx fiber.Ctx) error {
	data := ap.userRepo.FindAll()
	var response []dto.UserResponseDto
	for _, item := range data {
		response = append(response, dto.UserResponseDto{
			Email: item.Email,
			Id:    item.Id,
			Name:  item.Name,
		})
	}

	return ctx.Status(200).JSON(dto.ResponseWithData[[]dto.UserResponseDto](response))
}
