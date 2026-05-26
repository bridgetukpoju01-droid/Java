package main

import (
	"strings"
)

var asciiDigits = map[rune][5]string{
	'1': {
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
	},
	'2': {
		" ___ ",
		"    |",
		" ___|",
		"|    ",
		"|___ ",
	},
}

func StringToArt(input string) string {
	if input == "" {
		return ""

	}

	var digit strings.Builder

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		for _, r := range line {
			if _, exists := asciiDigits[r]; !exists {
				return ""
			}
		}
		for row := 0; row < 5; row++ {
			for _, r := range line {
				digit.WriteString(asciiDigits[r][row])
			}
			digit.WriteString("\n")
		}
	}

	return digit.String()
}
