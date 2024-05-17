package giraph

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	model "github.com/carsoncall/giraph/pkg/model"
	"github.com/carsoncall/giraph/pkg/queue"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

type Giraph struct {
	Ctx          context.Context
	DB           DB
	Parser       sitter.Parser
	CodebaseRoot string
}

func BirthGiraph(ctx context.Context, uri, username, password, cbroot string) (*Giraph, error) {
	database := NewDatabase()
	err := database.Connect(ctx, uri, username, password)
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	if err != nil {
		fmt.Printf("failed to build Giraph: %s", err)
	}
	return &Giraph{
		Ctx:          ctx,
		DB:           database,
		Parser:       *parser,
		CodebaseRoot: cbroot,
	}, nil
}

func (giraph Giraph) Walk(path string) error {
	error := filepath.Walk(path, giraph.parse)
	return error
}

// This function is of type WalkFunc (filepath.WalkFunc), as is necessary to be the parameter
// to filepath.Walk.
func (giraph Giraph) parse(path string, info os.FileInfo, err error) error {
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
			importer := model.Node{
				Name:     path,
				Contents: path, // for now??
			}

			importee := model.Node{
				Name:     filePath,
				Contents: filePath,
			}

			relationship := model.Relationship{
				Name: "imports",
			}

			// now an actual database operation
			err := giraph.DB.PutRelationship(importer, importee, relationship)

			if err != nil {
				fmt.Printf("Error writing to database: %s\n", err)
			}
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
