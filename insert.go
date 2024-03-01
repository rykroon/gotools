package gotools

import (
	"fmt"
	"strings"
)

func BuildInsertQuery(table string, columns []string) string {
	columnsStr := strings.Join(columns, ", ")
	valuesStr := strings.Join(strings.Split(strings.Repeat("?", len(columns)), ""), ", ")
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columnsStr, valuesStr)
}
