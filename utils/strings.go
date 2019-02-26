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
	pascalCaseRE = regexp.MustCompile("^[A-Z][a-z0-9]+(?:[A-Z][a-z0-9]+|ID)*$")
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

var indentMatcher = regexp.MustCompile(`(?m)^([^\S\n]*)`)

func StripIndent(input string) string {
	// trim leading new lines
	input = strings.Trim(input, "\n")

	// trim trailing lines and whitespace
	input = strings.TrimRight(input, "\n\t ")

	// match all indents
	matches := indentMatcher.FindAllStringSubmatch(input, -1)
	if matches == nil {
		return input
	}

	// find the smallest indent size
	minIndent := len(matches[1][0])
	for _, match := range matches {
		if len(match[0]) < minIndent {
			minIndent = len(match[0])
		}
	}

	// trim the smallest indent from all lines
	output := []string{}
	for _, line := range strings.Split(input, "\n") {
		output = append(output, line[minIndent:])
	}

	return strings.Join(output, "\n")
}

func Yaml(input string) string {
	input = StripIndent(input)
	return strings.Replace(input, "\t", "  ", -1)
}
