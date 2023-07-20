package main

import (
	"github.com/google/uuid"
)

func NewCryptoAccount(amount int) *CryptoAccount {
	id := uuid.New()

	A := CryptoAccount{
		CryptoAddr: id.String(),
		Amount:     amount,
		Market:     &cryptoXMarket,
	}
	return &A
}

func NewBankingAccount(amount int) *BankingAccount {
	id := uuid.New()

	A := BankingAccount{
		Id:     id.String(),
		Amount: amount,
		Market: &bankXMarket,
	}
	return &A
}
