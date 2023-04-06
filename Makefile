.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s -run TestSuite$$ ./...

.PHONY: docker
docker:
	docker compose -f ./build/docker-compose.local.yaml up -d

.DEFAULT_GOAL := build