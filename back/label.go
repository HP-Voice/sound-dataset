package main

import (
	"context"
	"math/rand"
)

type Label struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Amount         int    `json:"amount"`
	AmountApproved int    `json:"amountApproved"`
}

func readLabels() ([]*Label, error) {
	ctx := context.Background()
	rows, err := db.Query(ctx, `SELECT id, name FROM label`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*Label, 0)
	for rows.Next() {
		label := &Label{}
		err = rows.Scan(&label.Id, &label.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, label)
	}

	return result, nil
}

func getRandomSpellId() (int, error) {
	ctx := context.Background()
	rows, err := db.Query(ctx, `SELECT id, amount FROM label_state WHERE id > 1`)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	spells := make([]*Label, 0)
	for rows.Next() {
		spell := &Label{}
		err = rows.Scan(&spell.Id, &spell.Amount)
		if err != nil {
			return 0, err
		}
		spells = append(spells, spell)
	}

	return spells[rand.Intn(len(spells))].Id, nil
}
