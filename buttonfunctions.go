package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Function called when operation button is clicked
func calcOpButtonFunc(calculatorState CalcState, buttonLabel string, expressionField *widget.Label, expressionFieldDefaultPos fyne.Position) {
	if len(expressionField.Text) <= 29 {
		expressionField.Text = setOpCharacter(expressionField.Text, buttonLabel)
		expressionField.Refresh()
		setPosition(expressionField, expressionFieldDefaultPos)
		calculatorState.Calculated = false
	}
}

// Function called when clearButton is clicked
func calcClearButtonFunc(calculatorState *CalcState, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	expressionField.Text = ""
	resultField.Text = ""
	setPosition(expressionField, expressionFieldDefaultPos)
	setPosition(resultField, resultFieldDefaultPos)
	expressionField.Refresh()
	resultField.Refresh()
	calculatorState.Calculated = false
}

// Function called when a number button is clicked
func calcNumButtonFunc(settings Settings, calculatorState *CalcState, buttonLabel string, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	if calculatorState.Calculated {
		expressionField.Text = buttonLabel
		setPosition(expressionField, expressionFieldDefaultPos)
		expressionField.Refresh()
		calculate(settings, expressionField, resultField, resultFieldDefaultPos)
		calculatorState.Calculated = false
		return
	}

	if len(expressionField.Text) <= 29 {
		expressionField.Text = expressionField.Text + buttonLabel
		setPosition(expressionField, expressionFieldDefaultPos)
		expressionField.Refresh()
		calculate(settings, expressionField, resultField, resultFieldDefaultPos)
	}
}

// Function called when decimal point button clicked
func calcDecPointButtonFunc(settings Settings, calculatorState *CalcState, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	if len(expressionField.Text) <= 29 && !hasDecPoint(expressionField.Text) {
		expressionField.Text = expressionField.Text + "."
		setPosition(expressionField, expressionFieldDefaultPos)
		expressionField.Refresh()
		calculate(settings, expressionField, resultField, resultFieldDefaultPos)
		calculatorState.Calculated = false
	}
}

// Function called when opening parenthesis button is clicked
func calcParLeftButtonFunc(settings Settings, calculatorState *CalcState, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	if len(expressionField.Text) <= 29 {
		expressionField.Text = expressionField.Text + "("
		setPosition(expressionField, expressionFieldDefaultPos)
		expressionField.Refresh()
		calculate(settings, expressionField, resultField, resultFieldDefaultPos)
		calculatorState.Calculated = false
	}
}

// Function called when closing parenthesis button is clicked
func calcParRightButtonFunc(settings Settings, calculatorState *CalcState, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	if len(expressionField.Text) <= 29 && canAddClosingParenthesis(expressionField.Text) {
		expressionField.Text = expressionField.Text + ")"
		setPosition(expressionField, expressionFieldDefaultPos)
		expressionField.Refresh()
		calculate(settings, expressionField, resultField, resultFieldDefaultPos)
		calculatorState.Calculated = false
	}
}

// Function called when exeButton is clicked
func calcExeButtonFunc(settings Settings, calculatorState *CalcState, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	calculate(settings, expressionField, resultField, resultFieldDefaultPos)
	expressionField.Text = resultField.Text
	setPosition(expressionField, expressionFieldDefaultPos)
	expressionField.Refresh()
	resultField.Text = ""
	setPosition(resultField, resultFieldDefaultPos)
	resultField.Refresh()
	calculatorState.Calculated = true
}

// Function called when TogglePosNegButton is clicked
func calcTogPosNegButtonFunc(settings Settings, calculatorState *CalcState, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	calcExeButtonFunc(settings, calculatorState, expressionField, resultField, expressionFieldDefaultPos, resultFieldDefaultPos)
	if expressionField.Text != "" {
		if expressionField.Text[0] != '-' {
			expressionField.Text = "-" + expressionField.Text
		} else {
			expressionField.Text = expressionField.Text[1:]
		}
	}
	setPosition(expressionField, expressionFieldDefaultPos)
	expressionField.Refresh()
	setPosition(resultField, resultFieldDefaultPos)
	resultField.Refresh()
}

// Function for the del button
func calcDelButtonFunc(settings Settings, calculatorState *CalcState, expressionField, resultField *widget.Label, expressionFieldDefaultPos, resultFieldDefaultPos fyne.Position) {
	if len(expressionField.Text) > 0 {
		expressionField.Text = expressionField.Text[:len(expressionField.Text)-1]
		setPosition(expressionField, expressionFieldDefaultPos)
		expressionField.Refresh()
		calculate(settings, expressionField, resultField, resultFieldDefaultPos)
	}
	calculatorState.Calculated = false
}
