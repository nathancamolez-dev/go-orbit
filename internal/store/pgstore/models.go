// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package pgstore

import (
	"time"

	"github.com/google/uuid"
)

type Goal struct {
	ID                     uuid.UUID `json:"id"`
	Title                  string    `json:"title"`
	Desiredweeklyfrequency int32     `json:"desiredweeklyfrequency"`
	Createdat              time.Time `json:"createdat"`
}

type Goalscompletion struct {
	ID        uuid.UUID `json:"id"`
	Goalid    uuid.UUID `json:"goalid"`
	Createdat time.Time `json:"createdat"`
}
