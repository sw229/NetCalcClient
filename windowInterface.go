package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func initMainWindow() {
	const (
		keyboardYOffset = 40
		buttonWidth     = 64
		buttonHeight    = 36
		buttonGap       = 2
		buttonXPos1     = 0
		buttonXPos2     = buttonWidth + buttonGap
		buttonXPos3     = buttonWidth*2 + buttonGap*2
		buttonXPos4     = buttonWidth*3 + buttonGap*3
		buttonYPos1     = keyboardYOffset
		buttonYPos2     = keyboardYOffset + buttonHeight + buttonGap
		buttonYPos3     = keyboardYOffset + buttonHeight*2 + buttonGap*2
		buttonYPos4     = keyboardYOffset + buttonHeight*3 + buttonGap*3
		buttonYPos5     = keyboardYOffset + buttonHeight*4 + buttonGap*4
	)
	//calcClient := newCalcClient("http", "127.0.0.1", "8080")
	calculator := Calc{}
	a := app.New()
	w := a.NewWindow("calc")
	w.Resize(fyne.NewSize(270, 320))
	typingField := widget.NewEntry()
	typingField.Resize(fyne.NewSize(262, 36))
	divButton := widget.NewButton("/", func() { // Finish this and put in a separate function
		calculator.Op = Div
		var err error
		if calculator.FirstOp {
			calculator.OperandX, err = strconv.ParseFloat(typingField.Text, 64)
			if err != nil {
				typingField.Text = ""
				return
			}
		}
	})
	divButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	divButton.Move(fyne.NewPos(buttonXPos2, buttonYPos1))

	mulButton := widget.NewButton("*", func() {})
	mulButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	mulButton.Move(fyne.NewPos(buttonXPos3, buttonYPos1))

	subButton := widget.NewButton("-", func() {})
	subButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	subButton.Move(fyne.NewPos(buttonXPos4, buttonYPos1))

	percentButton := widget.NewButton("%", func() {
		typingField.Text = typingField.Text + "%"
		typingField.Refresh()
	})
	percentButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	percentButton.Move(fyne.NewPos(buttonXPos1, buttonYPos1))

	sevenButton := widget.NewButton("7", func() {
		typingField.Text = typingField.Text + "7"
		typingField.Refresh()
	})
	sevenButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	sevenButton.Move(fyne.NewPos(buttonXPos1, buttonYPos2))

	eightButton := widget.NewButton("8", func() {
		typingField.Text = typingField.Text + "8"
		typingField.Refresh()
	})
	eightButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	eightButton.Move(fyne.NewPos(buttonXPos2, buttonYPos2))

	nineButton := widget.NewButton("9", func() {
		typingField.Text = typingField.Text + "9"
		typingField.Refresh()
	})
	nineButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	nineButton.Move(fyne.NewPos(buttonXPos3, buttonYPos2))

	fourButton := widget.NewButton("4", func() {
		typingField.Text = typingField.Text + "4"
		typingField.Refresh()
	})
	fourButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	fourButton.Move(fyne.NewPos(buttonXPos1, buttonYPos3))

	fiveButton := widget.NewButton("5", func() {
		typingField.Text = typingField.Text + "5"
		typingField.Refresh()
	})
	fiveButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	fiveButton.Move(fyne.NewPos(buttonXPos2, buttonYPos3))

	sixButton := widget.NewButton("6", func() {
		typingField.Text = typingField.Text + "6"
		typingField.Refresh()
	})
	sixButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	sixButton.Move(fyne.NewPos(buttonXPos3, buttonYPos3))

	oneButton := widget.NewButton("1", func() {
		typingField.Text = typingField.Text + "1"
		typingField.Refresh()
	})
	oneButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	oneButton.Move(fyne.NewPos(buttonXPos1, buttonYPos4))

	twoButton := widget.NewButton("2", func() {
		typingField.Text = typingField.Text + "2"
		typingField.Refresh()
	})
	twoButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	twoButton.Move(fyne.NewPos(buttonXPos2, buttonYPos4))

	threeButton := widget.NewButton("3", func() {
		typingField.Text = typingField.Text + "3"
		typingField.Refresh()
	})
	threeButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	threeButton.Move(fyne.NewPos(buttonXPos3, buttonYPos4))

	zeroButton := widget.NewButton("0", func() {
		typingField.Text = typingField.Text + "0"
		typingField.Refresh()
	})
	zeroButton.Resize(fyne.NewSize(2*buttonWidth+buttonGap, 36))
	zeroButton.Move(fyne.NewPos(buttonXPos1, buttonYPos5))

	decPointButton := widget.NewButton(".", func() {
		typingField.Text = typingField.Text + "."
		typingField.Refresh()
	})
	decPointButton.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	decPointButton.Move(fyne.NewPos(buttonXPos3, buttonYPos5))

	addButton := widget.NewButton("+", func() {})
	addButton.Resize(fyne.NewSize(buttonWidth, buttonHeight*2+buttonGap))
	addButton.Move(fyne.NewPos(buttonXPos4, buttonYPos2))

	exeButton := widget.NewButton("=", func() {})
	exeButton.Resize(fyne.NewSize(buttonWidth, buttonHeight*2+buttonGap))
	exeButton.Move(fyne.NewPos(buttonXPos4, buttonYPos4))

	content := container.NewWithoutLayout(
		typingField,
		sevenButton,
		eightButton,
		nineButton,
		addButton,
		fourButton,
		fiveButton,
		sixButton,
		subButton,
		oneButton,
		twoButton,
		threeButton,
		mulButton,
		zeroButton,
		decPointButton,
		exeButton,
		divButton,
		percentButton,
	)
	w.SetContent(content)
	w.ShowAndRun()
}
