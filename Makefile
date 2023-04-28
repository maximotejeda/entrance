export DEV=1
include .env
export 
OS:=${shell go env GOOS}
ARCH=$(shell go env GOARCH)
OOSS="linux"
ARRCHS="arm 386"
DEBUG=1
SERVICE=entrance
.PHONY: all clean

image:
	docker build -t $(SERVICE)-service:latest --build-arg OS=$(OS) --build-arg ARCH=$(ARCH)  .
run-image: image
	@docker run -p 4000:4000 -p 8083:8083 $(SERVICE)-service:latest

image-debug:
	docker build -f ./Dockerfile-debug -t $(SERVICE)-service:debug --build-arg OS=$(OS) --build-arg ARCH=$(ARCH)  .
run-debug: image-debug
	@docker run -p 4000:4000 -p 8080:8080 --env-file .env $(SERVICE)-service:debug

run-local:build
	@./bin/$(SERVICE)-$(OS)-$(ARCH)
build: clean
	@mkdir bin 
	@go build -o bin/$(SERVICE)-$(OS)-$(ARCH) ./main/

test:
	@go test ./...
clean:
	@rm -rf bin
