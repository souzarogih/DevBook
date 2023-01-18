package modelos

import "time"

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id"`
	Nome uint64 `json:"nome"`
	Email uint64 `json:"email"`
	Nick uint64 `json:"nick"`
	CriadoEm time.Time `json:"criadoEm"`
	Seguidores []Usuario `json:"seguidores"`
	Seguindo []Usuario `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}