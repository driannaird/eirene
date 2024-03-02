package service

import (
	"context"
	"fmt"
	"os"

	"github.com/rulanugrh/eirene/src/helper"
	"go.opentelemetry.io/otel/trace"
)

type ImageService interface {
	GetImage(username string) (*[]helper.Image, error)
	DeleteImage(username string, image string) error
}

type imageservivce struct {
	trace trace.Tracer
}

func NewImageService(trace trace.Tracer) ImageService {
	return &imageservivce{
		trace: trace,
	}
}

func (img *imageservivce) GetImage(username string) (*[]helper.Image, error) {
	_, span := img.trace.Start(context.Background(), "getAllImage")
	defer span.End()

	path := fmt.Sprintf("./data/image/%s", username)
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	var response []helper.Image
	for _, f := range files {
		image := helper.Image{
			File: f.Name(),
			Link: fmt.Sprintf("./data/image/%s/%s", username, f.Name()),
		}

		response = append(response, image)
	}

	return &response, nil
}

func (img *imageservivce) DeleteImage(username string, image string) error {
	_, span := img.trace.Start(context.Background(), "delete-image")
	defer span.End()

	path := fmt.Sprintf("/data/image/%s/%s", username, image)
	err := os.Remove(path)
	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return helper.Success("success delete image", image)
}
