package stringutils

import "strings"

//Upper converts string to uppercase
func Upper(s string) string {
	return strings.ToUpper(s)
}

//Lower converts string to lowercase
func Lower(s string) string {
	return strings.ToLower(s)
}
