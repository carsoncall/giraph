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

	typeName := ""

	found := func(node *sitter.Node) bool {
		return node.Type() == typeName
	}

	root_node := tree.RootNode()

	typeName = "import_statement"
	parentNode := bfs(root_node, found)

	parentModelNode := model.Node{
		Name:     info.Name(),
		Contents: path,
	}

	typeName = "string_fragment"
	childNode := bfs(parentNode, found)
	childNodeContents := getContentsOfNode(childNode, file)

	childModelNode := model.Node{
		Name:     childNode.String(),
		Contents: filepath.Join(giraph.CodebaseRoot, string(childNodeContents)),
	}

	rel := model.Relationship{
		Name: "imports",
	}

	err = giraph.DB.PutRelationship(parentModelNode, childModelNode, rel)

	if err != nil {
		fmt.Printf("Error parsing import statement: %s", err.Error())
	}

	return nil
}

func buildModelNode(name, contents string) *model.Node {
	return &model.Node{
		Name:     name,
		Contents: contents,
	}
}

func getContentsOfNode(node *sitter.Node, file []byte) (contents string) {
	startByte := node.StartByte()
	endByte := node.EndByte()
	return string(file[startByte:endByte])
}

func bfs(rootNode *sitter.Node, found func(node *sitter.Node) bool) *sitter.Node {
	queue := queue.NewQueue()
	queue.Enqueue(rootNode)

	for !queue.IsEmpty() {
		var node *sitter.Node = queue.Dequeue().(*sitter.Node)

		if found(node) {
			return node
		}

		numChildren := node.ChildCount()
		for i := 0; i < int(numChildren); i++ {
			queue.Enqueue(node.Child(i))
		}
	}
	return nil
}
