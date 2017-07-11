package builder

import (
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

func Build(directory, hash, imageName string, client *docker.Client) error {

	err := client.BuildImage(docker.BuildImageOptions{
		Name:         imageName,
		Dockerfile:   "Dockerfile",
		OutputStream: os.Stdout,
		ContextDir:   directory,
	})
	if err != nil {
		log.Printf("Cannot Build Docker image - %v", err)
		return err
	}
	return nil
}
