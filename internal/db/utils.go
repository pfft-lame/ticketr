package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToNullString(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}

	/*
		pgtype.Text{String: *s, Valid: s != nil}
		we can't do this directly because during runtime it could panic if s is infact nil and
		`String` tries to dereference the value first
	*/

	return pgtype.Text{String: *s, Valid: true}
}

func ToNullInt32(n *int) pgtype.Int4 {
	if n == nil {
		return pgtype.Int4{}
	}

	return pgtype.Int4{Int32: int32(*n), Valid: true}
}

func ToPgTimeTz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{}
	}

	return pgtype.Timestamptz{Time: *t, Valid: true}
}

func ToPgUUID(id *uuid.UUID) pgtype.UUID {
	if id == nil {
		return pgtype.UUID{Valid: false}
	}

	return pgtype.UUID{Valid: true, Bytes: *id}
}
