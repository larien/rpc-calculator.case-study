package main

import(
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"
	"calculadora/client/clientrpc"
)

func setupFunctionsEndpoints(api *mux.Router) {
	log.Println("Chegou aqui1")
	api.HandleFunc("/calculadora", GetOnly(Calculate)).Methods("GET")
	api.HandleFunc("/funcao", GetOnly(Functions)).Methods("GET")
}

// Calculate é a função que recebe a expressão básica do usuário e se comunica com o
// servidor via RPC para fazer o cálculo
func Calculate(w http.ResponseWriter, r *http.Request) {

	log.Println(fmt.Sprintf("R = %v",r))

	// Recebe expressão
	vars := mux.Vars(r)
	exp := vars["Expressao"]

	log.Println(fmt.Sprintf("[MAIN-CLIENT] Expressão foi recebida pelo backend. Enviando para cálculo no servidor: %s",exp))

	resposta, err := clientrpc.SendExpression(Client, exp)
    if err != nil {
        log.Fatal("Falha na comunicação com o servidor RPC")
    }

	var res struct {
		Result float64
	}

	res.Result = resposta

	log.Println("[MAIN-CLIENT] Resposta calculada: ",res.Result)

	serveResult(w, res)
}

func Functions (w http.ResponseWriter, r *http.Request) {

	// Recebe expressão
	vars := mux.Vars(r)
	funcao := vars["Funcao"]
	valor, _ := strconv.ParseFloat(vars["Valor"], 64)
	
	var res struct {
		Result float64
	}

	resposta, err := clientrpc.SendFunction(Client, funcao, valor)
    if err != nil {
        log.Fatal("Falha na comunicação com o servidor RPC")
    }

	res.Result = resposta

	log.Println("[MAIN-CLIENT] Resposta calculada: ", res.Result)

	serveResult(w, res)
}


// serveResult converte um objeto para formato JSON e o envia para o cliente
func serveResult(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
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

// GetOnly é a função que responde o método GET
func GetOnly(h handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "get only", http.StatusMethodNotAllowed)
	}
}

// PostOnly é a função que responde o método POST
func PostOnly(h handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}
