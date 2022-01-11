package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	answers := createWordList("files/answer-list.txt")
	//guesses := createWordList("files/full-word-list.txt")
	answer := answers[rand.Intn(len(answers))]

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please type your guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Please try again.")
	}

	if input == answer {
		fmt.Println("Congrats!")
	}

}

func createWordList(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(data), "\n")
	return list
}
