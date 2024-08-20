# MyApiGo Project

Este é o projeto `MyApiGo`, uma aplicação back-end escrita em Go, configurada para ser executada dentro de contêineres Docker para fácil desenvolvimento e implantação.

## Estrutura do Projeto

O projeto está organizado da seguinte maneira:

/myapigo/
│
├── cmd/
│ └── myapi/
│ └── main.go # Entrada principal da aplicação Go
│
├── pkg/ # Pacotes reutilizáveis
│
├── internal/ # Código interno específico deste projeto
│
├── Dockerfile # Dockerfile para construir a imagem do contêiner
│
├── docker-compose.yml # Configuração do Docker Compose
│
├── go.mod # Gerenciamento de módulos e dependências Go
├── go.sum # Soma de verificação das dependências
│
└── README.md # Documentação do projeto

## Pré-requisitos

Antes de iniciar, certifique-se de ter instalado em sua máquina:

-   Docker
-   Docker Compose
-   Go (opcional para máquina local, necessário para configurações e testes)

## Configuração e Execução

### Construindo a Imagem Docker

Para construir a imagem Docker para a aplicação, execute:

```bash
docker-compose up -d --build
```

Após iniciar o contêiner, a aplicação estará disponivel em:

http://localhost:8080
