package main

import "fmt"

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner}
}

func (a *Account) SetBalance(quantity float64) error {
	if quantity < 0 {
		fmt.Println("Отрицательный баланс")
	}
	a.balance = quantity
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(quantity float64) error {
	if quantity < 0 {
		fmt.Println("Отрицательный или нулевой депозит")
	}
	a.balance += quantity // зачисление средств
	return nil
}

func (a *Account) Withdraw(quantity float64) error {
	if quantity < 0 {
		fmt.Println("Отрицательный депозит")
	}
	if a.balance-quantity < 0 {
		fmt.Println("Отрицательный депозит")
	}
	a.balance -= quantity // снятие средств с депозита
	return nil
}

func main() {
	a := &Account{owner: "owner"}
	a.SetBalance(100)

	fmt.Println(a.GetBalance())
}
