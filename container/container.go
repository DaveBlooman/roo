package container

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

func Start(client *docker.Client, opts docker.CreateContainerOptions) error {
	container, err := client.CreateContainer(opts)
	if err != nil {
		return err
	}
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
