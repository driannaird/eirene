package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
)

func UserRoutes(f *fiber.App, endpoint endpoint.UserEndpoint) {
	u := f.Group("/api/v1/user")
	u.Post("/register", endpoint.Register)
	u.Post("/login", endpoint.Login)
	u.Put("/update", endpoint.Update)
}
