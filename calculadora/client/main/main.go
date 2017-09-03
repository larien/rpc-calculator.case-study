package main

import (
    "fmt"
    "log"
    "net/rpc"
   // "os"
)

type Args struct {
    Expressao string
}

type Quotient struct {
    Quo, Rem int
}


func main() {
    // if len(os.Args) != 2 {
    //     fmt.Println("Usage: ", os.Args[0], "server")
    //     os.Exit(1)
    // }
    serverAddress := "127.0.0.1"

    client, err := rpc.DialHTTP("tcp", serverAddress+":3334")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    // Synchronous call

    args := Args{"Testando"}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: resposta=%d\n", reply)

}