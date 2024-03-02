package docker

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/eirene/src/helper"
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
}

type container struct {
	client *docker.Client
	trace  trace.Tracer
}

func NewDockerContainer(client *docker.Client) DockerContainer {
	return &container{client: client, trace: otel.Tracer("container-service")}
}

func (c *container) Create(req Container) (*helper.Container, error) {
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
