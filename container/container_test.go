package container

import (
	"testing"

	"github.com/DaveBlooman/deliveroo/docker"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/stretchr/testify/assert"
)

func TestContainers(t *testing.T) {
	client, err := utils.NewDockerClient()
	if err != nil {
		t.Error(err)
	}

	hash := "f8e9f807c23bd12b81e30ae57eb9619404aed110"

	opts := docker.CreateContainerOptions{
		Name:       "test-app-" + hash,
		HostConfig: &docker.HostConfig{},
		Config: &docker.Config{
			Image:  "davey/hello",
			Labels: map[string]string{hash: "true"},
		},
	}

	err = Start(client, opts)
	assert.NoError(t, err)
	err = Stop(client, hash)
	assert.NoError(t, err)

	containers, err := ListContainers(hash, client)
	assert.NoError(t, err, "list containers")
	assert.Equal(t, 0, len(containers))

}
