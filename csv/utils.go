package csv

import "strings"

func ToRow(s []string) string {
	return strings.Join(s, ",")
}
