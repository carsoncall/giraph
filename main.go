package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

type Giraph struct {
	Ctx    context.Context
	DbConn neo4j.SessionWithContext
	Parser sitter.Parser
}

func BirthGiraph(ctx context.Context, uri, username, password string) (*Giraph, error) {
	connection, err := Connect(ctx, uri, username, password)
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	if err != nil {
		fmt.Printf("failed to build Giraph: %s", err)
	}
	return &Giraph{
		Ctx:    ctx,
		DbConn: connection,
		Parser: *parser,
	}, nil
}

func Connect(ctx context.Context, uri, username, password string) (neo4j.SessionWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// greeting, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
	// 	result, err := transaction.Run(ctx,
	// 		"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
	// 		map[string]any{"message": "hello, world"})
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	if result.Next(ctx) {
	// 		return result.Record().Values[0], nil
	// 	}

	// 	return nil, result.Err()
	//})
	if err != nil {
		fmt.Printf("failed to connect to Neo4j with this error: %s", err)
		return nil, err
	}
	return session, nil
}

func parse(giraph Giraph, path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing filepath %s: %v\n", path, err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	fmt.Printf("Parsing file %s\n", path)
	file, err := os.ReadFile(path)
	tree, err := giraph.Parser.ParseCtx(giraph.Ctx, nil, file)
	if err != nil {
		fmt.Printf("Failed to parse file: %s", err)
		return err
	}

	return nil
}

func main() {
	fmt.Println("Hello, World!")
	ctx := context.Background()
	var path string = os.Args[1]

	giraph, err := BirthGiraph(ctx, "bolt://localhost:7687", "neo4j", "2ShutYoTrap")
	if err != nil {
		fmt.Print(err)
	}

	parseFunc := func(path string, info os.FileInfo, err error) error {
		return parse(*giraph, path, info, err)
	}
	filepath.Walk(path, parseFunc)
}
