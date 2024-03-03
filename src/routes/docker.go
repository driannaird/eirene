package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
)

func DockerRoutes(f *fiber.App, endpoint endpoint.DockerEndpoint) {
	ctn := f.Group("/api/v1/docker/container")
	ctn.Post("/create", endpoint.CreateContainer)
	ctn.Get("/getAll", endpoint.ListContainer)
	ctn.Get("/get/:id", endpoint.InspectContainer)
	ctn.Get("/resource/:id", endpoint.DownloadResourceContainer)
	ctn.Post("/pause/:id", endpoint.PauseContainer)
	ctn.Get("/logs/:id", endpoint.ContainerLogs)
	ctn.Delete("/delete/:id", endpoint.DeleteContainer)

	img := f.Group("/api/v1/docker/image")
	img.Post("/pull", endpoint.PullImage)
	img.Get("/getAll", endpoint.ListImage)
	img.Get("/get/:id", endpoint.InspectImage)
	img.Delete("/delete/:id", endpoint.DeleteImage)
	img.Get("/history/:name", endpoint.ImageHistory)

	vlm := f.Group("/api/v1/docker/volume")
	vlm.Post("/create", endpoint.CreateVolume)
	vlm.Get("/getAll", endpoint.ListVolume)
	vlm.Get("/get/:id", endpoint.InspectVolume)
	vlm.Delete("/delete/:id", endpoint.DeleteVolume)
}
