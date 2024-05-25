package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/carsoncall/giraph/internal/giraph"
	"github.com/carsoncall/giraph/internal/server"
)

func main() {
	fmt.Println("Hello, World!")
	log.Println("Starting server...")
	server.StartServer()
	ctx := context.Background()
	var path string = os.Args[1]

	giraph, err := giraph.BirthGiraph(ctx, "bolt://localhost:7687", "neo4j", "2ShutYoTrap", path)
	if err != nil {
		fmt.Print(err)
	}

	giraph.Walk(path)
}
