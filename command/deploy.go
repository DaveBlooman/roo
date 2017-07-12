package command

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DaveBlooman/deliveroo/builder"
	"github.com/DaveBlooman/deliveroo/clone"
	container "github.com/DaveBlooman/deliveroo/container"
	"github.com/DaveBlooman/deliveroo/docker"
	docker "github.com/fsouza/go-dockerclient"
)

type DeployCommand struct {
	Meta
}

func (c *DeployCommand) Run(args []string) int {

	if len(args) == 0 {
		log.Println("please provide a hash from the repo github.com/DaveBlooman/rubydockerapp")
		return 1
	}

	hash := args[0]

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
		return 1
	}
	directory := dir + "/" + hash + time.Now().String()

	err = clone.Fetch(directory, hash)
	if err != nil {
		log.Printf("Unable to clone snapshot - %v", err)
		return 1
	}

	client, err := utils.NewDockerClient()
	if err != nil {
		log.Println("Unable to setup Docker client")
	}

	log.Println("Building Image")

	imageName := "rubyapp:" + hash

	err = builder.Build(directory, hash, imageName, client)
	if err != nil {
		return 1
	}

	dbName := "postgres-" + hash

	err = container.Start(client, docker.CreateContainerOptions{
		Name: dbName,
		HostConfig: &docker.HostConfig{
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "localhost",
			Image:    "postgres:9.5.3",
			Labels:   map[string]string{hash: "true"},
		},
	})
	if err != nil {
		log.Println("Failed to start container" + dbName)
		return 1
	}

	err = container.Start(client, docker.CreateContainerOptions{
		Name: "ruby-app-" + hash,
		HostConfig: &docker.HostConfig{
			Links:           []string{dbName + ":postgres"},
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "deliveroo.localhost",
			Image:    imageName,
			Labels:   map[string]string{hash: "true"},
		},
	})
	if err != nil {
		log.Println("Failed to start container" + dbName)
		return 1
	}

	err = os.RemoveAll(directory)
	if err != nil {
		log.Println("Failed to cleanup directory")
		return 1
	}

	return 0
}

func (c *DeployCommand) Synopsis() string {
	return ""
}

func (c *DeployCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
