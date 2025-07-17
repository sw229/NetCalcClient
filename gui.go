package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func initMainWindow() {
	const (
		windowWidth        = 270
		windowHeight       = 320
		buttonGap          = 2
		lastColumnGap      = 6
		keyboardYOffset    = 80
		actualWindowWidth  = windowWidth - 8
		actualWindowHeight = windowHeight - 7
		buttonWidth        = (actualWindowWidth - buttonGap*4) / 5
		buttonHeight       = (actualWindowHeight - keyboardYOffset - buttonGap*4) / 5
		buttonXPos1        = 0
		buttonXPos2        = buttonWidth + buttonGap
		buttonXPos3        = buttonWidth*2 + buttonGap*2
		buttonXPos4        = buttonWidth*3 + buttonGap*3
		buttonXPos5        = buttonWidth*4 + buttonGap*3 + lastColumnGap
		buttonYPos1        = keyboardYOffset
		buttonYPos2        = keyboardYOffset + buttonHeight + buttonGap
		buttonYPos3        = keyboardYOffset + buttonHeight*2 + buttonGap*2
		buttonYPos4        = keyboardYOffset + buttonHeight*3 + buttonGap*3
		buttonYPos5        = keyboardYOffset + buttonHeight*4 + buttonGap*4
	)

	calcIconData, _ := embedFS.ReadFile("calc_icon.png")
	calcIcon := fyne.NewStaticResource("calc_icon", calcIconData)

	standardSize := fyne.NewSize(buttonWidth, buttonHeight)
	doubleWideSize := fyne.NewSize(buttonWidth*2+buttonGap, buttonHeight)
	doubleHighSize := fyne.NewSize(buttonWidth, buttonHeight*2+buttonGap)

	a := app.New()
	mainWindow := a.NewWindow("calc")
	mainWindow.Resize(fyne.NewSize(windowWidth, windowHeight))
	mainWindow.SetIcon(calcIcon)

	// Label showing expression that is being typed
	expressionField := widget.NewLabel("")
	expressionFieldDefaultPos := fyne.NewPos(250, 0)
	expressionField.Move(expressionFieldDefaultPos)
	expressionField.TextStyle.Bold = true

	// Label showing result of the expression
	resultField := widget.NewLabel("")
	resultFieldDefaultPos := fyne.NewPos(250, 40)
	resultField.Move(resultFieldDefaultPos)
	resultField.TextStyle.Bold = true

	// Addition
	addButton := widget.NewButton("+", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = setOpCharacter(expressionField.Text, "+")
			expressionField.Refresh()
			setPosition(expressionField, expressionFieldDefaultPos)
		}
	})
	addButton.Resize(doubleHighSize)
	addButton.Move(fyne.NewPos(buttonXPos4, buttonYPos2))

	// Subtraction
	subButton := widget.NewButton("-", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = setOpCharacter(expressionField.Text, "-")
			expressionField.Refresh()
			setPosition(expressionField, expressionFieldDefaultPos)
		}
	})
	subButton.Resize(standardSize)
	subButton.Move(fyne.NewPos(buttonXPos4, buttonYPos1))

	// Multiplication
	mulButton := widget.NewButton("*", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = setOpCharacter(expressionField.Text, "*")
			expressionField.Refresh()
			setPosition(expressionField, expressionFieldDefaultPos)
		}
	})
	mulButton.Resize(standardSize)
	mulButton.Move(fyne.NewPos(buttonXPos3, buttonYPos1))

	// Division
	divButton := widget.NewButton("/", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = setOpCharacter(expressionField.Text, "/")
			expressionField.Refresh()
			setPosition(expressionField, expressionFieldDefaultPos)
		}
	})
	divButton.Resize(standardSize)
	divButton.Move(fyne.NewPos(buttonXPos2, buttonYPos1))

	// Button removes last character in expressionField
	delButton := widget.NewButton("del", func() {
		if len(expressionField.Text) > 0 {
			expressionField.Text = expressionField.Text[:len(expressionField.Text)-1]
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	delButton.Resize(standardSize)
	delButton.Move(fyne.NewPos(buttonXPos1, buttonYPos1))

	// Decimal point
	decPointButton := widget.NewButton(".", func() {
		if len(expressionField.Text) <= 29 && !hasDecPoint(expressionField.Text) {
			expressionField.Text = expressionField.Text + "."
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
		}
	})
	decPointButton.Resize(standardSize)
	decPointButton.Move(fyne.NewPos(buttonXPos3, buttonYPos5))

	// Parenthesis
	parenthesisLeftButton := widget.NewButton("(", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "("
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			calculate(expressionField, resultField, resultFieldDefaultPos)
		}
	})
	parenthesisLeftButton.Resize(standardSize)
	parenthesisLeftButton.Move(fyne.NewPos(buttonXPos5, buttonYPos3))

	parenthesisRightButton := widget.NewButton(")", func() {
		if len(expressionField.Text) <= 29 && canAddClosingParenthesis(expressionField.Text) {
			expressionField.Text = expressionField.Text + ")"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			calculate(expressionField, resultField, resultFieldDefaultPos)
		}
	})
	parenthesisRightButton.Resize(standardSize)
	parenthesisRightButton.Move(fyne.NewPos(buttonXPos5, buttonYPos4))

	// Switching + and -  DOES NOT WORK
	togglePosNegButton := widget.NewButton("+/-", func() {
		if expressionField.Text != "" {
			if expressionField.Text[0] == '-' {
				expressionField.Text = expressionField.Text[1:]
			} else {
				expressionField.Text = "-" + expressionField.Text
			}
			expressionField.Refresh()
		}
	})
	togglePosNegButton.Resize(standardSize)
	togglePosNegButton.Move(fyne.NewPos(buttonXPos5, buttonYPos5))

	// Button to execute operation
	exeButton := widget.NewButton("=", func() {
		err := calculate(expressionField, resultField, resultFieldDefaultPos)
		if err != nil {
			// Add error handling
		}
	})
	exeButton.Resize(doubleHighSize)
	exeButton.Move(fyne.NewPos(buttonXPos4, buttonYPos4))

	// Button to open options
	optionsButton := widget.NewButton("opt", func() {
		initOptionsWindow(a)
	})
	optionsButton.Resize(standardSize)
	optionsButton.Move(fyne.NewPos(buttonXPos5, buttonYPos1))

	// Clear input
	clearButton := widget.NewButton("C", func() {
		expressionField.Text = ""
		resultField.Text = ""
		setPosition(expressionField, expressionFieldDefaultPos)
		setPosition(resultField, resultFieldDefaultPos)
		expressionField.Refresh()
		resultField.Refresh()
	})
	clearButton.Resize(standardSize)
	clearButton.Move(fyne.NewPos(buttonXPos5, buttonYPos2))

	// Digits
	zeroButton := widget.NewButton("0", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "0"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	zeroButton.Resize(doubleWideSize)
	zeroButton.Move(fyne.NewPos(buttonXPos1, buttonYPos5))

	oneButton := widget.NewButton("1", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "1"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	oneButton.Resize(standardSize)
	oneButton.Move(fyne.NewPos(buttonXPos1, buttonYPos4))

	twoButton := widget.NewButton("2", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "2"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	twoButton.Resize(standardSize)
	twoButton.Move(fyne.NewPos(buttonXPos2, buttonYPos4))

	threeButton := widget.NewButton("3", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "3"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	threeButton.Resize(standardSize)
	threeButton.Move(fyne.NewPos(buttonXPos3, buttonYPos4))

	fourButton := widget.NewButton("4", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "4"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	fourButton.Resize(standardSize)
	fourButton.Move(fyne.NewPos(buttonXPos1, buttonYPos3))

	fiveButton := widget.NewButton("5", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "5"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	fiveButton.Resize(standardSize)
	fiveButton.Move(fyne.NewPos(buttonXPos2, buttonYPos3))

	sixButton := widget.NewButton("6", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "6"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	sixButton.Resize(standardSize)
	sixButton.Move(fyne.NewPos(buttonXPos3, buttonYPos3))

	sevenButton := widget.NewButton("7", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "7"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	sevenButton.Resize(standardSize)
	sevenButton.Move(fyne.NewPos(buttonXPos1, buttonYPos2))

	eightButton := widget.NewButton("8", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "8"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	eightButton.Resize(standardSize)
	eightButton.Move(fyne.NewPos(buttonXPos2, buttonYPos2))

	nineButton := widget.NewButton("9", func() {
		if len(expressionField.Text) <= 29 {
			expressionField.Text = expressionField.Text + "9"
			setPosition(expressionField, expressionFieldDefaultPos)
			expressionField.Refresh()
			err := calculate(expressionField, resultField, resultFieldDefaultPos)
			if err != nil {
				// Add error handling
			}
		}
	})
	nineButton.Resize(standardSize)
	nineButton.Move(fyne.NewPos(buttonXPos3, buttonYPos2))

	content := container.NewWithoutLayout(
		expressionField,
		resultField,
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
		clearButton,
		optionsButton,
		parenthesisLeftButton,
		parenthesisRightButton,
		togglePosNegButton,
		divButton,
		delButton,
	)
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

// Window containing connection options (server IP, port, username, etc.)
func initOptionsWindow(a fyne.App) {
	settingsIconData, _ := embedFS.ReadFile("settings_icon.png")
	settingsIcon := fyne.NewStaticResource("settings_icon", settingsIconData)

	optionsWindow := a.NewWindow("Options")
	optionsWindow.Resize(fyne.NewSize(218, 245))
	optionsWindow.SetIcon(settingsIcon)

	// Entry for server IP
	serverAddressLabel := widget.NewLabel("Server address and port:")
	serverAddressLabel.Move(fyne.NewPos(-5, -10))
	semicololLabel := widget.NewLabel(":")
	semicololLabel.Move(fyne.NewPos(145, 19))
	serverAddressField := widget.NewEntry()
	serverAddressField.Resize(fyne.NewSize(150, 36))
	serverAddressField.Move(fyne.NewPos(0, 20))
	serverAddressField.Text = settings.ServerIP
	serverAddressField.Refresh()
	// Entry for server port
	serverPortField := widget.NewEntry()
	serverPortField.Resize(fyne.NewSize(50, 36))
	serverPortField.Move(fyne.NewPos(160, 20))
	serverPortField.Text = settings.serverPort
	serverPortField.Refresh()
	// Username field
	usernameLabel := widget.NewLabel("Username:")
	usernameLabel.Move(fyne.NewPos(-5, 50))
	usernameField := widget.NewEntry()
	usernameField.Resize(fyne.NewSize(210, 36))
	usernameField.Move(fyne.NewPos(0, 80))
	usernameField.Text = settings.Username
	// Password field
	passwordLabel := widget.NewLabel("Password:")
	passwordLabel.Move(fyne.NewPos(-5, 110))
	passwordField := widget.NewEntry()
	passwordField.Resize(fyne.NewSize(210, 36))
	passwordField.Move(fyne.NewPos(0, 140))
	passwordField.Text = settings.Password
	// Checkbox for setting wether client does the calculations locally or uses a server
	localCalcCheckbox := widget.NewCheck("Calculate locally", func(checked bool) {
		settings.CalculateLocally = checked
	})
	localCalcCheckbox.Checked = settings.CalculateLocally
	localCalcCheckbox.Resize(fyne.NewSize(36, 36))
	localCalcCheckbox.Move(fyne.NewPos(-5, 170))
	// Button to apply settings
	applyButton := widget.NewButton("ok", func() {
		settings.ServerIP = serverAddressField.Text
		settings.serverPort = serverPortField.Text
		settings.Username = usernameField.Text
		settings.Password = passwordField.Text
		optionsWindow.Close()
	})
	applyButton.Resize(fyne.NewSize(50, 36))
	applyButton.Move(fyne.NewPos(105, 200))
	// Button to cancel settings
	cancelButton := widget.NewButton("cancel", func() {
		optionsWindow.Close()
	})
	cancelButton.Resize(fyne.NewSize(50, 36))
	cancelButton.Move(fyne.NewPos(160, 200))

	content := container.NewWithoutLayout(
		serverAddressLabel,
		semicololLabel,
		serverAddressField,
		serverPortField,
		usernameLabel,
		usernameField,
		passwordLabel,
		passwordField,
		localCalcCheckbox,
		applyButton,
		cancelButton,
	)
	optionsWindow.SetContent(content)
	optionsWindow.Show()
}
