package command

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DaveBlooman/roo/clone"
	container "github.com/DaveBlooman/roo/container"
)

type DeployCommand struct {
	Meta
}

func (c *DeployCommand) Run(args []string) int {

	if len(args) == 0 {
		log.Println("please provide a hash from the repo github.com/DaveBlooman/go-app")
		return 1
	}

	hash := args[0]

	client, err := utils.NewDockerClient()
	if err != nil {
		log.Println("Unable to setup Docker client")
	}

	exists, err := container.ListContainers(hash, client)
	if err != nil {
		log.Println("Error determining if containers are running")
		return 1
	}
	if len(exists) > 0 {
		log.Println("snapshot already running")
		return 1
	}

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

	err = container.CreateContainers(hash, directory, client)
	if err != nil {
		log.Println("Unable to start containers, is Docker running?")
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
	return "Starts app and database"
}

func (c *DeployCommand) Help() string {
	helpText := `
There is a single argument to the deploy command, a git hash.

`
	return strings.TrimSpace(helpText)
}
