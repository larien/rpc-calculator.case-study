package main

import (
    "fmt"
    "log"
    "net/rpc"
    "net/http"
    "github.com/gorilla/mux"
)

var Client *rpc.Client

// Server é a estrutura que contém os dados para conexão do servidor
type Server struct {
    ip string
    port string
}

// Port é a string que contém a porta para comunicação com o front
var Port string

func main() {
    log.Println("Iniciando main do client")

    // Inicializa struct com dados do servidor
    serv := new (Server)

    // Define IP e porta do servidor
    serv.ip = "127.0.0.1"
    serv.port = ":3334"

    log.Println("[MAIN-CLIENT] Conectando com servidor")
    
    // "Liga" para o servidor
    client, err := rpc.DialHTTP("tcp", serv.ip+serv.port)
    Client = client
   if err != nil {
        log.Fatal(fmt.Sprintf("[MAIN-CLIENT] Falha ao conectar com o servidor: %v", err))
        return
    }
    log.Println("[MAIN-CLIENT] Conexão com o servidor realizada com sucesso") 

    log.Println("[MAIN-CLIENT] Conetando com o interface da calculadora RPC") 
    // Conecta com o front
    router := mux.NewRouter()
	api := router.PathPrefix("/main/").Subrouter()
	setupFunctionsEndpoints(api)
    http.ListenAndServe(":8080", router)
    
    log.Println("[MAIN-CLIENT] Conexão com interface da calculadora RPC realizada com sucesso") 

    // fmt.Print("Insira a expressão: ")
    // scanner := bufio.NewScanner(os.Stdin)
    // scanner.Scan()
    // expressao := scanner.Text()
    // resposta, err := clientrpc.SendExpression(client, expressao)
    // if err != nil {
    //     log.Fatal("Falha na comunicação com o servidor RPC")
    // }
    //fmt.Printf("Resposta = %.2f",resposta)

}