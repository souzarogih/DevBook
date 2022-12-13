package modelos

import (
	"errors"
	"strings"
	"time"
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
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (Usuario *Usuario) validar() error {
	if Usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em banco.")
	}

	if Usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em banco.")
	}

	if Usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em banco.")
	}

	if Usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em banco.")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}