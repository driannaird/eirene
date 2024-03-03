package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rulanugrh/eirene/src/config"
	"github.com/rulanugrh/eirene/src/docker"
	"github.com/rulanugrh/eirene/src/endpoint"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/repository"
	"github.com/rulanugrh/eirene/src/internal/service"
	"github.com/rulanugrh/eirene/src/routes"
)

func main() {
	app := fiber.New()
	getConfig := config.GetConfig()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Accept, Content-Type, Authorization",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodDelete,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodOptions,
		}, ""),
		AllowOrigins: "*",
	}))

	file, err := os.OpenFile("./log/fiber.log", os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	app.Use(logger.New(logger.Config{
		Format:   "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone: "Asia/Jakarta",
		Output:   file,
	}))

	validate := middleware.NewValidation()
	dbConn, err := config.GetDB()
	if err != nil {
		log.Println("error cant connect to db")
	}

	getDocker, err := config.DockerConnection()
	if err != nil {
		log.Println("error cant connect to docker client")
	}

	userRepository := repository.NewUserRepository(dbConn)
	mailRepository := repository.NewMailRepository(dbConn)

	userService := service.NewUserService(userRepository, validate)
	mailService := service.NewMailService(mailRepository, validate)
	metricService := service.NewMetric()
	fileService := service.NewFileService()
	imageService := service.NewImageService()

	dockerContainer := docker.NewDockerContainer(getDocker)
	dockerImage := docker.NewDockerImage(getDocker)
	dockerVolume := docker.NewDockerVolume(getDocker)

	userEndpoint := endpoint.NewUserEndpoint(userService)
	mailEndpoint := endpoint.NewMailEndpoint(mailService)
	metricEndpoint := endpoint.NewMetricEndpoint(metricService)
	fileEndpoint := endpoint.NewFileEndpoint(fileService)
	imageEndpoint := endpoint.NewImageEndpoint(imageService)
	dockerEndpoint := endpoint.NewDockerEndpoint(dockerContainer, dockerImage, dockerVolume)

	routes.UserRoutes(app, userEndpoint)
	routes.MailRoutes(app, mailEndpoint)
	routes.FileRoutes(app, fileEndpoint)
	routes.MetricRoutes(app, metricEndpoint)
	routes.ImageRoutes(app, imageEndpoint)
	routes.DockerRoutes(app, dockerEndpoint)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", getConfig.Server.Host, getConfig.Server.Port)))
}
