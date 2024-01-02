package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("Choose your move: (1) Rock, (2) Paper, (3) Scissors, (4) Quit")
		var Choice int
		fmt.Scan(&Choice)

		switch Choice {
		case 1:
			playRound("Rock")
		case 2:
			playRound("Paper")
		case 3:
			playRound("Scissors")
		case 4:
			fmt.Println("Thanks for playing. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please choose a number between 1 and 4.")
		}
	}
}

func playRound(userMove string) {
	choices := []string{"Rock", "Paper", "Scissors"}
	computerMove := choices[rand.Intn(len(choices))]

	fmt.Printf("You chose %s.\n", userMove)
	fmt.Printf("Computer chose %s.\n", computerMove)

	switch {
	case userMove == computerMove:
		fmt.Println("It's a tie!")
	case (userMove == "Rock" && computerMove == "Scissors") ||
		(userMove == "Paper" && computerMove == "Rock") ||
		(userMove == "Scissors" && computerMove == "Paper"):
		fmt.Println("You win!")
	default:
		fmt.Println("You lose!")
	}
}