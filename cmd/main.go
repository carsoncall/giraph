package main

import (
	"context"
	"fmt"
	"os"

	"github.com/carsoncall/giraph/internal/giraph"
)

// func GetText(startByte, endByte int64, file os.File) {
// 	// Seek to the start position
// 	_, err := file.Seek(startByte, io.SeekStart)
// 	if err != nil {
// 		fmt.Println("Error seeking file:", err)
// 		return
// 	}

// 	// Read the specific slice of bytes
// 	buf := make([]byte, length)
// 	_, err = file.Read(buf)
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// }

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
