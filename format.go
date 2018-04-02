package main

import "strings"

func indentFunc(depth int) string {
	if depth < 0 {
		depth = 0
	}

	indent := "    "
	return strings.Repeat(indent, depth)
}
