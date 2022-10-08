package parsing

import (
	"fmt"
	"phonebook/internal/domain"
	"strings"
)

func ParsQuery() (domain.Operation, error) {
	var queryType string
	fmt.Printf("Enter a query name:\n\t1: get;\n\t2: insert;\n\t3: update;\n\t4: delete\n\t5: exit\n")
	_, err := fmt.Scan(&queryType)
	if err != nil {
		return "", err
	}
	queryType = strings.ToLower(queryType)
	switch queryType {
	case "get":
		return domain.Operation("get"), nil
	case "1":
		return domain.Operation("get"), nil
	case "insert":
		return domain.Operation("insert"), nil
	case "2":
		return domain.Operation("insert"), nil
	case "update":
		return domain.Operation("update"), nil
	case "3":
		return domain.Operation("update"), nil
	case "delete":
		return domain.Operation("delete"), nil
	case "4":
		return domain.Operation("delete"), nil
	case "exit":
		return domain.Operation("exit"), nil
	case "5":
		return domain.Operation("exit"), nil
	default:
		return "", nil
	}
}

func ParsFilter() (*domain.NumbersFilter, error) {
	var limit int
	fmt.Printf("Enter a count of rows: ")
	_, err := fmt.Scan(&limit)
	if err != nil {
		return nil, err
	}
	return &domain.NumbersFilter{Limit: limit}, nil
}
