package main

const MaxGuesses = 6
const MaxWordLength = 5

type ANSIColor int

const ANSIColorReset = "\033[0m"
const ANSIBold = "\u001b[1m"

const (
	ANSIColorUnknown ANSIColor = iota
	ANSIColorExactMatch
	ANSIColorExistMatch
	ANSIColorNoMatch
)
