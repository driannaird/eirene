package service

import (
	"fmt"
	"os"

	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type ImageService interface {
	GetImage(username string) (*[]helper.Image, error)
	DeleteImage(username string, image string) error
}

type imageservivce struct {
	trace trace.Tracer
}

func NewImageService() ImageService {
	return &imageservivce{
		trace: otel.Tracer("image-service"),
	}
}

func (img *imageservivce) GetImage(username string) (*[]helper.Image, error) {
	span, err := util.Tracer(img.trace, "getAllImage")
	if err != nil {
		return nil, err
	}

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
	span, err := util.Tracer(img.trace, "deleteImage")
	if err != nil {
		return err
	}

	defer span.End()

	path := fmt.Sprintf("/data/image/%s/%s", username, image)
	err = os.Remove(path)
	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return helper.Success("success delete image", image)
}
