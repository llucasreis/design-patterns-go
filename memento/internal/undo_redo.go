package internal

import "fmt"

type Memento2 struct {
	Balance int
}

type BankAccount2 struct {
	balance int
	changes []*Memento2
	current int
}

func NewBankAccount2(balance int) *BankAccount2 {
	ba := &BankAccount2{balance: balance}
	ba.changes = append(ba.changes, &Memento2{balance})
	return ba
}

func (b *BankAccount2) String() string {
	return fmt.Sprint("Balance = $", b.balance, ", current = ", b.current)
}

func (b *BankAccount2) Deposit(amount int) *Memento2 {
	b.balance += amount
	m := Memento2{b.balance}
	b.changes = append(b.changes, &m)
	b.current++
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
	return &m
}

func (b *BankAccount2) Restore(m *Memento2) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount2) Undo() *Memento2 {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount2) Redo() *Memento2 {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func RunUndoRedo() {
	ba := NewBankAccount2(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1:", ba)

	ba.Undo()
	fmt.Println("Undo 2:", ba)

	ba.Redo()
	fmt.Println("Redo 2:", ba)
}
