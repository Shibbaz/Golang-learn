package main

import(
    "github.com/google/uuid"
)

var cryptoXMarket = Market[*CryptoAccount]{
	Name: "Shibba",
	Volume: 1092222222222222222,
	Hash: uuid.New().String(),
	transactions: map[string]Transaction[*CryptoAccount, *CryptoAccount]{},
}

var bankXMarket = Market[*BankingAccount]{
	Name: "Shibba",
	Volume: 1092222222222222222,
	Hash: uuid.New().String(),
 	transactions: map[string]Transaction[*BankingAccount, *BankingAccount]{},
}