package sql

import (
	"database/sql"
)

func ScanToSlice(rows *sql.Rows) ([]any, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}
	err = rows.Scan(valuePtrs...)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func ScanToMap(rows *sql.Rows) (map[string]any, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}
	err = rows.Scan(valuePtrs...)
	if err != nil {
		return nil, err
	}
	result := make(map[string]any)
	for i, column := range columns {
		result[column] = values[i]
	}
	return result, nil
}

func ScanToStruct(rows *sql.Rows, dest any) error {
	return nil
}
