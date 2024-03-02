package endpoint

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/service"
)

type FileEndpoint interface {
	Save(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type fileendpoint struct {
	file service.FileService
}

func NewFileEndpoint(file service.FileService) FileEndpoint {
	return &fileendpoint{
		file: file,
	}
}

func (f *fileendpoint) Save(ctx *fiber.Ctx) error {
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	path := fmt.Sprintf("/data/file/%s/%s", token.Username, file.Filename)
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

func (f *fileendpoint) GetAll(ctx *fiber.Ctx) error {
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	files, err := f.file.GetAll(token.Username)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success get all files", files))
}

func (f *fileendpoint) GetOne(ctx *fiber.Ctx) error {
	file := ctx.Params("file")
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	path := fmt.Sprintf("/data/file/%s/%s", token.Username, file)
	return ctx.Status(200).SendFile(path)
}

func (f *fileendpoint) Delete(ctx *fiber.Ctx) error {
	file := ctx.Params("file")
	claim := ctx.Get("Authorization")
	token, err := middleware.CheckToken(claim)
	if err != nil {
		return ctx.Status(401).JSON(helper.Unauthorize(err.Error()))
	}

	err = f.file.Delete(token.Username, file)
	if err != nil {
		return ctx.Status(500).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success delete image", file))
}
