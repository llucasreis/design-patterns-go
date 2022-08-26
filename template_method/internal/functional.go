package internal

import "fmt"

func PlayGame2(start, takeTurn func(),
	haveWinner func() bool, winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}

	fmt.Printf("Player %d wins!\n", winningPlayer())
}

func RunFunctional() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting game of chess")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n", turn, currentPlayer)
		currentPlayer = 1 - currentPlayer
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame2(start, takeTurn, haveWinner, winningPlayer)

}
