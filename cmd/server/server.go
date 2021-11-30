package main

import (
	"github.com/zchary-ma/pre/server"
	"log"
)

func main() {
	port := "50051"
	s := server.NewServer()
	log.Fatalln(s.ListenAndServe(port))
}
