package main

import (
	"fmt"
	"errors"
)

// Wallet
type Wallet struct {
	balance		float64
	currency	string
}

// Deposit (amount float64)
func (w *Wallet) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("az összeg csak pozitív lehet")
	} else {
		w.balance = w.balance + amount
	}
	return nil
}

// Withdraw(amount float64) err
func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return fmt.Errorf("nincs elég fedezet. Jelenlegi egyenleged: %.2f", w.balance)
	} else {
		w.balance = w.balance - amount
	}
	return nil
}

// GetBalance() float64
func (w Wallet) GetBalance() float64 {
	return w.balance
}
// GetInfo() string
func (w Wallet) GetInfo() string {
	return fmt.Sprintf("A jelenlegi egyenleg %.2f %s", w.balance, w.currency)
}

// Transfer(target *Wallet, amount float64) error
func (w1 *Wallet) Transfer(amount float64, w2 *Wallet) error {
	if amount > w1.balance {
		return errors.New("nincs elegendő fedezet")
	} else {
		w1.balance = w1.balance - amount
		w2.balance = w2.balance + amount
		return nil
	}
}


//TODO:
// Hozz létre két pénztárcát (pl. 10000 HUF és 5000 HUF)
// Fizess be az egyikbe
// Vegyél ki a másikból
// Utalj át egyik pénztárcából a másikba
// Írd ki mindkettő info-ját
// Próbálj többet kivenni mint amennyi van (error handling!)

func main() {

	wallet1 := Wallet{
		balance: 10000,
		currency: "HUF",
	}

	wallet2 := Wallet{
		balance: 5000,
		currency: "HUF",
	}

	wallet1.Deposit(3000)

	wallet2.Withdraw(2000)

	fmt.Printf("Jelenlegi egyenleg: %.2f\n", wallet1.GetBalance())
	fmt.Println(wallet2.GetInfo())

	err := wallet1.Transfer(2000, &wallet2)
	if err != nil {
		fmt.Println("Transfer hiba:", err)
	} else {
		fmt.Println("Sikeres átutalás!")
	}

	err1 := wallet1.Withdraw(500000)
	if err1 != nil {
		fmt.Println("Withdraw hiba:", err1)
	}

	fmt.Println(wallet1.GetInfo())
	fmt.Println(wallet2.GetInfo())
}
