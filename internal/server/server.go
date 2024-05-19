// internal/server/server.go

package server

import (
	"context"
	"net/http"

	"github.com/carsoncall/giraph/internal/server/protobuf"
	"github.com/carsoncall/giraph/pkg/neo4j"
	"github.com/golang/protobuf/proto"
)

func StartServer() {
	http.HandleFunc("/graph", GraphHandler)
	http.ListenAndServe(":8080", nil)
}

func GraphHandler(w http.ResponseWriter, r *http.Request) {
	var req protobuf.GraphRequest
	if err := proto.Unmarshal(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := &neo4j.Neo4j{}
	ctx := context.Background()
	if err := db.Connect(ctx, "bolt://localhost:7687", "neo4j", "password"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var nodes []*protobuf.Node
	var edges []*protobuf.Edge
	for _, name := range req.NodeNames {
		node := &protobuf.Node{Name: name}
		edges, err := fetchEdges(db, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nodes = append(nodes, node)
		edges = append(edges, edges...)
	}

	resp := &protobuf.GraphResponse{
		Nodes: nodes,
		Edges: edges,
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(data)
}

func fetchEdges(db *neo4j.Neo4j, nodeName string) ([]*protobuf.Edge, error) {
	panic("fetchEdges is not implemented yet")
}
