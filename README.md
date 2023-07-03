# PDI-01-2023

## subir banco de dados : 
docker-compose -f build/package/docker/docker-compose.yml up -d


## teste de cobertura :

go test ./internal/repository/users -coverprofile=coverage.out -covermode=count && goto -func=converage.out


## resultado dos Testes em HTML
go tool cover -html=coverage.out

## Rodar Aplicação 
go run cmd/pdi-go-kafka-db/main.go