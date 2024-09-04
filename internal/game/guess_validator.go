package game

import "strings"
import "wordle-cli/internal/ui"

func IsValidWord(word string, wordMap map[string]struct{}) bool {
	if _, exists := wordMap[word]; exists {
		return true
	}
	return false
}

func GetColorArrayForGuess(input string, answer string) []ui.ANSIColor {
	var colorArray []ui.ANSIColor
	for index, char := range input {
		if string(char) == string(answer[index]) {
			colorArray = append(colorArray, ui.ANSIColorExactMatch)
		} else if strings.Contains(answer, string(char)) {
			colorArray = append(colorArray, ui.ANSIColorExistMatch)
		} else {
			colorArray = append(colorArray, ui.ANSIColorNoMatch)
		}
	}
	return colorArray
}
