package main

import (
//    "errors"
    "fmt"
    "net/http"
    "net/rpc"
    "calculadora/server/serverrpc"
)

func main() {

    ziguifryda := new(serverrpc.Ziguifryda)
    rpc.Register(ziguifryda)
    rpc.HandleHTTP()

    err := http.ListenAndServe(":3334", nil)
    if err != nil {
        fmt.Println(err.Error())
    }
}