package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var previousGuesses []string

	answers := createWordList("words/answer-list.txt")
	guesses := createWordList("words/full-word-list.txt")
	guessesMap := convertToMap(guesses)
	answer := answers[rand.Intn(len(answers))]

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Wordle CLI! ヾ(＾ ∇ ＾).\n")

	for attempt := MaxGuesses; attempt > 0; attempt-- {
		fmt.Print("Please type your guess: ")
		guess, err := reader.ReadString('\n')
		guess = strings.TrimSpace(strings.ToLower(guess))
		if len(guess) > MaxWordLength {
			fmt.Println("Word can only be 5 characters!")
			attempt++
			continue
		}

		if !isValidWord(guess, guessesMap) {
			fmt.Println("This word cannot be found in our word list")
			attempt++
			continue
		}

		if err != nil {
			fmt.Println("Something went wrong. Please try again.")
			attempt++
			continue
		}
		formattedGuess := getGuessDisplay(guess, answer)
		previousGuesses = append(previousGuesses, formattedGuess)
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

func createWordList(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(data), "\n")
	return list
}

func isValidWord(word string, wordMap map[string]struct{}) bool {
	if _, exists := wordMap[word]; exists {
		return true
	}
	return false
}

func convertToMap(wordList []string) map[string]struct{} {
	wordMap := make(map[string]struct{})
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	return wordMap
}
