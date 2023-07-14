package main

type Wallet struct {
	ballance int
}

func (w *Wallet) Deposit(ammount int){
	w.ballance += ammount
}

func (w *Wallet) Balance() int {
	return w.ballance
}