package controllers

import (
	"api/src/structs"
	"encoding/json"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := structs.Usuario{"Jairo", 21, structs.Endereco{"rua 15", 1, "B", "Sao leo"}}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(usuario)

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := structs.Usuario{"Silas", 21, structs.Endereco{"rua 16", 111, "Bdd", "Sao leo dd"}}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(usuario)
}
