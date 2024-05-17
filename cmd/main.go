package main

import (
	"context"
	"fmt"
	"os"

	"github.com/carsoncall/giraph/internal/giraph"
)

func main() {
	fmt.Println("Hello, World!")
	ctx := context.Background()
	var path string = os.Args[1]

	giraph, err := giraph.BirthGiraph(ctx, "bolt://localhost:7687", "neo4j", "2ShutYoTrap", path)
	if err != nil {
		fmt.Print(err)
	}

	giraph.Walk(path)
}
