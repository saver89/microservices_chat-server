package prettier

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// PlaceholderDollar is a placeholder for PostgreSQL
	PlaceholderDollar = "$"
	// PlaceholderQuestion is a placeholder for MySQL
	PlaceholderQuestion = "?"
)

// Pretty replaces placeholders in the query with the values of the arguments
func Pretty(query string, placeholder string, args ...any) string {
	for i, param := range args {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("%q", v)
		case []byte:
			value = fmt.Sprintf("%q", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}

		query = strings.Replace(query, fmt.Sprintf("%s%s", placeholder, strconv.Itoa(i+1)), value, -1)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.TrimSpace(query)
}
