package neo4j

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Connect(ctx context.Context, uri, username, password string) (neo4j.SessionWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	// greeting, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
	// 	result, err := transaction.Run(ctx,
	// 		"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
	// 		map[string]any{"message": "hello, world"})
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	if result.Next(ctx) {
	// 		return result.Record().Values[0], nil
	// 	}

	// 	return nil, result.Err()
	//})
	if err != nil {
		fmt.Printf("failed to connect to Neo4j with this error: %s", err)
		return nil, err
	}
	return session, nil
}
