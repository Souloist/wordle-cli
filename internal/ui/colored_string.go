package ui

import (
	"fmt"
	"strings"
)

type ColoredString string

func (s ColoredString) FormatWithColorArray(colors []ANSIColor) (string, error) {
	if len(s) != len(colors) {
		return "", fmt.Errorf("string and color array have different lengths: %d vs %d", len(s), len(colors))
	}

	var formattedChars []string
	for i, char := range s {
		color := colors[i].String()
		formattedChar := ANSIBold + color + string(char) + ANSIColorReset
		formattedChars = append(formattedChars, formattedChar)
	}

	return strings.Join(formattedChars, ""), nil
}
