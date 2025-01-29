<h1 align="center" style="font-weight: bold;">Quanto Gasto?</h1>


Este é o backend do sistema de controle financeiro, desenvolvido em Golang. Ele fornece endpoints para gerenciamento de usuários, despesas e renda.

## Tecnologias Utilizadas

- Golang
- Gin (Framework para criação de APIs)
- GORM (ORM para Golang)
- PostgreSQL (Banco de dados)
- Docker
- Air (Hot reload para desenvolvimento)

## Configuração e Execução

### Pré-requisitos

- Go instalado ([Download Go](https://golang.org/dl/))
- PostgreSQL instalado e configurado
- Air instalado globalmente para hot reload (`go install github.com/cosmtrek/air@latest`)

### Configuração do Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```
DB_URL=postgres://usuario:senha@localhost:5432/nome_do_banco
DB_HOST=localhost
DB_PORT=5432
DB_USER=usuario
DB_PASSWORD=senha
DB_NAME=nome_do_banco
SECRETKEY=sua_secret_key

```

### Rodando o Projeto

Para iniciar o servidor com hot reload, utilize o Air:

```sh
air
```

Caso queira rodar sem o Air, use:

```sh
go run main.go
```

## Endpoints

### Autenticação e Usuários

- `POST /usuarios/` - Cadastra um novo usuário
- `POST /usuarios/foto` - Upload de foto de perfil

### Gastos e Renda (Requer Autenticação)

- `GET /:id` - Obtém dados de um usuário
- `PUT /gastos-fixos/:id` - Edita um gasto fixo
- `PUT /gastos-variaveis/:id` - Edita um gasto variável
- `DELETE /gastos-fixos/:id` - Remove um gasto fixo
- `DELETE /gastos-variaveis/:id` - Remove um gasto variável
- `POST /renda` - Adiciona renda
- `POST /gastos-fixos` - Adiciona gasto fixo
- `POST /gastos-variaveis` - Adiciona gasto variável
- `GET /resumo` - Obtém resumo financeiro

## Middleware de Autenticação

As rotas protegidas utilizam um middleware de autenticação para validar os tokens dos usuários antes de permitir o acesso.

---

## Contribuição 🤝

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

---

## Licença 📄

Este projeto está licenciado sob a MIT License.


