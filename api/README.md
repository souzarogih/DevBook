# Aplicação DevBook

## API DevBook

### Tecnologias

- [x] Golang
- [x] JWT Token
- [x] Middlewares
- [x] MySQL

### Comandos básicos para Golang

### Criar um módulo

`go mod init api`

### Atualizar o módulo da aplicação com os novos imports

`$ go mod tidy`

### Rodar a aplicação

`$ go run main.go`

### Gerar build da aplicação(arquivo executável)

`$ go build`

### Executar o build

`$ ./api.exe`

### Instalação de módulos

<ul>
  <ol>go get github.com/gorilla/mux</ol>
  <ol>go get github.com/joho/godotenv</ol>
  <ol>go get github.com/go-sql-driver/mysql</ol>
  <ol>go get github.com/badoux/checkmail</ol>
  <ol>go get golang.org/x/crypto/bcrypt</ol>
  <ol>go get github.com/dgrijalva/jwt-go</ol>
</ul>

##### Futuras implementações

#### Funcionalides

- [ ] Funcionalidade de comentário persistindo comentário no MongoDB
- [ ] Aplicação rodar em Docker

#### Executando o BackEnd
```bash
# Crie um arquivo .env no projeto e insira os campos:

DB_USUARIO=Usuário do banco de dados
DB_SENHA=Senha do usuário do banco de dados
DB_NOME=Nome do banco de dados
API_PORT=Porta que o serviço será executado
SECRET_KEY=Chave de seguração

# Agora inicie o serviço
$ go run main.go
```

#### Links uteis
- [Loggly](https://www.loggly.com/use-cases/logging-in-golang-how-to-start/)