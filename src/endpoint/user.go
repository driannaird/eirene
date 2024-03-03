package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/service"
)

type UserEndpoint interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type userendpoint struct {
	user service.UserService
}

func NewUserEndpoint(user service.UserService) UserEndpoint {
	return &userendpoint{
		user: user,
	}
}

func (u *userendpoint) Register(ctx *fiber.Ctx) error {
	var req entity.UserRegister
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("cannot parser body"))
	}

	data, err := u.user.Register(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success register", data))

}

func (u *userendpoint) Login(ctx *fiber.Ctx) error {
	var req entity.UserLogin
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("cannot parser body"))
	}

	data, err := u.user.Login(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success login", data))
}
func (u *userendpoint) Update(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize("token invalid"))
	}

	var req entity.User
	err = ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("cannot parser body"))
	}

	data, err := u.user.Update(claims.Username, req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success login", data))
}
