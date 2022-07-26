package main

import "context"

type Stats struct {
	Labels              []*Label `json:"labels"`
	TotalAmount         int      `json:"totalAmount"`
	TotalAmountApproved int      `json:"totalAmountApproved"`
}

func readStats() (*Stats, error) {
	ctx := context.Background()
	rows, err := db.Query(ctx, "SELECT id, name, amount, amount_approved FROM label_state")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := &Stats{
		Labels: make([]*Label, 0),
	}
	for rows.Next() {
		label := &Label{}
		err = rows.Scan(&label.Id, &label.Name, &label.Amount, &label.AmountApproved)
		if err != nil {
			return nil, err
		}
		stats.Labels = append(stats.Labels, label)
		stats.TotalAmount += label.Amount
		stats.TotalAmountApproved += label.AmountApproved
	}

	return stats, nil
}
