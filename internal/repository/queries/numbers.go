package queries

import (
	"context"
	"phonebook/internal/domain"
)

const selectNumbersQuery = `SELECT Id, FirstName, LastName, Country, Number FROM numbers`

func (q *Queries) Numbers(ctx context.Context) ([]domain.Number, error) {
	rows, err := q.pool.Query(ctx, selectNumbersQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var numbers []domain.Number
	for rows.Next() {
		var n domain.Number
		err = rows.Scan(&n.Id, &n.FirstName, &n.LastName, &n.Country, &n.Number)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
