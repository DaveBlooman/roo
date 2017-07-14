package container

import (
	"log"

	"github.com/DaveBlooman/roo/builder"
	docker "github.com/fsouza/go-dockerclient"
)

func CreateContainers(hash, directory string, client *docker.Client) error {
	imageName := "go-app:" + hash
	err := builder.Build(directory, hash, imageName, client)
	if err != nil {
		return err
	}

	dbName := "postgres-" + hash
	err = Start(client, docker.CreateContainerOptions{
		Name: dbName,
		HostConfig: &docker.HostConfig{
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "localhost",
			Image:    "postgres:9.5.3",
			Env:      []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres", "POSTGRES_DB=test", "- PGPASSWORD=postgres"},
			Labels:   map[string]string{hash: "true"},
		},
	})
	if err != nil {
		log.Println("Failed to start container" + dbName)
		return err
	}

	err = Start(client, docker.CreateContainerOptions{
		Name: "go-app-" + hash,
		HostConfig: &docker.HostConfig{
			Links:           []string{dbName + ":db"},
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "localhost",
			Image:    imageName,
			Labels:   map[string]string{hash: "true"},
		},
	})
	if err != nil {
		log.Println("Failed to start container" + dbName)
		return err
	}

	return nil
}
