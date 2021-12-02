package main

import (
	"fmt"
	"github.com/zchary-ma/grpc-server/server"
	"log"
)

func main() {
	port := "50051"
	s := server.NewServer()
	fmt.Println("Server started on port", port)
	log.Fatalln(s.ListenAndServe(port))
}
