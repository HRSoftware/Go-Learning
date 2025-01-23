package games

import (
	"Learning/utils"
	"fmt"
	"math/rand"
)

func Play() {
	fmt.Println("Welcome to the Number guessing game! Guess a number between 0 and 100")

	guessNum := 0
	num := rand.Intn(100)

	for guessNum != num {
		guessNum, _ = utils.GetUserInputNumber()

		switch {
		case guessNum == num:
			fmt.Println("Congratulations - You guessed the number!")
		case guessNum > num:
			fmt.Println("Too high")
		case guessNum < num:
			fmt.Println("Too low")
		}
	}
}
