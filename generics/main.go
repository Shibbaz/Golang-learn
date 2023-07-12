package main

import(
    "github.com/google/uuid"
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

func NewTransaction[A Account, B Account](Source A, Direction B, Amount int, Currency string) (Transaction[A, B], Transaction[B, A]){
	TransactionA := Transaction[A, B]{
		Source: Source,
		Direction: Direction,
		Amount: Amount,
		Currency: Currency,
	}

	TransactionB := Transaction[B, A]{
		Source: Direction,
		Direction: Source,
		Amount: -1*Amount,
		Currency: Currency,
	}
	return TransactionA, TransactionB
}
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

var banks = BankMarket{
	relation: map[string]Market[*BankingAccount]{},
}

type CryptoMarket struct{
	relation map[string]Market[*CryptoAccount]
}


type BankMarket struct {
	relation map[string]Market[*BankingAccount]
}

func (cryptoMarket *CryptoMarket) Append(market *Market[*CryptoAccount]) {
	crypto := *cryptoMarket
	crypto.relation[uuid.New().String()] = *market
}

func (cryptoMarket *CryptoMarket) Delete(key string){
	delete(cryptoMarket.relation,key)
}

func (cryptoMarket *CryptoMarket) Update(key string, market *Market[*CryptoAccount]) {
	cryptoMarket.relation[key] = *market
}


func (bankMarket *BankMarket) Append(market *Market[*BankingAccount]) {
	(bankMarket.relation)[uuid.New().String()] = *market
}

func (bankMarket *BankMarket) Delete(key string){
	delete(bankMarket.relation,key)
}

func (bankMarket *BankMarket) Update(key string, market *Market[*BankingAccount]) {
	(bankMarket.relation)[key] = *market
}

func NewCryptoAccount(amount int) *CryptoAccount{
	id := uuid.New()

	A := CryptoAccount{
		CryptoAddr: id.String(),
		Amount: amount,
		Market: &cryptoXMarket,
	}
	return &A
}

func NewBankingAccount(amount int) *BankingAccount{
	id := uuid.New()

	A := BankingAccount{
		Id: id.String(),
		Amount: amount,
		Market: &bankXMarket,
	}
	return &A
}

func main(){
	A := NewBankingAccount(100)

	B := NewBankingAccount(100)
	TransactionA, TransactionB := NewTransaction[*BankingAccount, *BankingAccount](A, B, 100, "usd")
	bankXMarket.transactions["0"] = TransactionA
	bankXMarket.transactions["1"] = TransactionB
	banks.Append(&bankXMarket)
	fmt.Println(&banks.relation)
}
