package main

import (
	"testing"
)

func TestSearch(t *testing.T){
	t.Run("Search method", func(t *testing.T) {
		dictionary := map[string]string{
			"0": "appa",
			"1": "aang",
			"2": "korra",
		}
		got, _ := Search(dictionary, "0")
		want := "appa"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}