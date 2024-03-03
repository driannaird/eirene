package docker

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Volume struct {
	Name       string            `json:"name"`
	Driver     string            `json:"driver"`
	Labels     map[string]string `json:"label"`
	DriverOpts map[string]string `json:"driver_ops"`
}

type DockerVolume interface {
	Create(req Volume) (*helper.Volume, error)
	ListVolume() (*[]helper.Volume, error)
	InspectVolume(name string) (*helper.Volume, error)
	DeleteVolume(name string) error
}

type dockervolume struct {
	trace  trace.Tracer
	client *docker.Client
}

func NewDockerVolume(client *docker.Client) DockerVolume {
	return &dockervolume{client: client, trace: otel.Tracer("docker-volume")}
}

func (v *dockervolume) Create(req Volume) (*helper.Volume, error) {
	span, err := util.Tracer(v.trace, "createVolume")
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	defer span.End()

	data, err := v.client.CreateVolume(docker.CreateVolumeOptions{
		Driver:     req.Driver,
		Name:       req.Name,
		DriverOpts: req.DriverOpts,
		Context:    context.Background(),
		Labels:     req.Labels,
	})

	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	response := helper.Volume{
		Name:       data.Name,
		Driver:     data.Driver,
		Labels:     data.Labels,
		DriverOpts: data.Options,
	}

	return &response, nil
}

func (v *dockervolume) ListVolume() (*[]helper.Volume, error) {
	span, err := util.Tracer(v.trace, "listVolume")
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	defer span.End()

	var response []helper.Volume
	data, err := v.client.ListVolumes(docker.ListVolumesOptions{Context: context.Background()})
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	for _, dt := range data {
		result := helper.Volume{
			Name:       dt.Name,
			Driver:     dt.Driver,
			DriverOpts: dt.Options,
			Labels:     dt.Labels,
		}

		response = append(response, result)
	}

	return &response, nil
}

func (v *dockervolume) InspectVolume(name string) (*helper.Volume, error) {
	span, err := util.TracerWithAttribute(v.trace, "inspectVolume", name)
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	defer span.End()

	data, err := v.client.InspectVolume(name)

	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	response := helper.Volume{
		Name:       data.Name,
		Driver:     data.Driver,
		Labels:     data.Labels,
		DriverOpts: data.Options,
	}

	return &response, nil
}

func (v *dockervolume) DeleteVolume(name string) error {
	span, err := util.TracerWithAttribute(v.trace, "inspectVolume", name)
	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	defer span.End()

	err = v.client.RemoveVolumeWithOptions(docker.RemoveVolumeOptions{Name: name, Context: context.Background()})

	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return nil
}
