package main

import ("fmt"
	"net/http"
	"html/template")


type goods struct{
	Naming string
	Price float32

}

func VitaminsCategory(w http.ResponseWriter, r *http.Request) {

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
	medcine := goods{"Larista", 202.22}
	tmpl, _ :=template.ParseFiles("templates/index.html")
	tmpl.Execute(w,medcine)
}

func infoPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "улица Пушкина дом Колотушкина")

}

func initSite(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/info/",infoPage)
	http.HandleFunc("/category/vitamins/",VitaminsCategory)
	http.ListenAndServe(":5000",nil)
}

func main(){
	initSite()
}