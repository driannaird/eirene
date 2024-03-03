package docker

import (
	"context"
	"io"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type HostConfig struct {
	Binds       []string                             `json:"binding"`
	NetworkMode string                               `json:"network_mode"`
	PortBinding map[docker.Port][]docker.PortBinding `json:"port_binding"`
}

type ContainerConfig struct {
	Hostname   string                   `json:"hostname,omitempty"`
	Domainname string                   `json:"domain_name,omitempty"`
	Image      string                   `json:"image,omitempty"`
	Tty        bool                     `json:"tty"`
	OpenStdin  bool                     `json:"bool"`
	Env        []string                 `json:"env"`
	Port       map[docker.Port]struct{} `json:"port"`
}

type Container struct {
	Name       string           `json:"name"`
	Platform   string           `json:"platform"`
	Config     *ContainerConfig `json:"config"`
	HostConfig *HostConfig      `json:"host_config"`
}

type DockerContainer interface {
	Create(req Container) (*helper.Container, error)
	ListContainer() (*[]helper.ListContainer, error)
	InspectContainer(id string) (*helper.InspectContainer, error)
	DeleteContainer(id string) error
	ContainerLog(name string, w io.Writer) error
	DownloadResources(id string, w io.Writer) error
}

type container struct {
	client *docker.Client
	trace  trace.Tracer
}

func NewDockerContainer(client *docker.Client) DockerContainer {
	return &container{client: client, trace: otel.Tracer("container-service")}
}

func (c *container) Create(req Container) (*helper.Container, error) {
	span, err := util.Tracer(c.trace, "createContainer")
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	defer span.End()

	data, err := c.client.CreateContainer(docker.CreateContainerOptions{
		Name: req.Name,
		HostConfig: &docker.HostConfig{
			Binds:        req.HostConfig.Binds,
			NetworkMode:  req.HostConfig.NetworkMode,
			PortBindings: req.HostConfig.PortBinding,
		},
		Config: &docker.Config{
			Hostname:     req.Config.Hostname,
			Image:        req.Config.Image,
			Domainname:   req.Config.Domainname,
			Tty:          req.Config.Tty,
			Env:          req.Config.Env,
			ExposedPorts: req.Config.Port,
			OpenStdin:    req.Config.OpenStdin,
		},
		Platform: req.Platform,
	})

	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	response := helper.Container{
		ID:           data.ID,
		Image:        data.Image,
		Name:         data.Name,
		Created:      data.Created,
		Path:         data.Path,
		HostnamePath: data.HostnamePath,
		HostsPath:    data.HostsPath,
		Config: &helper.Config{
			Hostname:   data.Config.Hostname,
			Domainname: data.Config.Domainname,
			Image:      data.Config.Image,
			OpenStdin:  data.Config.OpenStdin,
			Tty:        data.Config.Tty,
			Env:        data.Config.Env,
			Port:       data.Config.ExposedPorts,
		},
	}

	return &response, nil
}

func (c *container) ListContainer() (*[]helper.ListContainer, error) {
	span, err := util.Tracer(c.trace, "listContainer")
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	defer span.End()

	data, err := c.client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	var response []helper.ListContainer
	for _, ct := range data {
		var listPort []helper.Port
		for _, ports := range ct.Ports {
			port := helper.Port{
				PrivatePort: ports.PrivatePort,
				PublicPort:  ports.PublicPort,
				IP:          ports.IP,
				Type:        ports.Type,
			}

			listPort = append(listPort, port)
		}

		res := helper.ListContainer{
			ID:      ct.ID,
			Command: ct.Command,
			Created: ct.Created,
			Ports:   listPort,
			Image:   ct.Image,
			Status:  ct.Status,
			State:   ct.State,
		}

		response = append(response, res)
	}

	return &response, nil
}

func (c *container) InspectContainer(id string) (*helper.InspectContainer, error) {
	span, err := util.TracerWithAttribute(c.trace, "inspectContainer", id)
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	defer span.End()

	data, err := c.client.InspectContainerWithOptions(docker.InspectContainerOptions{
		ID: id,
	})
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	response := helper.InspectContainer{
		ID:           data.ID,
		Image:        data.Image,
		Name:         data.Name,
		HostnamePath: data.HostnamePath,
		HostsPath:    data.HostsPath,
		Port:         data.Config.ExposedPorts,
		Env:          data.Config.Env,
		TTY:          data.Config.Tty,
		OpenStdin:    data.Config.OpenStdin,
	}

	return &response, nil
}

func (c *container) DeleteContainer(id string) error {
	span, err := util.TracerWithAttribute(c.trace, "deleteContainer", id)
	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	defer span.End()

	err = c.client.RemoveContainer(docker.RemoveContainerOptions{
		ID:      id,
		Context: context.Background(),
	})

	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return nil
}

func (c *container) ContainerLog(name string, w io.Writer) error {
	span, err := util.TracerWithAttribute(c.trace, "containerLog", name)
	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	defer span.End()

	err = c.client.Logs(docker.LogsOptions{
		Context:      context.Background(),
		Container:    name,
		OutputStream: w,
		Stdout:       true,
		RawTerminal:  true,
	})

	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return nil
}

func (c *container) DownloadResources(id string, w io.Writer) error {
	span, err := util.TracerWithAttribute(c.trace, "downloadResourceContainer", id)
	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	defer span.End()

	err = c.client.DownloadFromContainer(id, docker.DownloadFromContainerOptions{
		Path:              "./data/docker/",
		OutputStream:      w,
		Context:           context.Background(),
		InactivityTimeout: time.Duration(1 * time.Minute),
	})

	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return nil
}
