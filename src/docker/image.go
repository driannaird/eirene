package docker

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/eirene/src/config"
	"github.com/rulanugrh/eirene/src/helper"
)

type Image struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Platform   string `json:"platform"`
}

type DockerImage interface {
	Create(req Image) error
	ListImage() (*[]helper.DockerImage, error)
	InspectImage(id string) (*helper.InspectDockerImage, error)
	DeleteImage(id string) error
}

type imageclient struct {
	client *docker.Client
	config *config.App
}

func NewDockerImage(client *docker.Client) DockerImage {
	return &imageclient{client: client, config: config.GetConfig()}
}

func (img *imageclient) Create(req Image) error {
	err := img.client.PullImage(docker.PullImageOptions{
		Repository: req.Repository,
		Platform:   req.Platform,
		Tag:        req.Tag,
	}, docker.AuthConfiguration{
		Username: img.config.Docker.Username,
		Password: img.config.Docker.Password,
		Email:    img.config.Docker.Email,
	})

	if err != nil {
		return helper.InternalServerError(err.Error())
	}

	return nil
}

func (img *imageclient) ListImage() (*[]helper.DockerImage, error) {
	data, err := img.client.ListImages(docker.ListImagesOptions{All: true})
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	var response []helper.DockerImage
	for _, images := range data {
		image := helper.DockerImage{
			ID:          images.ID,
			Tag:         images.RepoTags,
			Created:     images.Created,
			Size:        images.Size,
			VirtualSize: images.VirtualSize,
			Labels:      images.Labels,
		}

		response = append(response, image)
	}

	return &response, nil
}

func (img *imageclient) InspectImage(id string) (*helper.InspectDockerImage, error) {
	data, err := img.client.InspectImage(id)
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	response := helper.InspectDockerImage{
		ID:            data.ID,
		Tag:           data.RepoTags,
		Created:       data.Created,
		Container:     data.Container,
		Size:          data.Size,
		VirtualSize:   data.VirtualSize,
		OS:            data.OS,
		Architecture:  data.Architecture,
		DockerVersion: data.DockerVersion,
		Author:        data.Author,
	}

	return &response, nil
}

func (img *imageclient) DeleteImage(id string) error {
	err := img.client.RemoveImage(id)
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	return nil
}
