package main

import (
	"testing"
)

func TestArea(t *testing.T){
	t.Run("area of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Width: 15.2,
			Height: 13.2,
		}

		got := Area(rectangle)
		want := 200.64
		assertCorrectMessage(t, got, want)
	})

	t.Run("area of empty Rectangle", func(t *testing.T) {
		rectangle := Rectangle{}

		got := Area(rectangle)
		want := 0.0
		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkArea(b *testing.B) {
	rectangle := Rectangle{
		Width: 23.2,
		Height: 16.3,
	}

	for i := 0; i < b.N; i++ {
		Area(rectangle)
	}
}

func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}