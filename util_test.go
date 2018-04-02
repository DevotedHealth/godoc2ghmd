package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	assert.Panics(t, func() { Must(errors.New("foo")) })
	assert.NotPanics(t, func() { Must(nil) })
}
