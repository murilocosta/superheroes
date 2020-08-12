# Superheroes

A simple API to store your favorites heroes/villains made with [Go 1.14](https://golang.org/dl/)

## All the commands bellow need to be run from the root folder.

### To install the dependencies:
```
go get ./...
```

### To run the tests:
```
go test ./...
```

### To migrate the schema:
```
go run ./cmd/migrate/migrate.go
```

### To start the server:
```
go run ./cmd/superhero/superhero.go -config ./configs/config.yml
```
