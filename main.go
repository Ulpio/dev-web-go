package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB{
	connection := "user=postgres dbname=loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id 				 int
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
	db := connectDB()
	selectAll, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	p:= Produto{}
	produtos := []Produto{}

	for selectAll.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float32

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	temp.ExecuteTemplate(w, "Index", produtos)
}
