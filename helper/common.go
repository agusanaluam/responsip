package helper

import (
	"regexp"
	"strings"
)

var letters = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// SetLowerAndAddSpace is function to convert camelCase to snake_case
func setLowerAndAddSpace(str string) string {
	lower := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	lower = matchAllCap.ReplaceAllString(lower, "${1} ${2}")
	return strings.ToLower(lower)
}
