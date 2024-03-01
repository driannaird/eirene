package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/service"
)

type ModuleEndpoint interface {
	Install(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	AddSSHKey(ctx *fiber.Ctx) error
}

type moduleendpoint struct {
	mod service.ModuleService
}

func NewModuleEndpoint(mod service.ModuleService) ModuleEndpoint {
	return &moduleendpoint{
		mod: mod,
	}
}

func (m *moduleendpoint) Install(ctx *fiber.Ctx) error {
	var req entity.Module
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("something error went parser"))
	}

	data, err := m.mod.Install(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success install package", data))
}

func (m *moduleendpoint) Update(ctx *fiber.Ctx) error {
	var req entity.Module
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("something error went parser"))
	}

	err = m.mod.Update(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success update package", "update done"))
}

func (m *moduleendpoint) Delete(ctx *fiber.Ctx) error {
	var req entity.Module
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("something error went parser"))
	}

	data, err := m.mod.Delete(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success delete package", data))
}

func (m *moduleendpoint) AddSSHKey(ctx *fiber.Ctx) error {
	var req entity.SSHKey
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError("something error went parser"))
	}

	err = m.mod.AddSSHKey(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success append sshkey", nil))
}
