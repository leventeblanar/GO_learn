package main

import (
	"fmt"
	"errors"
)

type Account struct {
	owner		string
	balance		float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("%s befizetett %.2f. Új egyenleg: %.2f\n", a.owner, amount, a.balance)
}

func (a *Account) Withdrawn(amount float64) error {
	if a.balance < amount {
		return errors.New("nincs elég fedezet a számlán")
	}
	a.balance -= amount
	fmt.Printf("%s kivett %.2f. Új egyenleg: %.2f\n", a.owner, amount, a.balance)
	return nil
}

func (a Account) GetBalance() float64 {
	return a.balance
}

func main() {
	lblanarAccount := Account{
		owner: "Blanar Levente",
		balance: 10000.0,
	}

	fmt.Printf("Kezdő egyenleg: %.2f\n\n", lblanarAccount.GetBalance())
	lblanarAccount.Deposit(200)

	lblanarAccount.Withdrawn(100)
}