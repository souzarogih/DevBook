### Serviço WebApp

Rodando na porta 3000

`$ go mod init webapp`

### Bibliotecas

`$ go get github.com/gorilla/mux`
`$ go get github.com/joho/godotenv`
`$ go get github.com/gorilla/securecookie`

##### Futuras implementações

#### Botão de esqueci minha senha

- [ ] Funcionalidade de enviar e-mail para recuperar senha.
- [ ] Exibir o nome do usuário logado ao lado do botão sair.
- [ ] Permitir comentar na publicação de outros usuário.
- [ ] Aplicação rodar em Docker

#### Executando o FrontEnd
```bash
# Crie um arquivo .env no projeto e insira uma porta e um secret no arquivo.
  API_URL=URL do backend.
  APP_PORT=Porta que o frontend vai executar.
  HASH_KEY=Um hash para ser usado como chave de segurança.
  BLOCK_KEY=Um outro hash para ser usado no projeto.


# Agora inicie o serviço
$ go run main.go
```
#### Links uteis
- [ReactJs](https://reactjs.org)