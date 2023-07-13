package main

import (
	"fmt"
)

func main(){
	A := NewBankingAccount(100)

	B := NewBankingAccount(100)
	TransactionA, TransactionB := NewTransaction[*BankingAccount, *BankingAccount](A, B, 100, "usd")
	bankXMarket.transactions["0"] = TransactionA
	bankXMarket.transactions["1"] = TransactionB
	banks.Append(&bankXMarket)
	fmt.Println(&banks.relation)
}
