package internal

import "fmt"

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for i := range c.commands {
		c.commands[len(c.commands)-1-i].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceded(value)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))
	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeded()
		} else {
			cmd.SetSucceded(false)
		}
	}
}

func RunCompositeCommands() {
	from := BankAccount{100}
	to := BankAccount{0}
	mtc := NewMoneyTransferCommand(&from, &to, 25)
	mtc.Call()
	fmt.Println(from, to)

	mtc.Undo()
	fmt.Println(from, to)
}
