package giraph

import (
	"context"
	"github.com/carsoncall/pkg/neo4j"
)

// DB defines the behavior expected from a database.
type DB interface {
	Connect(ctx context.Context, uri, username, password string) error
	PutRelationship(node1 Node, node2 Node, rel Relationship) error
}

// NewDatabase creates and returns an instance of a database implementation.
// This is where the current implementation is changed.
func NewDatabase() DB {
	return &Neo4j{}
}
