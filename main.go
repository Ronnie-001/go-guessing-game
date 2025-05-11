package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the number guessing game! You have 5 guesses")
	
	numOfGuesses := 5
	var hasWon bool = false

	value := rand.Intn(101)
	for numOfGuesses > 0 {
		fmt.Println("Guess a number between 1 and 100 (inclusive)")
		
		// Getting the users guess
		var guess int
		fmt.Scanln(&guess)
		
		if guess < value {
			fmt.Println("Guess is too low! Try again!")
		}

		if guess > value {
			fmt.Println("Guess is too high! Try again!")
		}
		
		if guess == value {
			fmt.Println("Guess is correct! You win!")
			hasWon = true
			break
		}
		numOfGuesses--
	}

	if numOfGuesses == 0 && !hasWon {
		fmt.Println("You ran out of guesses! You Lose")
	}
}
