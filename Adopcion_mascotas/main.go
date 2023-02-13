package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := `
		<html>
			<head>
				<title>Adopta a una mascota</title>
			</head>
			<body>
				<h1>Adopta a una mascota</h1>
				<form action="/adopt" method="post">
					<label for="pet-type">Escoja el tipo de mascota:</label>
					<select id="pet-type" name="pet-type">
						<option value="Perro">Perro</option>
						<option value="Gato">Gato</option>
						<option value="Aves">Aves</option>
					</select>
					<br><br>
					<label for="name">Escriba su nombre nombre:</label>
					<input type="text" id="name" name="name">
					<br><br>
					<input type="submit" value="Adoptar">
				</form>
			</body>
		</html>
		`
		fmt.Fprint(w, page)
	})
	http.HandleFunc("/adopt", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			petType := r.FormValue("pet-type")
			name := r.FormValue("name")
			fmt.Fprintf(w, "Gracias, %s, por adoptar un %s!", name, petType)
		}
	})
	http.ListenAndServe(":8080", nil)
}
