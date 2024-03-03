package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
)

func ImageRoutes(f *fiber.App, endpoint endpoint.ImageEndpoint) {
	f.Static("/data/image", "./data/image")

	img := f.Group("/api/v1/image")
	img.Post("/upload", endpoint.Save)
	img.Get("/getAll", endpoint.GetAll)
	img.Get("/get/:img", endpoint.GetOne)
	img.Delete("/delete/:img", endpoint.Delete)
}
