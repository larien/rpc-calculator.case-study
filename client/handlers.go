package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Expression struct {
	InsertedExpression string
}

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

	// Retorna a resposta para o front end
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}
