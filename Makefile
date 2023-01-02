.PHONY: build
build:
    swag init -g ../../cmd/apiserver/main.go -d ./internal/app	
	go build -v ./cmd/apiserver

.PHONY: test
test:
	swag init -g ../../cmd/apiserver/main.go -d ./internal/app	
	go test -v -race -timeout 30s ./...
run:
	swag init -g ../../cmd/apiserver/main.go -d ./internal/app	
	go run -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
.DEFAULT_GOAL := build