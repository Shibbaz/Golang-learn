package main

import (
	"testing"
)

func BenchmarkTestWallet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wallet := Wallet{
			ballance: 100,
		}
		wallet.Deposit(100)
		got := wallet.Balance()
		want := 200
		assertCorrectMessage(b, got, want)
	}
}

func TestWallet(t *testing.T){
	t.Run("Deposit method of wallet", func(t *testing.T) {
		wallet := Wallet{
			ballance: 100,
		}
		wallet.Deposit(100)
		got := wallet.Balance()
		want := 200
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}