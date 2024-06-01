package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/carsoncall/giraph/internal/server/protobuf"
	"google.golang.org/protobuf/proto"
)

func TestGraphEndpoint(t *testing.T) {
	go StartServer()

	req := &protobuf.GraphRequest{
		Request:       "Test Request",
		ProjectRoot:   "/path/to/project",
		NumSteps:      proto.Int32(5),
		StartNodeHash: proto.String("start_hash"),
	}

	reqBytes, err := proto.Marshal(req)
	if err != nil {
		t.Fatalf("Error marshaling request: %v", err)
	}

	resp, err := http.PostForm("http://localhost:8080/graph", url.Values{"request": {string(reqBytes)}})
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response: %v", err)
	}

	var graphResp protobuf.GraphResponse
	err = proto.Unmarshal(respBytes, &graphResp)
	if err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}

	fmt.Println("Nodes:")
	for _, node := range graphResp.Nodes {
		fmt.Printf("  Filepath: %s, StartByte: %d, EndByte: %d, Contents: %s, Hash: %s\n",
			node.Filepath, node.StartByte, node.EndByte, node.Contents, node.Hash)
	}

	fmt.Println("Edges:")
	for _, edge := range graphResp.Edges {
		fmt.Printf("  Name: %s, ParentHash: %s, ChildHash: %s\n",
			edge.Name, edge.ParentHash, edge.ChildHash)
	}
}
