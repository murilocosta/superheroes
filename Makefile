
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: 
	go build -o cmd/superhero pkg/superhero.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build

	docker build . -t superhero:1.0.0
