// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: datum.sql

package postgres

import (
	"context"
	"database/sql"
	"time"
)

const allDatum = `-- name: AllDatum :many
SELECT
	datum.datum, 
	datum.previous_hash, 
	datum.hash, 
	datum.created_at, 
	users.firstname, 
	users.lastname
FROM
	datum
	INNER JOIN
	users
	ON 
		datum.author_id = users."id"
`

type AllDatumRow struct {
	Datum        string         `json:"datum"`
	PreviousHash sql.NullString `json:"previous_hash"`
	Hash         sql.NullString `json:"hash"`
	CreatedAt    time.Time      `json:"created_at"`
	Firstname    string         `json:"firstname"`
	Lastname     string         `json:"lastname"`
}

func (q *Queries) AllDatum(ctx context.Context) ([]AllDatumRow, error) {
	rows, err := q.db.QueryContext(ctx, allDatum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllDatumRow
	for rows.Next() {
		var i AllDatumRow
		if err := rows.Scan(
			&i.Datum,
			&i.PreviousHash,
			&i.Hash,
			&i.CreatedAt,
			&i.Firstname,
			&i.Lastname,
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
