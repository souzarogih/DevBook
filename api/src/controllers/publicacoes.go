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
		respostas.Erro(w, http.StatusUnauthorized, erro, "")
		log.Printf("")
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro, "")
		log.Printf("")
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "")
		log.Printf("")
		return
	}

	publicacao.AutorId = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "")
		log.Printf("")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "")
		log.Printf("")
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

//DeletarPublicao exclui os dados de uma publicação
func DeletarPublicao(w http.ResponseWriter, r *http.Request) {

}