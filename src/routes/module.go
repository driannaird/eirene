package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
)

func ModuleRoutes(f *fiber.App, endpoint endpoint.ModuleEndpoint) {
	mod := f.Group("/api/v1/module")
	mod.Post("/install", endpoint.Install)
	mod.Put("/update", endpoint.Update)
	mod.Post("/add", endpoint.AddSSHKey)
	mod.Delete("/delete", endpoint.Delete)
}
