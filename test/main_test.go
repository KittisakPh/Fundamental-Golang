package main

import (
	"testing"
	// "github.com/stretchr/testify"
)

func TestAdd(t *testing.T) {
	t.Run("add numbers", func(t *testing.T) {
		r := Add(1, 2)
		if r != 3 {
			t.Error("Add(1,2) should be 3")
		}
	})

	t.Run("when input 1 and 2 should return 3", func(t *testing.T) {
		r := Add(1, 2)
		if r != 3 {
			t.Error("Add(1,2) should be 3")
		}
	})
}