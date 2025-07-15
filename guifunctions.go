package main

import (
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Function to move a label left as length of its content changes
// Function returns false if label goes over the boundary of the window
func setPosition(l *widget.Label, defaultPos fyne.Position) bool {
	newXCoord := defaultPos.X - float32(len(l.Text))*8
	if newXCoord < -5 {
		return false
	}
	l.Move(fyne.NewPos(newXCoord, defaultPos.Y))
	return true
}

// Function adds a character representing an operation to the string
// If there is such character in the end of the string,
// it is replaced with the one passed to the function
func setOpCharacter(exp string, char string) string {
	if exp == "" {
		return exp
	}
	if unicode.IsDigit(rune(exp[len(exp)-1])) || exp[len(exp)-1] == '%' || exp[len(exp)-1] == ')' || (exp[len(exp)-1] == '(' && (char == "+" || char == "-")) {
		exp += char
	} else {
		exp = exp[:len(exp)-1] + char
	}
	return exp
}

// Function checks if a string contains a dot
func hasDecPoint(exp string) bool {
	for _, char := range exp {
		if char == '.' {
			return true
		}
	}
	return false
}

func canAddClosingParenthesis(exp string) bool {
	var count int
	for _, char := range exp {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
	}
	if count > 0 {
		return true
	}
	return false
}
