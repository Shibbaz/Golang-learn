package main
import(
	"fmt"
)

type BankingAccount struct{
	bankAddr string
	amount	int
}

type CryptoAccount struct {
	cryptoAddr string
	amount int
}

type Account interface{
	*BankingAccount | *CryptoAccount
}

type Transaction[A Account, B Account] struct{
	from A
	to B
	amount int
	currency string
}

func NewTransaction[A Account, B Account](from A, to B, amount int, currency string) *Transaction[A, B]{
	return &Transaction[A, B]{
		from: from,
		to: to,
		amount: amount,
		currency: currency,
	}
}

func main(){
	bank := BankingAccount{
		bankAddr: "XD",
		amount: 100,
	}
	
	crypto := CryptoAccount{
		cryptoAddr: "YD",
		amount: 200,
	}
	transaction := Transaction[*BankingAccount, *CryptoAccount]{
		from: &bank,
		to: &crypto,
		amount: 100,
		currency: "usd",
	}
	fmt.Println(transaction.currency)
}