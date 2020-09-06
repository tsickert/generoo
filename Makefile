.DEFAULT_GOAL :=  setup
.PHONY: build test run

setup: bootstrap

bootstrap:
	exec scripts/bootstrap.sh

build: go-build go-test

go-build:
	go build -v -x -o bin/generoo cmd/cli.go

go-run:
	go run cmd/cli.go

go-test:
	go test ./...

compile: compile-linux compile-darwin compile-windows

compile-linux: compile-linux-amd64 compile-linux-arm
compile-darwin: compile-darwin-amd64
compile-windows: compile-windows-amd64

compile-darwin-amd64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-w' -o bin/{{project_name_dashes}}-darwin-amd64 cmd/{{project_name_dashes}}-server/main.go

compile-linux-amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w' -o bin/{{project_name_dashes}}-linux-amd64 cmd/{{project_name_dashes}}-server/main.go

compile-linux-arm:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -ldflags '-w' -o bin/{{project_name_dashes}}-linux-arm cmd/{{project_name_dashes}}-server/main.go

compile-windows-amd64:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 build -ldflags '-w' -o bin/{{project_name_dashes}}-windows-amd64 cmd/{{project_name_dashes}}-server/main.go
