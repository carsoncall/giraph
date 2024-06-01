package server

import (
	"log"
	"net/http"

	"github.com/carsoncall/giraph/internal/server/protobuf"
	"google.golang.org/protobuf/proto"
)

func StartServer() {
	http.HandleFunc("/graph", graphHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	var req protobuf.GraphRequest
	err := proto.Unmarshal([]byte(r.FormValue("request")), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Replace with actual logic to retrieve nodes and edges
	nodes, edges := getDummyData()

	resp := &protobuf.GraphResponse{
		Nodes: nodes,
		Edges: edges,
	}

	respBytes, err := proto.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(respBytes)
}

func getDummyData() ([]*protobuf.Node, []*protobuf.Edge) {
	// Dummy data for testing
	nodes := []*protobuf.Node{
		{Filepath: "/path/to/file1", StartByte: 0, EndByte: 100, Contents: "File 1 contents", Hash: "hash1"},
		{Filepath: "/path/to/file2", StartByte: 0, EndByte: 200, Contents: "File 2 contents", Hash: "hash2"},
	}

	edges := []*protobuf.Edge{
		{Name: "Edge 1", ParentHash: "hash1", ChildHash: "hash2"},
	}

	return nodes, edges
}
