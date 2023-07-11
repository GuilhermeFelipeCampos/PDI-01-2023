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

.PHONY: test-coverage
test-coverage: 
	go test ./internal/repository/users -coverprofile=coverage.out -covermode=count && goto -func=converage.out	

.PHONY: result-test-html
result-test-html:
	go tool cover -html=coverage.out
