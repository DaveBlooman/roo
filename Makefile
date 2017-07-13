tests:
	docker pull davey/hello
	go test `glide novendor`

build:
	go build
