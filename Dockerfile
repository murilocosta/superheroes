FROM golang:1.14-alpine AS builder
RUN apk add git
RUN git clone https://github.com/murilocosta/superheroes.git /go/src/github.com/murilocosta/superheroes
WORKDIR /go/src/github.com/murilocosta/superheroes
RUN go get -d -v ./...
RUN go build -o cmd/superhero pkg/superhero.go

FROM alpine
ADD configs/config.yml /config.yml
COPY --from=builder /go/src/github.com/murilocosta/superheroes/cmd/superhero /superhero
RUN ["chmod", "+x", "/superhero"]
ENTRYPOINT [ "/superhero" ]
