package main

import (
	"strings"
)

func getANSIColor(c ANSIColor) string {
	switch c {
	case ANSIColorExactMatch:
		return "\033[32m" // green
	case ANSIColorExistMatch:
		return "\033[33m" // yellow
	case ANSIColorNoMatch:
		return "\033[37m" // grey
	case ANSIColorUnknown:
		return "\033[97m" // white
	default:
		panic("Not supported color")
	}
}

func formatCharWithColor(s string, c ANSIColor) string {
	return ANSIBold + getANSIColor(c) + s + ANSIColorReset
}

func formatWordWithColor(s string, colorArray []ANSIColor) string {
	var arr []string
	if len(s) != len(colorArray) {
		panic("String and colorArray have different lengths")
	}
	for i := range s {
		arr = append(arr, formatCharWithColor(string(s[i]), colorArray[i]))
	}
	return strings.Join(arr, " ")
}

func getColorArrayForWord(input string, answer string) []ANSIColor {
	var colorArray []ANSIColor
	for index, char := range input {
		if string(char) == string(answer[index]) {
			colorArray = append(colorArray, ANSIColorExactMatch)
			continue
		}
		if strings.Contains(answer, string(char)) {
			colorArray = append(colorArray, ANSIColorExistMatch)
		} else {
			colorArray = append(colorArray, ANSIColorNoMatch)
		}
	}
	return colorArray
}

func getGuessDisplay(input string, answer string) string {
	colorArray := getColorArrayForWord(input, answer)
	return formatWordWithColor(input, colorArray)
}
