package main

import (
    "github.com/zchary-ma/pre/server"
    "log"
)

func main()  {
    port := "50051"
    log.Fatalln(server.ListenAndServe(port))
}