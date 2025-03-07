package db

type Queries struct{}

// Query provides access to the generated database queries defined in the SQL files
//
// Example usage with a read operation:
//
//	import (
//	    "context"
//	    "github.com/unkeyed/unkey/go/pkg/db"
//	)
//
//	func GetKeyByHash(ctx context.Context, database db.Database, hash string) (key gen.Key, err error) {
//	    // Use Query singleton with the read replica
//	    return db.Query.FindKeyByHash(ctx, database.RO(), hash)
//	}
//
// Example usage with a write operation:
//
//	func InsertWorkspace(ctx context.Context, database db.Database, params gen.InsertWorkspaceParams) error {
//	    // Use Query singleton with the write replica
//	    return db.Query.InsertWorkspace(ctx, database.RW(), params)
//	}
//
// The Query object contains all the database operations defined in the SQL files
// and automatically generated by sqlc.
var Query Querier = &Queries{}
