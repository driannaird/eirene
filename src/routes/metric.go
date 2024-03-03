package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
)

func MetricRoutes(f *fiber.App, endpoint endpoint.MetricEndpint) {
	met := f.Group("/api/v1/metric")
	met.Get("/", endpoint.GetAllMetric)
}
