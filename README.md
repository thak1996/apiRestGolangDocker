# Projeto Go com Docker

Este projeto é uma aplicação Go que utiliza Docker para facilitar a configuração do ambiente e a execução da aplicação.
## Requisitos

- Docker
- Docker Compose
- Configuração

## Instalação

Antes de iniciar, certifique-se de que você tem o Docker e o Docker Compose instalados em sua máquina.

Comandos para Executar o Projeto

### Passo 1: Construir os contêineres

Depois de limpar o ambiente, construa as imagens Docker e inicie os contêineres em segundo plano:
```
docker-compose up -d --build
```
### Passo 3: Verificar logs (opcional)

Para verificar se tudo está rodando conforme esperado, você pode inspecionar os logs dos contêineres:
```
docker-compose logs -f
```

### Passo 4: Parar a aplicação

Para parar a aplicação e remover os contêineres, utilize o comando:
```
docker-compose down
```
### Acesso à Aplicação

Após a execução bem-sucedida, a aplicação estará disponível na porta 8080. Você pode acessar a aplicação no seu navegador através do endereço:

http://localhost:8080

### Notas

Volumes: Os volumes são utilizados para persistir dados do banco de dados e outros dados importantes fora do ciclo de vida dos contêineres.
Scripts: O script wait-for-it.sh é utilizado para garantir que a aplicação Go só seja iniciada após o banco de dados estar acessível.
Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

Licença

Este projeto está sob a licença MIT.