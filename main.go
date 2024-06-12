package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float32
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil) // Inicia o servidor na porta 3000
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Camiseta preta", 39.90, 10},
		{"Calça", "Calça jeans", 79.90, 5},
		{"Tênis", "Tênis esportivo", 129.90, 3},
		{"Cueca", "Cueca Insider", 150, 5},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
