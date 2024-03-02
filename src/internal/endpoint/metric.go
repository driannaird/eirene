package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/service"
)

type MetricEndpint interface {
	GetAllMetric(ctx *fiber.Ctx) error
}

type mtrcendpoint struct {
	service service.IMetric
}

func NewMetricEndpoint(service service.IMetric) MetricEndpint {
	return &mtrcendpoint{service: service}
}

func (mtr *mtrcendpoint) GetAllMetric(ctx *fiber.Ctx) error {
	data, err := mtr.service.GetMetric()
	if err != nil {
		return ctx.Status(400).JSON(helper.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(data)
}
