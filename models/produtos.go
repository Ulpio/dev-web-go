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

func DeleteProd(id string){
	db := db.ConnectDB()

	deleteData, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil{
		panic(err.Error())
	}
	deleteData.Exec(id)

	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConnectDB()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
			panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
			var id, quantidade int
			var nome, descricao string
			var preco float32

			err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
			if err != nil {
					panic(err.Error())
			}
			produtoParaAtualizar.Id = id
			produtoParaAtualizar.Nome = nome
			produtoParaAtualizar.Descricao = descricao
			produtoParaAtualizar.Preco = preco
			produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
			panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}