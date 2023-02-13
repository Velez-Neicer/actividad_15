package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola, bienvenido a mi backend en Go!")
	})
	http.ListenAndServe(":8080", nil)
}
