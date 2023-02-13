package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Usuario struct {
	Nombres     string
	Nombres2    string
	Habilidades string
	Edad        int
	Edad2       int
}

func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	user := Usuario{"Maria Herrera", "Neicer Velez", "estudiantes", 21, 20}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}

}

func main() {

	//routes
	http.HandleFunc("/", Index)
	http.HandleFunc("/contact", contactHandler)

	//start the server
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func contactHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Pagina de contacto"))
}
