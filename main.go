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
	Ctx          context.Context
	DbConn       neo4j.SessionWithContext
	Parser       sitter.Parser
	CodebaseRoot string
}

func BirthGiraph(ctx context.Context, uri, username, password, cbroot string) (*Giraph, error) {
	connection, err := Connect(ctx, uri, username, password)
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	if err != nil {
		fmt.Printf("failed to build Giraph: %s", err)
	}
	return &Giraph{
		Ctx:          ctx,
		DbConn:       connection,
		Parser:       *parser,
		CodebaseRoot: cbroot,
	}, nil
}

func Connect(ctx context.Context, uri, username, password string) (neo4j.SessionWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

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

	file, err := os.ReadFile(path)
	tree, err := giraph.Parser.ParseCtx(giraph.Ctx, nil, file)
	if err != nil {
		fmt.Printf("Failed to parse file: %s", err)
		return err
	}

	dbOp := func(node *sitter.Node) {
		fmt.Printf("Match: %s", info.Name())
		print := func(node *sitter.Node) {

			startByte := node.StartByte()
			endByte := node.EndByte()
			fileName := file[startByte:endByte]
			filePath := filepath.Join(giraph.CodebaseRoot, string(fileName))
			fmt.Printf("File: %s \n", filePath)

			query := fmt.Sprintf(` MERGE (a:Node{name: "%s"})
MERGE (b:Node{name: "%s"})
MERGE (a)-[r:Imports]-(b)
RETURN a,b,r`, path, filePath)

			// now an actual database operation
			result, err := giraph.DbConn.Run(giraph.Ctx, query, nil)
			if err != nil {
				fmt.Printf("Error writing to database: %s\n", err)
			}
			if result.Next(giraph.Ctx) {
				fmt.Printf("Wrote relationship to database\n")
			}
			result.Consume(giraph.Ctx)
		}
		bfs(node, "string_fragment", print)
	}
	root_node := tree.RootNode()
	bfs(root_node, "import_statement", dbOp)
	return nil
}

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

func bfs(rootNode *sitter.Node, nodeType string, dbOp func(*sitter.Node)) {
	queue := NewQueue()
	queue.Enqueue(rootNode)

	for !queue.IsEmpty() {
		var node *sitter.Node = queue.Dequeue().(*sitter.Node)

		if node.Type() == nodeType {
			dbOp(node)
		}

		numChildren := node.ChildCount()
		for i := 0; i < int(numChildren); i++ {
			queue.Enqueue(node.Child(i))
		}
	}
}

func main() {
	fmt.Println("Hello, World!")
	ctx := context.Background()
	var path string = os.Args[1]

	giraph, err := BirthGiraph(ctx, "bolt://localhost:7687", "neo4j", "2ShutYoTrap", path)
	if err != nil {
		fmt.Print(err)
	}

	parseFunc := func(path string, info os.FileInfo, err error) error {
		return parse(*giraph, path, info, err)
	}
	filepath.Walk(path, parseFunc)
}
