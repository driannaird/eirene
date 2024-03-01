package endpoint

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/service"
)

type MailEndpoint interface {
	Inbox(ctx *fiber.Ctx) error
	Sent(ctx *fiber.Ctx) error
	Starred(ctx *fiber.Ctx) error
	Archive(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type mailendpoint struct {
	mail service.MailService
}

func NewMailEndpoint(mail service.MailService) MailEndpoint {
	return &mailendpoint{
		mail: mail,
	}
}

func (m *mailendpoint) Inbox(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize("token invalid"))
	}

	data, err := m.mail.Inbox(claims.Username)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("data not found"))
	}

	return ctx.Status(200).JSON(helper.Success("success email found", data))
}

func (m *mailendpoint) Sent(ctx *fiber.Ctx) error {
	var req entity.Mail
	token := ctx.Get("Authorization")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize("token invalid"))
	}

	req.UserEmail = claims.Email
	err = ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("cannot parser request"))
	}

	data, err := m.mail.Sent(req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("sorry something error"))
	}

	return ctx.Status(200).JSON(helper.Success("success sent email", data))

}

func (m *mailendpoint) Starred(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize("token invalid"))
	}

	data, err := m.mail.Starred(claims.Username)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("data not found"))
	}

	return ctx.Status(200).JSON(helper.Success("success email found", data))
}

func (m *mailendpoint) Archive(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize("token invalid"))
	}

	data, err := m.mail.Archived(claims.Username)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("data not found"))
	}

	return ctx.Status(200).JSON(helper.Success("success email found", data))
}

func (m *mailendpoint) Delete(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("cannot convert to int"))
	}

	err = m.mail.Delete(uint(id))
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("sorry cant delete mail"))
	}

	return ctx.Status(200).JSON(helper.Success("success email found", nil))

}

func (m *mailendpoint) Update(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("cannot convert to int"))
	}

	var req entity.Mail
	err = ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("cannot parser request"))
	}

	data, err := m.mail.Update(uint(id), req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest("cant update mail"))
	}

	return ctx.Status(200).JSON(helper.Success("success update mail", data))

}
