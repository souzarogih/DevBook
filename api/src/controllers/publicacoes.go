package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "300")
		log.Printf("Ocorreu um erro ao tentar extrair o usuário - 300")
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro, "301")
		log.Printf("Ocorreu um erro ao tentar processar o corpo da requisição - 301")
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "302")
		log.Printf("Ocorreu um erro ao tentar converter os dados da requisição - 302")
		return
	}

	publicacao.AutorId = usuarioID
	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "303")
		log.Printf("Ocorreu um erro ao tentar preparar a publicação - 303")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "304")
		log.Printf("Ocorreu um erro ao tentar conectar com o banco de dados - 304")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "305")
		log.Printf("Ocorreu um erro ao tentar criar a publicação 305")
		return
	}

	log.Printf("Publicação com id %d criado com sucesso.", publicacao.ID)
	respostas.JSON(w, http.StatusCreated, publicacao)


}

// BuscarPublicacoes traz as publicações que apareceriam no feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// BuscarPublicacao traz única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao altera os dados de um apublicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

//DeletarPublicacao exclui os dados de uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}