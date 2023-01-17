package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarPublicacao chama a API para cadastrar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando tela de criar uma publicação")
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo": r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro no marshal das publicações - 223")
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		log.Printf("Ocorreu um erro ao enviar a requisição para o servidor - 224")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	// Este if foi colocado para resolver o problema do ajax que não consegue resolver resposta 200 com corpo vazio.
	if response.StatusCode == 200 {
		response.StatusCode = 204
		return
	}

	fmt.Println(">>", response.StatusCode)

	respostas.JSON(w, response.StatusCode, nil)
}

// CurtirPublicacao chama a API para curtir uma publicação
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro no parse do uint 305")
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.APIURL, publicacaoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro ao fazer a requisição - 304")
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		log.Printf("Erro no status code - 300")
		return
	}

	log.Printf("Requisição atendida com sucesso!")
	respostas.JSON(w, response.StatusCode, nil)
}

// DescurtirPublicacao chama a API para descurtir uma publicação
func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro no parse do uint - 301")
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/descurtir", config.APIURL, publicacaoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro ao enviar a requisição - 302")
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		log.Printf("Erro no status code - 303")
		return
	}

	log.Printf("Requisição atendida com sucesso!")
	respostas.JSON(w, response.StatusCode, nil)
}

// AtualizarPublicacao chama a API para atualizar uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro no parse do uint - 304")
		return
	}

	r.ParseForm()
	publicacao, erro := json.Marshal(map[string]string{
		"titulo": r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro no marshal das publicações - 305")
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		log.Printf("Algo aconteceu em AtualizarPublicacao - 306")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		log.Printf("Ocorreu um erro ao enviar a ao atualizar a requisição - 307")
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	log.Printf("Requisição atendida com sucesso!")
	respostas.JSON(w, response.StatusCode, nil)
}

// DeletarPublicacao chama a API para deletar uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro no parse do uint - 308")
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
	if erro != nil {
		log.Printf("Algo aconteceu em AtualizarPublicacao - 309")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		log.Printf("Ocorreu um erro ao enviar a ao atualizar a requisição - 310")
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	log.Printf("Requisição deletar atendida com sucesso!")
	respostas.JSON(w, response.StatusCode, nil)
}