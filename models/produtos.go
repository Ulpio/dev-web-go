package models

import (
	"Projeto/db"
	"fmt"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float32
	Quantidade int
}

func GetAllProds() []Produto {
	db := db.ConnectDB()
	selectAll, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectAll.Next() {
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
	return produtos
}

func CreateProd(nome,descricao string,preco float64,quantidade int){
	db := db.ConnectDB()

	insertData, err := db.Prepare("INSERT INTO produtos(nome,descricao,preco,quantidade) VALUES($1,$2,$3,$4)")
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("Inserido com sucesso")
	insertData.Exec(nome,descricao,preco,quantidade)

	defer db.Close()
}