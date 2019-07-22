[![CircleCI](https://circleci.com/gh/jhonata-menezes/luizalabs-desafio.svg?style=svg)](https://circleci.com/gh/jhonata-menezes/luizalabs-desafio)

#### Desafio Luiza Labs
Este projeto é um desafio proposto pela LuizaLabs. Para resolver o problema proposto estou usando [Golang](https://golang.org/), [VueJS](https://vuejs.org/) e [PostgreSQL](https://www.postgresql.org/).

### Executando com docker-compose
```bash
docker-compose up -d
```
- http://localhost:8000 - api
- http://localhost:8001 - interface de administração

### instalando as dependencias

```bash
cd front && npm i
```

### instalando postgres com docker-compose
```bash
docker-compose up -d
```

### Executando no ambiente de desenvolvimento
```bash
cd front && npm run dev
```
```bash
cd back && go run main/app.go
```

### Executando os testes
```bash
cd back/ && go test ./...
```

```bash
cd front/ && npm run test:unit
```



