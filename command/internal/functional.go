package internal

import "fmt"

type BankAccount2 struct {
	Balance int
}

func FunctionalDeposit(ba *BankAccount2, amount int) {
	fmt.Println("Depositing", amount)
	ba.Balance += amount
}

func FunctionalWithdraw(ba *BankAccount2, amount int) {
	if ba.Balance >= amount {
		fmt.Println("Withdrawing", amount)
		ba.Balance -= amount
	}
}

func RunFunctional() {
	ba := &BankAccount2{0}
	var commands []func()

	commands = append(commands, func() {
		FunctionalDeposit(ba, 100)
	})

	commands = append(commands, func() {
		FunctionalWithdraw(ba, 25)
	})

	for _, cmd := range commands {
		cmd()
	}

	fmt.Println(ba)
}
