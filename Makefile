
GOPATH:=$(shell go env GOPATH)

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: build
build:
	go build -o cmd/superhero/superhero cmd/superhero/superhero.go

.PHONY: docker
docker:
	docker build . -t superhero:1.0.0
