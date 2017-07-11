package command

import (
	"log"
	"strings"

	container "github.com/DaveBlooman/deliveroo/container"
	"github.com/DaveBlooman/deliveroo/docker"
)

type StopCommand struct {
	Meta
}

func (c *StopCommand) Run(args []string) int {

	hash := args[0]

	client, err := utils.NewDockerClient()
	if err != nil {
		log.Println("Unable to setup Docker client")
	}

	err = container.Stop(client, hash)
	if err != nil {
		log.Fatal("Unable to stop containers")
		return 1
	}

	return 0
}

func (c *StopCommand) Synopsis() string {
	return ""
}

func (c *StopCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
