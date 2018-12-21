package utils

import (
	"regexp"
	"strings"
)

// ContainsUpperChars returns true if the input
// string contains any upper-case letters.
func ContainsUpperChars(input string) bool {
	return strings.ToLower(input) != input
}

var (
	pascalCaseRE = regexp.MustCompile("^[A-Z][a-z0-9]+(?:[A-Z][a-z0-9]+)*$")
	camelCaseRE  = regexp.MustCompile("^[a-z][a-z0-9]+(?:[A-Z0-9]+[a-z0-9]*)*$")
	snakeCaseRE  = regexp.MustCompile("^[a-z]+(_[a-z0-9]+)*$")
	kebabCaseRE  = regexp.MustCompile("^[a-z]+(-[a-z0-9]+)*$")
)

func IsPascalCase(input string) bool {
	return pascalCaseRE.MatchString(input)
}

func IsCamelCase(input string) bool {
	return camelCaseRE.MatchString(input)
}

func IsSnakeCase(input string) bool {
	return snakeCaseRE.MatchString(input)
}

func IsKebabCase(input string) bool {
	return kebabCaseRE.MatchString(input)
}
