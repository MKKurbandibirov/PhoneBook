package parsing

import (
	"bufio"
	"fmt"
	"os"
	"phonebook/internal/domain"
	"strconv"
	"strings"
)

func ParsQuery() domain.Operation {
	var queryType string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a query name:\n\t1: get;\n\t2: insert;\n\t3: update;\n\t4: delete\n\t5: exit\n")
	scanner.Scan()
	queryType = strings.ToLower(scanner.Text())
	switch queryType {
	case "get":
		return domain.Operation("get")
	case "1":
		return domain.Operation("get")
	case "insert":
		return domain.Operation("insert")
	case "2":
		return domain.Operation("insert")
	case "update":
		return domain.Operation("update")
	case "3":
		return domain.Operation("update")
	case "delete":
		return domain.Operation("delete")
	case "4":
		return domain.Operation("delete")
	case "exit":
		return domain.Operation("exit")
	case "5":
		return domain.Operation("exit")
	default:
		return ""
	}
}

func ParsFilter() (*domain.NumbersFilter, error) {
	var limit int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a count of rows: ")
	scanner.Scan()
	text := scanner.Text()
	if text == "" {
		return &domain.NumbersFilter{Limit: -1}, nil
	}
	limit, err := strconv.Atoi(text)
	if err != nil {
		return nil, err
	}
	return &domain.NumbersFilter{Limit: limit}, nil
}
