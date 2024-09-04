package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CreateWordList(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(data), "\n")
	return list
}

// ClearTerminal clears the terminal screen depending on the OS
func ClearTerminal() {
	var clearCmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default: // Linux, macOS, etc.
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	err := clearCmd.Run()
	if err != nil {
		return
	}
}

func ListToSet(wordList []string) map[string]struct{} {
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	return wordSet
}
