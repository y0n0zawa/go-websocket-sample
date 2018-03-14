# Go websocket chat server

## Requirements

- Go (>= 1.10)
- dep (for dependency management)

### Install dependencies

```
$ dep ensure
```

### Update depdendencies

```
$ dep ensure -update
```

## Run test

```
$ go test ./...
```

## Run server (web and socket)
```
$ go run cmd/go-websocket-sample/main.go
```