package graphqlQuery

import "strings"

func lineShift(line string) string {
	lines := strings.Split(line, "\n")
	return strings.Join(lines, "\n  ")
}
