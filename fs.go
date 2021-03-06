package main

import (
	"strings"
)

type Tags []string

type OrderedTags map[string]byte

func OrderTags(tags Tags) OrderedTags {
	t := make(OrderedTags)
	for i, tag := range tags {
		weight := (len(tags) - 1 - i) * 255 / (len(tags) - 1)
		t[tag] += byte(weight)
	}

	return t
}

func SplitPath(path string) []string {
	isSeparator := func(c rune) bool {
		return c == '/' || c == '.'
	}

	return strings.FieldsFunc(path, isSeparator)
}
