.PHONY: test
test:
	go test -race ./...

TAG ?= 0.0.2

build-docker-server:
	docker build -t github.com/jmnelson12/distributed-world/server:$(TAG) -f ./build/server/Dockerfile .
