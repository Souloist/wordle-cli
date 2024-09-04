package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"wordle-cli/internal/game"
	"wordle-cli/internal/ui"
	"wordle-cli/internal/utils"
)

func main() {
	var previousGuesses []string

	answers := utils.CreateWordList("words/answer-list.txt")
	guesses := utils.CreateWordList("words/full-word-list.txt")
	guessesSet := utils.ListToSet(guesses)
	answer := answers[rand.Intn(len(answers))]

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Wordle CLI! ヾ(＾ ∇ ＾).")

	for attempt := game.MaxGuesses; attempt > 0; attempt-- {
		fmt.Print("Please type your guess: ")
		guess, err := reader.ReadString('\n')
		guess = strings.TrimSpace(strings.ToLower(guess))

		if guess == "exit" {
			fmt.Println("Bye!")
			return
		}

		if len(guess) > game.MaxWordLength {
			fmt.Println("Word can only be 5 characters!")
			attempt++
			continue
		}

		if !game.IsValidWord(guess, guessesSet) {
			fmt.Println("This word cannot be found in our word list")
			attempt++
			continue
		}

		if err != nil {
			fmt.Println("Something went wrong. Please try again.")
			attempt++
			continue
		}

		guessColorArray := game.GetColorArrayForGuess(guess, answer)
		formattedGuess, _ := ui.ColoredString(guess).FormatWithColorArray(guessColorArray)
		previousGuesses = append(previousGuesses, formattedGuess)
		utils.ClearTerminal()
		for _, guess := range previousGuesses {
			fmt.Println(guess)
		}

		if guess == answer {
			fmt.Println("Congrats you got the correct word! 🎉")
			return
		}
	}

	fmt.Println("Game over ☠️! You ran out of attempts. ( • ᴖ • ｡) ")
	fmt.Println("The Answer was:", answer)
}
