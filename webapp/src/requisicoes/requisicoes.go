package requisicoes

import (
	"io"
	"log"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutenticacao é utilizada para colocar o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		log.Printf("Ocorreu um erro ao tentar fazer a requisição - 122")
		return nil, erro
	}

	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		log.Printf("Ocorreu um erro quando o cliente enviou a requisição - 123")
		return nil, erro
	}

	log.Printf("Enviando requisição para a api")
	return response, nil
}