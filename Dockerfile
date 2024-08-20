FROM golang:1.20-alpine

WORKDIR /app

# Instalar curl e bash
RUN apk add --no-cache curl bash

# Copiar os arquivos go.mod e go.sum e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte para o diretório de trabalho
COPY . .

# Compilar o aplicativo para produção
RUN go build -o myapi ./cmd/myapi/main.go

# Baixar o script wait-for-it.sh e torná-lo executável
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Definir o ponto de entrada para esperar o banco de dados e depois iniciar a aplicação
ENTRYPOINT ["/bin/bash", "/wait-for-it.sh", "db:3307", "--", "/app/myapi"]

# Expor a porta que o aplicativo usará
EXPOSE 8080
