package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/docker"
	"github.com/rulanugrh/eirene/src/helper"
)

type DockerEndpoint interface {
	PullImage(ctx *fiber.Ctx) error
	DeleteImage(ctx *fiber.Ctx) error
	ImageHistory(ctx *fiber.Ctx) error
	ListImage(ctx *fiber.Ctx) error
	InspectImage(ctx *fiber.Ctx) error

	CreateContainer(ctx *fiber.Ctx) error
	ListContainer(ctx *fiber.Ctx) error
	DeleteContainer(ctx *fiber.Ctx) error
	InspectContainer(ctx *fiber.Ctx) error
	DownloadResourceContainer(ctx *fiber.Ctx) error
	ContainerLogs(ctx *fiber.Ctx) error
	PauseContainer(ctx *fiber.Ctx) error

	CreateVolume(ctx *fiber.Ctx) error
	ListVolume(ctx *fiber.Ctx) error
	InspectVolume(ctx *fiber.Ctx) error
	DeleteVolume(ctx *fiber.Ctx) error
}

type dockerendpoint struct {
	container docker.DockerContainer
	image     docker.DockerImage
	volume    docker.DockerVolume
}

func NewDockerEndpoint(container docker.DockerContainer, image docker.DockerImage, volume docker.DockerVolume) DockerEndpoint {
	return &dockerendpoint{
		container: container,
		image:     image,
		volume:    volume,
	}
}

func (d *dockerendpoint) PullImage(ctx *fiber.Ctx) error {
	var req docker.Image
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError(err.Error()))
	}

	err = d.image.Create(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(201).JSON(helper.Created("success pull image"))
}

func (d *dockerendpoint) DeleteImage(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	err := d.image.DeleteImage(param)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success delete image", param))
}

func (d *dockerendpoint) ImageHistory(ctx *fiber.Ctx) error {
	param := ctx.Params("name")
	data, err := d.image.ImageHistory(param)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("image history found", data))
}

func (d *dockerendpoint) ListImage(ctx *fiber.Ctx) error {
	data, err := d.image.ListImage()
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	if data == nil {
		return ctx.Status(404).JSON(helper.NotFound("sorry you not pull image"))
	}

	return ctx.Status(200).JSON(helper.Success("list your local image", data))
}

func (d *dockerendpoint) InspectImage(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	data, err := d.image.InspectImage(param)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("sorry image with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("image found", data))
}

func (d *dockerendpoint) CreateContainer(ctx *fiber.Ctx) error {
	var req docker.Container
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError(err.Error()))
	}

	data, err := d.container.Create(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success create container", data))
}

func (d *dockerendpoint) ListContainer(ctx *fiber.Ctx) error {
	data, err := d.container.ListContainer()
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	if data == nil {
		return ctx.Status(404).JSON(helper.NotFound("you not create container"))
	}

	return ctx.Status(200).JSON(helper.Success("list all container", data))
}

func (d *dockerendpoint) InspectContainer(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	data, err := d.container.InspectContainer(params)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("container with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("container found", data))
}

func (d *dockerendpoint) DeleteContainer(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	err := d.container.DeleteContainer(params)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("container with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("container successfull deleted", params))
}

func (d *dockerendpoint) ContainerLogs(ctx *fiber.Ctx) error {
	params := ctx.Params("name")
	err := d.container.ContainerLog(params, ctx.Request().BodyWriter())
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("container with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("container get logger container with this id", params))
}

func (d *dockerendpoint) PauseContainer(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	err := d.container.PauseContainer(params)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("container with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("container paused", params))
}

func (d *dockerendpoint) DownloadResourceContainer(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	err := d.container.DownloadResources(params, ctx.Request().BodyWriter())
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("container with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("success download resources", params))
}

func (d *dockerendpoint) CreateVolume(ctx *fiber.Ctx) error {
	var req docker.Volume
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(helper.InternalServerError(err.Error()))
	}

	data, err := d.volume.Create(req)
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(helper.Success("success create volume", data))
}

func (d *dockerendpoint) ListVolume(ctx *fiber.Ctx) error {
	data, err := d.volume.ListVolume()
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	if data == nil {
		return ctx.Status(404).JSON(helper.NotFound("data not found"))

	}

	return ctx.Status(200).JSON(helper.Success("list all volume", data))
}

func (d *dockerendpoint) InspectVolume(ctx *fiber.Ctx) error {
	params := ctx.Params("name")

	data, err := d.volume.InspectVolume(params)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("volume with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("volume found", data))
}

func (d *dockerendpoint) DeleteVolume(ctx *fiber.Ctx) error {
	params := ctx.Params("name")

	err := d.volume.DeleteVolume(params)
	if err != nil {
		return ctx.Status(404).JSON(helper.NotFound("volume with this id not found"))
	}

	return ctx.Status(200).JSON(helper.Success("succesfull deleted volume", params))
}
