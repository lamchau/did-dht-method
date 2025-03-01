// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package postgres

import (
	"context"
)

const listRecords = `-- name: ListRecords :many
SELECT key, value, sig, seq FROM pkarr_records
`

func (q *Queries) ListRecords(ctx context.Context) ([]PkarrRecord, error) {
	rows, err := q.db.Query(ctx, listRecords)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PkarrRecord
	for rows.Next() {
		var i PkarrRecord
		if err := rows.Scan(
			&i.Key,
			&i.Value,
			&i.Sig,
			&i.Seq,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readRecord = `-- name: ReadRecord :one
SELECT key, value, sig, seq FROM pkarr_records WHERE key = $1 LIMIT 1
`

func (q *Queries) ReadRecord(ctx context.Context, key string) (PkarrRecord, error) {
	row := q.db.QueryRow(ctx, readRecord, key)
	var i PkarrRecord
	err := row.Scan(
		&i.Key,
		&i.Value,
		&i.Sig,
		&i.Seq,
	)
	return i, err
}

const writeRecord = `-- name: WriteRecord :exec
INSERT INTO pkarr_records(key, value, sig, seq) VALUES($1, $2, $3, $4)
`

type WriteRecordParams struct {
	Key   string
	Value string
	Sig   string
	Seq   int64
}

func (q *Queries) WriteRecord(ctx context.Context, arg WriteRecordParams) error {
	_, err := q.db.Exec(ctx, writeRecord,
		arg.Key,
		arg.Value,
		arg.Sig,
		arg.Seq,
	)
	return err
}
