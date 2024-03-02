package service

import (
	"fmt"
	"os"

	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type FileService interface {
	GetAll(username string) (*[]helper.File, error)
	Delete(username string, name string) error
}

type fileservice struct {
	trace trace.Tracer
}

func NewFileService() FileService {
	return &fileservice{trace: otel.Tracer("file-service")}
}

func (f *fileservice) GetAll(username string) (*[]helper.File, error) {
	span, err := util.Tracer(f.trace, "getAllFile")
	if err != nil {
		return nil, err
	}

	defer span.End()

	path := fmt.Sprintf("./data/file/%s", username)
	file, err := os.ReadDir(path)
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	var response []helper.File
	for _, data := range file {
		files := helper.File{
			File: data.Name(),
			Link: fmt.Sprintf("/data/file/%s/%s", username, data.Name()),
		}

		response = append(response, files)
	}

	return &response, nil
}

func (f *fileservice) Delete(username string, name string) error {
	span, err := util.Tracer(f.trace, "deleteFile")
	if err != nil {
		return err
	}

	defer span.End()

	path := fmt.Sprintf("./data/file/%s/%s", username, name)
	err = os.Remove(path)
	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return helper.Success("success delete image", name)
}
