package db

import "github.com/jackc/pgx/v5/pgtype"

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
