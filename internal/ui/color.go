package ui

type ANSIColor int

const (
	ANSIColorReset = "\033[0m"
	ANSIBold       = "\u001b[1m"
)

const (
	ANSIColorUnknown ANSIColor = iota
	ANSIColorExactMatch
	ANSIColorExistMatch
	ANSIColorNoMatch
)

func (c ANSIColor) String() string {
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
		return ANSIColorReset
	}
}
