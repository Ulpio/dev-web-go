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