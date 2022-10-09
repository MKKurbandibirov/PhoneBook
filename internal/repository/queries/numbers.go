package queries

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"phonebook/internal/domain"
	"phonebook/pkg/parsing"
	"strconv"
	"strings"
)

const (
	SelectNumbersQuery       = `SELECT Id, FirstName, LastName, Country, Number FROM numbers`
	SelectNumbersQueryByName = `SELECT Id, FirstName, LastName, Country, Number FROM numbers WHERE FirstName = $1 AND LastName = $2`
)

func ParsParameters() (string, []string, error) {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a person full name: ")
	scanner.Scan()
	text := scanner.Text()
	if text == "" {
		return SelectNumbersQuery, nil, nil
	}
	name := strings.Split(text, " ")
	if name[0] == "" || name[1] == "" {
		return "", nil, errors.New("invalid name parameter")
	} else {
		return SelectNumbersQueryByName, name, nil
	}
}

func (q *Queries) Numbers(ctx context.Context) ([]domain.Number, error) {
	filter, err := parsing.ParsFilter()
	if err != nil {
		return nil, err
	}
	query, value, err := ParsParameters()
	if err != nil {
		return nil, err
	}
	var rows pgx.Rows
	if value != nil {
		rows, err = q.pool.Query(ctx, query, value[0], value[1])
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

	if len(numbers) > filter.Limit && filter.Limit > 0 {
		numbers = numbers[:filter.Limit]
	}
	return numbers, nil
}

const InsertNumberQuery = `INSERT INTO numbers (FirstName, LastName, Country, Number) VALUES($1, $2, $3, $4) RETURNING Id`

func GetParameters() []string {
	fmt.Print("Enter a person parameters:\n")
	var scanner = bufio.NewScanner(os.Stdin)
	var params = make([]string, 0)
	for _, param := range [...]string{"FirsName", "LastName", "Country", "Number"} {
		fmt.Printf("\t%s:", param)
		scanner.Scan()
		params = append(params, scanner.Text())
	}
	return params
}

func (q *Queries) AddNumber(ctx context.Context) (int, error) {
	params := GetParameters()
	num, err := strconv.Atoi(params[3])
	if err != nil {
		return -1, err
	}
	row := q.pool.QueryRow(ctx, InsertNumberQuery, params[0], params[1], params[2], num)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetNumberId() (int, error) {
	fmt.Print("Enter a person Id:\n\t")
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	id, err := strconv.Atoi(text)
	if err != nil {
		return -1, err
	}
	return id, nil
}

const UpdateNumberQuery = `UPDATE numbers SET FirstName = $1, LastName = $2, Country = $3, Number = $4  WHERE Id = $5`

func (q *Queries) ChangeNumber(ctx context.Context) error {
	id, err := GetNumberId()
	if err != nil {
		return err
	}
	params := GetParameters()
	num, err := strconv.Atoi(params[3])
	if err != nil {
		return err
	}
	ct, err := q.pool.Exec(ctx, UpdateNumberQuery, params[0], params[1], params[2], num, id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("no one row affected")
	}
	return nil
}

const DeleteNumberQuery = `DELETE FROM numbers WHERE id = $1`

func (q *Queries) DeleteNumber(ctx context.Context) error {
	id, err := GetNumberId()
	if err != nil {
		return err
	}
	ct, err := q.pool.Exec(ctx, DeleteNumberQuery, id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("no one row affected")
	}
	return nil
}
