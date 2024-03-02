package config

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/eirene/src/helper"
)

func DockerConnection() (*docker.Client, error) {
	docker, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	return docker, nil
}
