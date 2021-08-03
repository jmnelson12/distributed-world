.PHONY: test
test:
	go test -race ./...

TAG ?= 0.0.1

build-docker-logger:
	docker build -t github.com/jmnelson12/distributed-world/logger:$(TAG) -f ./build/server/Dockerfile .
