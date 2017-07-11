package utils

import "github.com/fsouza/go-dockerclient"

func NewDockerClient() (*docker.Client, error) {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}

	return client, nil
}
