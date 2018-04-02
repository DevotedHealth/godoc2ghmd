package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndentFunc(t *testing.T) {

	tests := []struct {
		depth    int
		expected string
	}{
		{-1, ""},
		{0, ""},
		{1, "    "},
		{4, "                "},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			assert.Equal(t, test.expected, indentFunc(test.depth))
		})
	}
}
