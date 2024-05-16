package neo4j

import (
	"context"
	"fmt"
	"github.com/carsoncall/giraph/pkg/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4j struct {
	session neo4j.SessionWithContext
	ctx     context.Context
}

func (DB *Neo4j) Connect(ctx context.Context, uri, username, password string) error {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	if err != nil {
		fmt.Printf("failed to connect to Neo4j with this error: %s", err)
		return err
	}
	DB.session = session
	DB.ctx = ctx
	return nil
}

func (DB *Neo4j) PutRelationship(node1 Node, node2 Node, rel Relationship) error {
	query := fmt.Sprintf(`MERGE (a:Node{name: "%s"})
						  MERGE (b:Node{name: "%s"})
						  MERGE (a)-[r:"%s"]-(b)
						  RETURN a,b,r`, node1.name, node2.name, rel.name)
	_, err := DB.session.ExecuteWrite(DB.ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(DB.ctx, query, map[string]any{})
		if err != nil {
			return nil, err
		}

		if result.Next(DB.ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	return err
}
