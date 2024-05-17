package giraph

import (
	"context"
	model "github.com/carsoncall/giraph/pkg/model"
	n4j "github.com/carsoncall/giraph/pkg/neo4j"
)

// DB defines the behavior expected from a database.
type DB interface {
	Connect(ctx context.Context, uri, username, password string) error
	PutRelationship(node1 model.Node, node2 model.Node, rel model.Relationship) error
}

// NewDatabase creates and returns an instance of a database implementation.
// This is where the current implementation is changed.
func NewDatabase() DB {
	return &n4j.Neo4j{}
}
