package main

import (
	"errors"
	"fmt"
)

// ошибки отрицательного депозита и баланса
var errNegativeBalance = errors.New("Negative balance")
var errNegativeDeposit = errors.New("Negative deposit")

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner}
}

func (a *Account) SetBalance(quantity float64) error {
	if quantity < 0 {
		return errNegativeBalance
	}
	a.balance = quantity
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(quantity float64) error {
	if quantity < 0 {
		return errNegativeDeposit
	}
	a.balance += quantity // зачисление средств на депозит
	return nil
}

func (a *Account) Withdraw(quantity float64) error {
	if quantity < 0 {
		return errNegativeDeposit
	}
	if a.balance-quantity < 0 {
		return errNegativeDeposit
	}
	a.balance -= quantity // снятие средств с депозита
	return nil
}

func main() {
	a := &Account{owner: "owner"}
	a.SetBalance(100)
	fmt.Println("Вы пополнили баланс на", a.GetBalance(), "рублей")

	err := a.Deposit(10)
	if err != nil {
		panic(errNegativeDeposit)
	}

	err = a.Withdraw(100)
	if err != nil {
		panic(errNegativeDeposit)
	}

	fmt.Println("У вас осталось", a.GetBalance(), "рублей")
}
