package modelos

import "time"

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64 `json:"id,omitempty"`
	Nome     string `json:"Nome,omitempty"`
	Nick     string `json:"Nick,omitempty"`
	Email    string `json:"Email,omitempty"`
	Senha    string `json:"Senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}