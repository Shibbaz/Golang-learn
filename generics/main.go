package main

import(
	"fmt"
)

type Bank struct{
	Name string
	NetWorth int
	Addr string
}

type Market[T Account] struct{
	Name string
	Volume int64
	Hash string
	transactions map[string]Transaction[T, T]
}

func (market *Market[Account]) Transactions()(*map[string]Transaction[Account, Account]){
	return &market.transactions
}


type BankingAccount struct{
	Id string
	Bank *Bank
	Amount	int
	Market *Market[*BankingAccount]
}

type CryptoAccount struct {
	CryptoAddr string
	Amount int
	Market *Market[*CryptoAccount]
}

type Account interface{
	*BankingAccount | *CryptoAccount
}

type Transaction[A Account, B Account] struct{
	Source A
	Direction B
	Amount int
	Currency string
}

func NewTransaction[A Account, B Account](Source A, Direction B, Amount int, Currency string) Transaction[A, B]{
	return Transaction[A, B]{
		Source: Source,
		Direction: Direction,
		Amount: Amount,
		Currency: Currency,
	}
}

func main(){

	market := Market[*CryptoAccount]{
		Name: "Shibba",
		Volume: 1092222222222222222,
		Hash: "X827273ghjsgjdsgjsgks",
	}

	A := CryptoAccount{
		CryptoAddr: "YD",
		Amount: 200,
		Market: &market,
	}

	B := CryptoAccount{
		CryptoAddr: "YD",
		Amount: 200,
		Market: &market,
	}

	market.transactions = map[string]Transaction[*CryptoAccount, *CryptoAccount]{
		"0": NewTransaction[*CryptoAccount, *CryptoAccount](&A, &B, 100, "usd"),
	}
	fmt.Println(market.Transactions())
}