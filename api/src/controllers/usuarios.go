package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro, "106")
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "105")
		return
	}
	
	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "104")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "103")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "102")
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "101")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "100")
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um usuário salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	log.Print("Função BuscarUsuario do pacote controllers em execução")
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "107")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "108")
		log.Print("Erro ao conectar com o banco de dados - 108")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "109")
		log.Print("Erro ao localizar o usuário no banco de dados - 109")
		return
	}
	
	respostas.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario altera as informações de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	log.Print("Função AtualizarUsuario em execução no pacote controllers")

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "109")
		log.Print("Erro ao fazer o parse do campo usuarioId, 109")
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "118")
		log.Printf("Erro ao extrair o usuário do token - 118")
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu."), "122")
		log.Printf("Usuário do token é diferente do usuário da requisição - 122")
		return
	}

	log.Printf("Atualizando usuário %d usando o token do usuário %d ", usuarioID, usuarioIDNoToken)

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro, "110")
		log.Print("Erro ao fazer processar o campo da requisição, 110")
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "111")
		log.Print("Erro ao processar o body da requisição, 111")
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "112")
		log.Print("Erro ao prepara a requisição para salvar em banco, 112")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "113")
		log.Print("Erro ao conectar com o banco de dados - 113")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "114")
		log.Print("Ocorreu um problema ao atualizar os dados do usuário no banco de dados - 114")
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario exclui as informações de um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "115")
		log.Print("Erro ao converter o usuarioId - 115")
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "123")
		log.Print("Erro ao extrair o usuário do token o usuarioId - 123")
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o seu."), "124")
		log.Print("Não é possível deletar um usuário que com id diferente na requisição e no token. - 124")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "116")
		log.Print("Erro ao conectar com o banco de dados - 116")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "117")
		log.Print("Erro ao criar um novo repositório - 117")
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// SeguirUsuario permite que um usuário siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "118")
		log.Printf("Erro ao estrair o seguidor do token - 118")
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "119")
		log.Printf("Erro ao converter o usuário di da requisição - 119")
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"), "120")
		log.Printf("Usuário tentando seguir a si mesmo - 120")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "121")
		log.Print("Erro ao conectar com o banco de dados - 121")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "122")
		log.Printf("Erro acessar salvar as informações no banco de dados - 122")
		return
	}

	log.Printf("Associação do usuário %d seguindo o usuário %d realizada com sucesso ",seguidorID,  usuarioID)
	respostas.JSON(w, http.StatusNoContent, nil)

}

// PararDeSeguirUsuario permite que um usuário pare de seguir outro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "123")
		log.Printf("Erro ao extrair o usuário que irá deixar de seguir - 123")
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "124")
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível parar de seguir você mesmo."), "125")
		log.Printf("Não é possível deixar de seguir a si mesmo - 125")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "121")
		log.Print("Erro ao conectar com o banco de dados - 121")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "122")
		log.Printf("Não foi possível se acessar a tabela para deixar de seguir um usuário - 122")
		return
	}

	log.Printf("Usuário %d deixou de seguir o usuário %d com sucesso.", seguidorID, usuarioID)
	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "123")
		log.Printf("Erro ao converter a identificação do usuario - 123")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "124")
		log.Printf("Erro ao converter a identificação do usuario - 123")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "124")
		return
	}

	log.Printf("Seguidores listados com sucesso.")
	respostas.JSON(w, http.StatusOK, seguidores)

}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "123")
		log.Printf("Erro ao converter a identificação do usuario - 123")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "124")
		log.Printf("Erro ao converter a identificação do usuario - 124")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "125")
		log.Printf("Erro interno ao localizar quem o usuário está seguindo - 125")
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)
}

// AtualizarSenha permite alterar a senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro, "126")
		log.Printf("Não foi possível extrar o usuárioID do token - 126")
		return
	}
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "127")
		log.Printf("Erro converter o usuario id da requisição - 127")
		return
	}

	if usuarioIDNoToken != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha do um usuário que não é o seu."), "128")
		log.Printf("Usuário está tentando atualizar a senha com o id diferente do seu. - 128")
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	var senha modelos.Senha
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "129")
		log.Printf("Erro ao converter o struc/json - 129")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "124")
		log.Printf("Erro ao converter a identificação do usuario - 124")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "125")
		log.Printf("Erro interno ao localizar a senha do banco - 125")
		return
	}

	if erro = seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("A senha atual não é a mesma salva no banco"), "126")
		log.Printf("Erro interno ao localizar a senha do banco - 126")
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro, "127")
		log.Printf("Não foi possível gerar o hash da senha nova - 127")
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro, "128")
		log.Printf("Erro interno ao atualizar a senha do usuário - 128")
		return
	}

	log.Printf("A nova senha do usuário foi atualizada com sucesso.")
	respostas.JSON(w, http.StatusNoContent, nil)
}