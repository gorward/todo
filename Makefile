.PHONY: build dev run

build:
	go build ./cmd/webapi/webapi.go

dev: 
	CompileDaemon -color -build="go build ./cmd/webapi/webapi.go" -command="./webapi" 

run:
	go run ./cmd/webapi/webapi.go
