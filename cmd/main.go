package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/carsoncall/giraph/internal/giraph"
)

func main() {
	fmt.Println("Hello, World!")
	ctx := context.Background()
	var path string = os.Args[1]
	rootDir := filepath.Base(path)
	fmt.Printf("Parsing codebase with root directory %s", rootDir)

	giraph, err := giraph.BirthGiraph(ctx, "bolt://localhost:7687", "neo4j", "2ShutYoTrap", path, rootDir)
	if err != nil {
		fmt.Print(err)
	}

	giraph.Walk()
}
