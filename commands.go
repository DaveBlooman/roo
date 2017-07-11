package main

import (
	"github.com/DaveBlooman/deliveroo/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"deploy": func() (cli.Command, error) {
			return &command.DeployCommand{
				Meta: *meta,
			}, nil
		},
		"stop": func() (cli.Command, error) {
			return &command.StopCommand{
				Meta: *meta,
			}, nil
		},
	}
}
