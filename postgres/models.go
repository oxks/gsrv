// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"database/sql"
	"time"
)

type Datum struct {
	ID           int64          `json:"id"`
	AuthorID     int64          `json:"author_id"`
	Datum        string         `json:"datum"`
	PreviousHash sql.NullString `json:"previous_hash"`
	Hash         sql.NullString `json:"hash"`
	CreatedAt    time.Time      `json:"created_at"`
}

type User struct {
	ID        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Deleted   bool      `json:"deleted"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}