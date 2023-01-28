package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaDeLogin vai renderizar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando tela de login")
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario vai carregar a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando tela de cadastro de usuário - 308")
		utils.ExecutarTemplate(w, "cadastro.html", nil)	
}

// CarregarPaginaPrincipal carrega a página principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando tela de home do usuário")
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		log.Printf("Erro ao carregar a pagina principal - 307")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		log.Printf("Erro no status code - 306")
		respostas.TratarStatusCodeErro(w, response)
		return
	}
	
	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		log.Printf("Erro no newdecoder - 307")
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	
	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID: usuarioID,
	})
}

// CarregarPaginaDeAtualizacaoDePublicacao carrega a página edição de publicação
func CarregarPaginaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando tela de edição de uma publicação")
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		log.Printf("Erro interno ao processar o json da requisição - 308")
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		log.Printf("Erro com status 400, verificar a requisição - 309")
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		log.Printf("Erro ao fazer o decoder do body - 310")
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	log.Printf("Carregando a pagina de atualizar uma publicação")
	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)
}

// CarregarPaginaDeUsuarios carrega a página com os usuários que atendem o filtro passado
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando pagina de pesquisa de usuário")
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOuNick)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		log.Printf("Ocorreu um erro ao fazer a autenticação - 311")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		log.Printf("Erro com status 400, verificar a requisição - 312")
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	var usuarios []modelos.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		log.Printf("Erro ao fazer o decoder - 313")
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	log.Printf("Retornando a lista de usuários")
	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario carrega a página do perfil do usuário
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		log.Printf("Erro ao fazer o parse do usuário")
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if usuarioID == usuarioLogadoID {
		http.Redirect(w, r, "/perfil", 302)
		return
	}

	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		log.Printf("Erro ao buscar os usuários")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario modelos.Usuario
		UsuarioLogadoID uint64
	}{
		Usuario: usuario,
		UsuarioLogadoID: usuarioLogadoID,
	})
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

// CarregarPaginaDeEdicaodoUsuario carrega a página para edição dos dados do usuário
func CarregarPaginaDeEdicaoDeUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	 
	canal := make(chan modelos.Usuario)
	go modelos.BuscarDadosDoUsuario(canal, usuarioID, r)
	usuario := <-canal

	if usuario.ID == 0 {
		log.Printf("Erro na função CarregarPaginaDeEdicaoDeUsuario")
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao buscar o usuário"})
		return
	}

	utils.ExecutarTemplate(w, "editar-usuario.html", usuario)
}

// CarregarPaginaDeAtualizacaoDeSenha carrega a página para atualização da senha do usuário
func CarregarPaginaDeAtualizacaoDeSenha(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando página atualizar-senha.html")
	utils.ExecutarTemplate(w, "atualizar-senha.html", nil)
}

// CarregarPays carrega a página payments
func CarregarPays(w http.ResponseWriter, r *http.Request) {
	log.Printf("Carregando página atualizar-senha.html")
	utils.ExecutarTemplate(w, "pays.html", nil)
}