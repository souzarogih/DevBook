package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacoId}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicao,
		RequerAutenticacao: true,
	},
}