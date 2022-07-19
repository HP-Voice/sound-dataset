package main

import (
	"database/sql"
	"io"
)

type Sample struct {
	Id        UUID   `json:"id"`
	LabelId   int    `json:"labelId"`
	LabelName string `json:"labelName"`
	Verdict   int    `json:"verdict"`
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
	row := db.QueryRow(`
		SELECT s.id, l.name AS label_name, s.verdict 
		FROM sample s LEFT JOIN label l ON l.id = s.label_id 
		WHERE s.verdict = 0
		ORDER BY s.id
		LIMIT 1
	`)
	if row == nil {
		return nil, sql.ErrNoRows
	}
	sample := &Sample{}
	err := row.Scan(&sample.Id, &sample.LabelName, &sample.Verdict)
	if err != nil {
		return nil, err
	}
	return sample, nil
}

func writeVerdict(sampleId UUID, verdict int) error {
	_, err := db.Exec("UPDATE sample SET verdict = $1 WHERE id = $2", verdict, sampleId)
	return err
}
