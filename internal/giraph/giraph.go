package giraph

import (
	"context"
	"fmt"
	
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript":w
	
)


type Giraph struct {
	Ctx          context.Context
	DbConn       neo4j.SessionWithContext
	Parser       sitter.Parser
	CodebaseRoot string
}

func BirthGiraph(ctx context.Context, uri, username, password, cbroot string) (*Giraph, error) {
	connection, err := db.Connect(ctx, uri, username, password)
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

func (giraph Giraph) Walk(path string) error {
	filepath.Walk(path, parse)
}

// This function is of type WalkFunc (filepath.WalkFunc), as is necessary to be the parameter
// to filepath.Walk.
func parse(path string, info os.FileInfo, err error) error {
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

func bfs(rootNode *sitter.Node, nodeType string, dbOp func(*sitter.Node)) {
	queue := queue.NewQueue()
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

func bfs(rootNode *sitter.Node, nodeType string, dbOp func(*sitter.Node)) {
	queue := queue.NewQueue()
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

