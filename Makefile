SHELL := /bin/bash
.PHONY: run
run: 
	go run cmd/pdi-go-kafka-db/main.go

.PHONY: docker-down
docker-down:
	docker-compose -f build/package/docker/docker-compose.yml down

.PHONY: docker-up
docker-up: docker-down
	docker-compose -f build/package/docker/docker-compose.yml up -d

