package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Iniciando main - conectando com front")

	router := mux.NewRouter()

	api := router.PathPrefix("/main/").Subrouter()
	setupTestEndpoints(api)

	//fs := http.FileServer(http.Dir())
	http.ListenAndServe(":8080", router)

}

func setupTestEndpoints(api *mux.Router) {
	log.Println("Chegou aqui1")
	api.HandleFunc("/calculadora", GetOnly(Calculate)).Methods("GET")
}

// serveResult converte um objeto para formato JSON e o envia para o cliente
func serveResult(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	log.Println("Chegou aqui2")

	var res struct {
		Result float64
	}

	res.Result = 50.4

	serveResult(w, res) // res - struct
}

// serverErrorEnvelope é um envelope dos erros retornados
// pelo servidor
type serverErrorEnvelope struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// serveInternalError envia um erro interno para o cliente
func serveInternalError(w http.ResponseWriter, msg string, args ...interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	env := serverErrorEnvelope{
		Error:   "internal",
		Message: fmt.Sprintf(msg, args...),
	}
	json.NewEncoder(w).Encode(env)
}

type handler func(w http.ResponseWriter, r *http.Request)

func GetOnly(h handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "get only", http.StatusMethodNotAllowed)
	}
}

func PostOnly(h handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}
