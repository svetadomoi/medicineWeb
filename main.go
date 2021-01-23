package main

import ("fmt"
	"net/http"
	"html/template")


type goods struct{
	Naming string
	Price float32

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



func main(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/info/",infoPage)
	http.ListenAndServe(":5000",nil)
}