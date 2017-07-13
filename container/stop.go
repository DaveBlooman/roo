package container

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

func Stop(client *docker.Client, hash string) error {

	containers, err := ListContainers(hash, client)
	if err != nil {
		log.Println("Unable to list containers")
		return err
	}

	if len(containers) == 0 {
		log.Println("No containers running")
		return nil
	}

	log.Println("Stopping Containers")

	for _, container := range containers {
		err := client.RemoveContainer(docker.RemoveContainerOptions{ID: (container.ID), Force: true})
		if err != nil {
			log.Fatal("Unable to stop container " + container.ID)
			return err
		}
		log.Printf(`Removed "%v"`+"\n", container.Names[0])
	}
	return nil
}

func ListContainers(hash string, client *docker.Client) ([]docker.APIContainers, error) {

	hashContainers := map[string][]string{
		"label": []string{
			hash,
		},
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{Filters: hashContainers})
	if err != nil {
		return nil, err
	}

	return containers, nil
}
