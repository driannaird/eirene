package endpoint

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/service"
)

type ImageEndpoint interface {
	Save(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type imgendpoint struct {
	image service.ImageService
}

func NewImageEndpoint(image service.ImageService) ImageEndpoint {
	return &imgendpoint{image: image}
}

func (img *imgendpoint) Save(ctx *fiber.Ctx) error {
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	path := fmt.Sprintf("/data/image/%s/%s", token.Username, file.Filename)
	if err = ctx.SaveFile(file, path); err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError(err.Error()))
	}

	response := helper.Response{
		Code:    200,
		Message: "success save file",
		Data:    file.Filename,
	}

	return ctx.Status(200).JSON(response)
}

func (img *imgendpoint) GetAll(ctx *fiber.Ctx) error {
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	response, err := img.image.GetImage(token.Username)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("sucess find all image", response))
}

func (img *imgendpoint) GetOne(ctx *fiber.Ctx) error {
	image := ctx.Params("img")
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	path := fmt.Sprintf("/data/image/%s/%s", token.Username, image)
	return ctx.Status(200).SendFile(path)
}

func (img *imgendpoint) Delete(ctx *fiber.Ctx) error {
	image := ctx.Params("img")
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	err = img.image.DeleteImage(token.Username, image)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))

	}

	return ctx.Status(200).JSON(helper.Success("success delete image", image))
}
