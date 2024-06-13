package main

import (
	"Projeto/routes"
	"net/http"
)



func main() {
	routes.CarregarRotas() // Carrega as rotas
	http.ListenAndServe(":3000", nil) // Inicia o servidor na porta 3000
}
