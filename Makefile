run:
	go run ./cmd/main.go config.yml

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/invite -a -ldflags '-extldflags "-static"' ./cmd/main.go