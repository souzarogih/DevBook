package utils

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates inserre os templates html na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTemplate renderiza uma página html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	log.Printf("Chamando função ExecutarTemplate de utils")
	templates.ExecuteTemplate(w, template, dados)
}