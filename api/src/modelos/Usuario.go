package modelos

import (
	"api/src/seguranca"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64 `json:"id,omitempty"`
	Nome     string `json:"Nome,omitempty"`
	Nick     string `json:"Nick,omitempty"`
	Email    string `json:"Email,omitempty"`
	Senha    string `json:"Senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		log.Printf("Ocorreu algum problema ao formatar os dados do usuário.")
		return erro
	}


	return nil
}

// middleware(schemas output) 
func (usuario *Usuario) validar( etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em banco.")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em banco.")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em banco.")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}


	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em banco.")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			log.Printf("Ocorreu um problema ao criar o Hash da senha.")
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}