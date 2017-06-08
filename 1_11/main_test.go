package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFIterative0Through5(t *testing.T) {
	for i := 0; i <= 15; i++ {
		assert.Equal(t, FRecursive(i), FIterative(i))
	}
}
