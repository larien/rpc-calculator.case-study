package main

import (
    "fmt"
    "log"
    "net/rpc"
    "clientrpc"
    "bufio"
    "os"
)

type Args struct {
    Expression string
}

type Quotient struct {
    Quo, Rem int
}

type Server struct {
    ip string
    port string
}

func main() {
    // Inicializa struct com dados do servidor
    serv := new (Server)

    // Define IP e porta do servidor
    serv.ip = "127.0.0.1"
    serv.port = ":3334"

    // "Liga" para o servidor
    client, err := rpc.DialHTTP("tcp", serv.ip+serv.port)
   if err != nil {
        log.Fatal("Discando:", err)
    }
    
    fmt.Print("Insira a expressão: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    expressao := scanner.Text()
    resposta, err := clientrpc.SendExpression(client, expressao)
    if err != nil {
        log.Fatal("Falha na comunicação com o servidor RPC")
    }
    fmt.Printf("Resposta = %.2f",resposta)

}