package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "")
		log.Printf("")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "304")
		log.Printf("Ocorreu um erro ao tentar conectar com o banco de dados - 304")
		return
	}
	defer db.Close()

	fmt.Println("usuario>", usuarioID)
	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.Buscar(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "")
		log.Printf("")
		return
	}

	log.Printf("")
	respostas.JSON(w, http.StatusOK, publicacoes)
}

// BuscarPublicacao traz única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "306")
		log.Printf("Erro ao extrair o campo publicacaoId da requisição - 306")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "304")
		log.Printf("Erro interno ao tentar conectar com o banco de dados - 304")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "305")
		log.Printf("Ocorreu um erro interno ao tentar localizar a publicação - 305")
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)
}

// AtualizarPublicacao altera os dados de um apublicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	log.Printf("repositorios - AtualizarPublicacao - Executando a atualização de uma publicação.")

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "306")
		log.Printf("Erro ao tentar identificar o usuário da requisição. - 306")
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "307")
		log.Printf("Não foi possível converter o campo publicacaoId da requisição. - 307")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "308")
		log.Printf("Erro interno ao tentar conectar com o banco de dados. - 308")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "309")
		log.Printf("Erro interno ao tentar conectar com o banco de dados. - 309")
		return
	}

	if publicacaoSalvaNoBanco.AutorId != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua."), "310")
		log.Printf("Usuário está tentando alterar uma publicação que não criou. - 310")
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro, "311")
		log.Printf("Erro ao tentar processar a requisição enviada. - 311")
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "312")
		log.Printf("Não foi possível converter o corpo da requisição. - 312")
		return
	}

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "313")
		log.Printf("Ocorreu um erro ao tentar preparar a atualização - 313")
		return
	}

	if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "314")
		log.Printf("Não foi possível atualizar a publicação - 314")
		return
	}

	log.Printf("Atualização da publicação realizada com sucesso.")
	respostas.JSON(w, http.StatusNoContent, nil)
}

//DeletarPublicacao exclui os dados de uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	log.Printf("repositorios - DeletarPublicacao - Executando a deleção de uma publicação.")

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "311")
		log.Printf("Erro ao tentar identificar o usuário da requisição. - 311")
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "312")
		log.Printf("Não foi possível converter o campo publicacaoId da requisição. - 312")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "313")
		log.Printf("Erro interno ao tentar conectar com o banco de dados. - 313")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "314")
		log.Printf("Erro interno ao tentar conectar com o banco de dados. - 314")
		return
	}

	if publicacaoSalvaNoBanco.AutorId != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar uma publicação que não seja sua."), "315")
		log.Printf("Usuário está tentando deletar uma publicação que não criou. - 315")
		return
	}

	if erro = repositorio.Deletar(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "316")
		log.Printf("Ocorreu um erro interno ao tentar deletar a publicação. - 316")
		return
	}

	log.Printf("Publicação %d removida com sucesso.", publicacaoID)
	respostas.JSON(w, http.StatusNoContent, nil)
}