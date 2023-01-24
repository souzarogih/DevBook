package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CriarUsuario chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick": r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		log.Printf("Ocorreu um erro ao tentar capturar os campos do formulário - 100")
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		log.Printf("Ocorreu um erro ao enviar a requisição para o servidor - 101")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

	//fmt.Println("resposta> ", response.Body)
	//fmt.Println(bytes.NewBuffer(usuario))
	// outra forma de pegar os dados do banco
	//nome := r.FormValue("nome")
	//fmt.Println(nome)
}

// PararDeSeguirUsuario chama a API para parar de seguir um usuário
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// SeguirUsuario chama a API para seguir um usuário
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// CarregarPerfilDoUsuarioLogado carregar a página do perfil do usuário logado
func CarregarPerfilDoUsuarioLogado(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)
}

