package internal

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{balance}, &Memento{balance}
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

func RunMemento() {
	ba, m0 := NewBankAccount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)
	fmt.Println(*ba)
	// restore to m1
	ba.Restore(m1)
	fmt.Println(*ba)
	// restore to m2
	ba.Restore(m2)
	fmt.Println(*ba)
	// restore to m0
	ba.Restore(m0)
	fmt.Println(*ba)
}
