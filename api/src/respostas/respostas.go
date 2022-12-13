package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para a requisição (schemas output)
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
			log.Print("Erro na função JSON")
		}
	}
	
}

// Erro retorna um erro em formato JSON (schemas output de erro)
func Erro(w http.ResponseWriter, statusCode int, erro error, codigo string) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
		Codigo string `json:"codigo"`
	}{
		Erro: erro.Error(),
		Codigo: codigo,
	})

}