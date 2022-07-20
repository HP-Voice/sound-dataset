package main

import (
	"context"
	"database/sql"
	"io"
	"os"
)

type Sample struct {
	Id        UUID   `json:"id"`
	LabelId   int    `json:"labelId"`
	LabelName string `json:"labelName"`
	Verdict   int    `json:"verdict"`
}

func writeSample(labelId int, data io.Reader) (_ *Sample, err error) {
	ctx := context.Background()
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err == nil {
			_ = tx.Commit(ctx)
		} else {
			_ = tx.Rollback(ctx)
		}
	}()

	rows, err := tx.Query(ctx, `INSERT INTO sample (label_id) VALUES ($1) RETURNING id`, labelId)
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
	err = rows.Scan(&sample.Id)
	if err != nil {
		return nil, err
	}

	err = writeFile(filenameOf(sample.Id), data)
	if err != nil {
		return nil, err
	}

	return sample, nil
}

func readSampleForApproval() (*Sample, error) {
	ctx := context.Background()
	row := db.QueryRow(ctx, `
		SELECT s.id, l.name AS label_name, s.verdict 
		FROM sample s LEFT JOIN label l ON l.id = s.label_id 
		WHERE s.verdict = 0
		ORDER BY s.id
		LIMIT 1
	`)
	sample := &Sample{}
	err := row.Scan(&sample.Id, &sample.LabelName, &sample.Verdict)
	if err != nil {
		return nil, err
	}
	return sample, nil
}

func writeVerdict(sampleId UUID, verdict int) error {
	ctx := context.Background()
	_, err := db.Exec(ctx, "UPDATE sample SET verdict = $1 WHERE id = $2", verdict, sampleId.Bytes)
	return err
}

func cleanSamples() (err error) {
	ctx := context.Background()
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			_ = tx.Commit(ctx)
		} else {
			_ = tx.Rollback(ctx)
		}
	}()

	rows, err := tx.Query(ctx, `SELECT id FROM sample`)
	if err != nil {
		return err
	}
	defer rows.Close()

	clean := make([]UUID, 0)

	for rows.Next() {
		uuid := UUID{}
		err = rows.Scan(&uuid)
		if err != nil {
			panic(err)
		}
		_, err = os.Stat(filenameOf(uuid))
		if err != nil {
			if os.IsNotExist(err) {
				clean = append(clean, uuid)
				continue
			}
			return err
		}
	}

	for _, c := range clean {
		_, err = tx.Exec(ctx, `DELETE FROM sample WHERE id = $1`, c.Bytes)
		if err != nil {
			return err
		}
	}

	return nil
}
