package utils

import "testing"

func TestDockerClient(t *testing.T) {
	docker, err := NewDockerClient()
	if err != nil {
		t.Error(err)
	}
	if docker == nil {
		t.Error("docker is nil!")
	}
}
