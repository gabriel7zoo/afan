// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ratelimit_override_find_matches.sql

package db

import (
	"context"
)

const findRatelimitOverrideMatches = `-- name: FindRatelimitOverrideMatches :many
SELECT id, workspace_id, namespace_id, identifier, ` + "`" + `limit` + "`" + `, duration, async, sharding, created_at_m, updated_at_m, deleted_at_m FROM ratelimit_overrides
WHERE
    workspace_id = ?
    AND namespace_id = ?
    AND identifier LIKE ?
`

type FindRatelimitOverrideMatchesParams struct {
	WorkspaceID string `db:"workspace_id"`
	NamespaceID string `db:"namespace_id"`
	Identifier  string `db:"identifier"`
}

// FindRatelimitOverrideMatches
//
//	SELECT id, workspace_id, namespace_id, identifier, `limit`, duration, async, sharding, created_at_m, updated_at_m, deleted_at_m FROM ratelimit_overrides
//	WHERE
//	    workspace_id = ?
//	    AND namespace_id = ?
//	    AND identifier LIKE ?
func (q *Queries) FindRatelimitOverrideMatches(ctx context.Context, db DBTX, arg FindRatelimitOverrideMatchesParams) ([]RatelimitOverride, error) {
	rows, err := db.QueryContext(ctx, findRatelimitOverrideMatches, arg.WorkspaceID, arg.NamespaceID, arg.Identifier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RatelimitOverride
	for rows.Next() {
		var i RatelimitOverride
		if err := rows.Scan(
			&i.ID,
			&i.WorkspaceID,
			&i.NamespaceID,
			&i.Identifier,
			&i.Limit,
			&i.Duration,
			&i.Async,
			&i.Sharding,
			&i.CreatedAtM,
			&i.UpdatedAtM,
			&i.DeletedAtM,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
