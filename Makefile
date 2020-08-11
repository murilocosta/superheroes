
GOPATH:=$(shell go env GOPATH)

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: build
build:
	go build -o cmd/superhero pkg/superhero.go

.PHONY: docker
docker: build
	docker build . -t superhero:1.0.0
