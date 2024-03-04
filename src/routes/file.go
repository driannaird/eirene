package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
	"github.com/rulanugrh/eirene/src/internal/middleware"
)

func FileRoutes(f *fiber.App, endpoint endpoint.FileEndpoint) {
	f.Static("/data/file", "./data/file")

	fl := f.Group("/api/v1/file")
	fl.Use(middleware.JWTVerify())
	fl.Post("/upload", endpoint.Save)
	fl.Get("/getAll", endpoint.GetAll)
	fl.Get("/get/:file", endpoint.GetOne)
	fl.Delete("/delete/:file", endpoint.Delete)
}
