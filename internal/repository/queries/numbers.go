package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"phonebook/internal/domain"
	"strings"
)

const (
	SelectNumbersQuery           = `SELECT Id, FirstName, LastName, Country, Number FROM numbers`
	SelectNumbersQueryByLastName = `SELECT Id, FirstName, LastName, Country, Number FROM numbers WHERE LastName = $1`
	SelectNumbersQueryByCountry  = `SELECT Id, FirstName, LastName, Country, Number FROM numbers WHERE Country = $1`
)

func ParsParameters() (string, string, error) {
	var param string
	fmt.Printf("Enter a parameter name:\n\t1: Last-name;\n\t2: Country\n")
	_, err := fmt.Scan(&param)
	if err != nil {
		return "", "", err
	}
	param = strings.ToLower(param)
	if param == "last-name" || param == "lastname" || param == "1" {
		var value string
		fmt.Printf("\tLastName: ")
		fmt.Scan(&value)
		return SelectNumbersQueryByLastName, value, nil
	} else if param == "country" || param == "2" {
		var value string
		fmt.Printf("\tCountry: ")
		fmt.Scan(&value)
		return SelectNumbersQueryByCountry, value, nil
	}
	return SelectNumbersQuery, "", nil
}

func (q *Queries) Numbers(ctx context.Context) ([]domain.Number, error) {
	query, value, err := ParsParameters()
	if err != nil {
		return nil, err
	}
	var rows pgx.Rows
	if value != "" {
		rows, err = q.pool.Query(ctx, query, value)
	} else {
		rows, err = q.pool.Query(ctx, query)
	}
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
