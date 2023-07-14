package main

import (
	"testing"
)

func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1,2,3,4,5}

		got := Sum(numbers[:])
		want := 15
		assertCorrectMessage(t, got, want)
	})

	t.Run("collection of 0 numbers", func(t *testing.T) {
		numbers := []int{}

		got := Sum(numbers[:])
		want := 0
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}