package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
	"github.com/rulanugrh/eirene/src/internal/middleware"
)

func UserRoutes(f *fiber.App, endpoint endpoint.UserEndpoint) {
	f.Post("/register", endpoint.Register)
	f.Post("/login", endpoint.Login)

	u := f.Group("/api/v1/user")
	u.Use(middleware.JWTVerify())
	u.Put("/update", endpoint.Update)
}
