package main

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func helloWorld(ctx context.Context, uri, username, password string) (string, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	greeting, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]any{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "failed to connect to database", err
	}

	return greeting.(string), nil
}

func parse(path string, info os.FileInfo, err error) error {
	if error != nil {
		fmt.Print("Error accessing filepath %s: %v\n", path, error)
		return nil
	}

	if info.IsDir() {
		return nil
	}

	fmt.Printf("Parsing file %s\n", path)
	
}

func main() {
	fmt.Println("Hello, World!")
	ctx := context.Background()
	var path String := os.Args[1]
	err := filepath.walk(path, )
	fmt.Println(helloWorld(ctx, "bolt://localhost:7687", "neo4j", "2ShutYoTrap"))
}
