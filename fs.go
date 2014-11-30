package main

import (
	"strings"
)

func SplitPath(path string) []string {
	path = strings.ToLower(path)

	isSeparator := func(c rune) bool {
		return c == '/' || c == '.'
	}

	return strings.FieldsFunc(path, isSeparator)
}
