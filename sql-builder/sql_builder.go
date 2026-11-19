package sqlbuilder

import (
	"fmt"
	"strings"
)

// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
	if len(columns) == 0 {
		return fmt.Sprintf("SELECT * FROM %s;", table)
	}

	return fmt.Sprintf("SELECT %s FROM %s;", strings.Join(columns, ", "), table)
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	result := make([]string, 0)

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		var columns strings.Builder
		var values strings.Builder

		for i, item := range row {
			if i > 0 {
				columns.WriteString(", ")
				values.WriteString(", ")
			}

			columns.WriteString(item.Column)
			values.WriteString("'")
			values.WriteString(item.Value)
			values.WriteString("'")
		}
		query := fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s);",
			table,
			columns.String(),
			values.String(),
		)
		result = append(result, query)
	}

	return result
}
