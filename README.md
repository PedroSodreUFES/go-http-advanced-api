![Go](https://img.shields.io/badge/Language-Go-00ADD8?logo=go)

## API de Leilão

### Contexto
+ API de leilão com tabela de usuários, produtos e lances
+ Uso de goroutines
+ Uso de websockets
+ Uso de channels

### Tecnologias usadas
+ Go Language
+ PostgreSQL
+ Chi framework
+ Pgx
+ Google UUID
+ Pacote json
+ Pacote errors
+ Pacote log
+ Pacote http
+ SQLc
+ Tern
+ Docker

### Como rodar o programa
```bash
docker-compose up -d
go mod tidy
go run ./cmd/api
```
OBS: Requer Go instalado em seu computador.

### Como parar o container Docker e excluir os dados
```bash
docker-compose down -v
```