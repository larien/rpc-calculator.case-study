package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Expression contém a expressão inserida
type Expression struct {
	InsertedExpression string
}

// Answer contém a resposta do cálculo
type Answer struct {
	Resposta float32 `json:"resultado"`
}

// Index é a função padrão
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bem vindo à calculadora RPC!")
}

// SubmitCalculation envia a expressão inserida para o client RPC
func SubmitCalculation(w http.ResponseWriter, r *http.Request) {
	exp := Expression{InsertedExpression: "TODO - Inserir expressão aqui"}

	// Envia expressão para o servidor
	resp, err := clientrpc.SendExpression(exp)
	if err != nil {
		log.Fatal("Erro")
	}

	resposta := Answer{Resposta: resp}

	// Retorna a resposta para o front end
	if err := json.NewEncoder(w).Encode(resposta); err != nil {
		panic(err)
	}
}
