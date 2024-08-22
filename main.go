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
	answers := createWordList("words/answer-list.txt")
	answerMap := convertToMap(answers)
	// guesses := createWordList("words/full-word-list.txt")
	answer := answers[rand.Intn(len(answers))]

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Welcome to Wordle CLI!")

	for attempt := MaxGuesses; attempt > 0; attempt-- {
		input, err := reader.ReadString('\n')
		if len(input) > MaxWordLength {
			fmt.Println("Word can only be 5 characters")
			continue
		}

		if !isValidWord(input, answerMap) {
			fmt.Println("Invalid word")
			continue
		}
		if input == answer {
			fmt.Println("Congrats you got the correct word!")
			return
		}

		if err != nil {
			fmt.Println("Please try again.")
			continue
		}
	}

	fmt.Println("Game over! You ran out of attempts")

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
