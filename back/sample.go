package main

import (
	"database/sql"
	"github.com/jackc/pgx/pgtype"
	"io"
)

type Sample struct {
	Id       pgtype.UUID
	LabelId  int
	Approved bool
}

func writeSample(labelId int, data io.Reader) (_ *Sample, err error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	rows, err := tx.Query(`INSERT INTO sample (label_id) VALUES ($1) RETURNING id`, labelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	sample := &Sample{
		LabelId: labelId,
	}
	_ = rows.Scan(&sample.Id)

	err = writeFile(filenameOf(sample.Id), data)
	if err != nil {
		return nil, err
	}

	return sample, nil
}
