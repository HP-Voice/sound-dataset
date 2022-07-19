package main

type Stats struct {
	Labels      []*Label
	TotalAmount int
}

func readStats() (*Stats, error) {
	rows, err := db.Query("SELECT id, name, amount FROM label_state")
	if err != nil {
		return nil, err
	}

	stats := &Stats{
		Labels: make([]*Label, 0),
	}
	for rows.Next() {
		label := &Label{}
		err = rows.Scan(&label.Id, &label.Name, &label.Amount)
		if err != nil {
			return nil, err
		}
		stats.Labels = append(stats.Labels, label)
		stats.TotalAmount += label.Amount
	}

	return stats, nil
}
