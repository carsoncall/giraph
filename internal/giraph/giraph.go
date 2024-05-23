package giraph

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	model "github.com/carsoncall/giraph/pkg/model"
	"github.com/carsoncall/giraph/pkg/queue"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

type Giraph struct {
	Ctx          context.Context
	DB           DB
	Parser       sitter.Parser
	CodebasePath string
	CodebaseRoot string
}

func BirthGiraph(ctx context.Context, uri, username, password, codebasePath, codebaseRoot string) (*Giraph, error) {
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
		CodebasePath: codebasePath,
		CodebaseRoot: codebaseRoot,
	}, nil
}

func (giraph Giraph) Walk() error {
	error := filepath.Walk(giraph.CodebasePath, giraph.parse)
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

	// we have to store the names of things without the extension, because imports don't include extensions
	// by default. This is stupid, and eventually we can implement the same convoluted logic that Typescript 
	// uses to determine what file you mean when you have two files of the same name but different extensions
	// in the same directory (in TSX, it checks .tsx, then .ts, then .jsx, then .js)
	parentModelNode := model.Node{
		Name:     removeExtension(info.Name()),
		Contents: filepath.Join(giraph.CodebaseRoot,removeExtension(path)),
	}

	typeName = "string_fragment"
	childNode := bfs(parentNode, found)
	childNodeName := getContentsOfNode(childNode, file)
	
	// determine whether import is a local file or an external dependency
	childNodeContents := getFileNameIfExistsLocally(filepath.Dir(path), childNodeName)

	childModelNode := model.Node{
		Name:     childNodeName,
		Contents: childNodeContents,
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

func removeExtension(filename string) string {
	extension := filepath.Ext(filename)
	return strings.TrimSuffix(filename, extension) 
}

func getFileNameIfExistsLocally(dir, filename string) string {
	if fileExistsWithoutExtension(dir, filename) {
		return filepath.Join(dir, filename)
	} else {
		return filename
	}
}

func fileExistsWithoutExtension(dir string, filename string) bool {
	var exists bool
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasPrefix(info.Name(), filename) {
			exists = true
		}
		return nil
	})
	return exists
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
