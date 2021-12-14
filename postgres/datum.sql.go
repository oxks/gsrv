// Code generated by sqlc. DO NOT EDIT.
// source: datum.sql

package postgres

import (
	"context"
	"database/sql"
	"time"
)

const getDatums = `-- name: GetDatums :many
SELECT
	datum.datum, 
	datum.hash, 
	datum.previous_hash, 
	datum.created_at, 
	users.nickname
FROM
	datum
	INNER JOIN
	users
	ON 
		datum.author_id = users."id"
ORDER BY
	datum."id" DESC
`

type GetDatumsRow struct {
	Datum        string         `json:"datum"`
	Hash         sql.NullString `json:"hash"`
	PreviousHash sql.NullString `json:"previous_hash"`
	CreatedAt    time.Time      `json:"created_at"`
	Nickname     string         `json:"nickname"`
}

func (q *Queries) GetDatums(ctx context.Context) ([]GetDatumsRow, error) {
	rows, err := q.db.QueryContext(ctx, getDatums)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDatumsRow
	for rows.Next() {
		var i GetDatumsRow
		if err := rows.Scan(
			&i.Datum,
			&i.Hash,
			&i.PreviousHash,
			&i.CreatedAt,
			&i.Nickname,
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
