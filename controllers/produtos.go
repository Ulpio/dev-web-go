package controllers

import (
	"Projeto/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAllProds()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome") // String 
		descricao := r.FormValue("descricao") //String
		preco :=  r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido,err := strconv.ParseFloat(preco, 64)
		if err != nil{
			panic(err.Error())
			log.Println("Erro na conversão do preço")
		}
		quantidadeConvertida,err := strconv.Atoi(quantidade)
		if err != nil{
			panic(err.Error())
			log.Println("Erro na conversão da quantidade")
		}

		models.CreateProd(nome,descricao,precoConvertido,quantidadeConvertida)
	}
	http.Redirect(w,r,"/",301)
}

func Delete(w http.ResponseWriter,r *http.Request){
	idProduto := r.URL.Query().Get("id")
	models.DeleteProd(idProduto)
	http.Redirect(w,r,"/",301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
			id := r.FormValue("id")
			nome := r.FormValue("nome")
			descricao := r.FormValue("descricao")
			preco := r.FormValue("preco")
			quantidade := r.FormValue("quantidade")

			idConvertido, err := strconv.Atoi(id)

			precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
			if err != nil {
					log.Println("Erro na convesão do preço para float64:", err)
			}

			quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
			if err != nil {
					log.Println("Erro na convesão da quantidade para int:", err)
			}

			models.AtualizaProduto(idConvertido, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}