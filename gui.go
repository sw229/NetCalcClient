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
	standardSize := fyne.NewSize(buttonWidth, buttonHeight)
	doubleWideSize := fyne.NewSize(buttonWidth*2+buttonGap, buttonHeight)
	doubleHighSize := fyne.NewSize(buttonWidth, buttonHeight*2+buttonGap)
	//calcClient := newCalcClient("http", "127.0.0.1", "8080")
	a := app.New()
	mainWindow := a.NewWindow("calc")
	mainWindow.Resize(fyne.NewSize(windowWidth, windowHeight))
	typingField := widget.NewEntry()
	typingField.Resize(fyne.NewSize(actualWindowWidth, 36))
	divButton := widget.NewButton("/", func() { // Finish this and put in a separate function

	})

	// Addition
	addButton := widget.NewButton("+", func() {})
	addButton.Resize(doubleHighSize)
	addButton.Move(fyne.NewPos(buttonXPos4, buttonYPos2))

	// Subtraction
	subButton := widget.NewButton("-", func() {})
	subButton.Resize(standardSize)
	subButton.Move(fyne.NewPos(buttonXPos4, buttonYPos1))

	// Multiplication
	mulButton := widget.NewButton("*", func() {})
	mulButton.Resize(standardSize)
	mulButton.Move(fyne.NewPos(buttonXPos3, buttonYPos1))

	// Division
	divButton.Resize(standardSize)
	divButton.Move(fyne.NewPos(buttonXPos2, buttonYPos1))

	// Percentage
	percentButton := widget.NewButton("%", func() {
		typingField.Text = typingField.Text + "%"
		typingField.Refresh()
	})
	percentButton.Resize(standardSize)
	percentButton.Move(fyne.NewPos(buttonXPos1, buttonYPos1))

	// Decimal point
	decPointButton := widget.NewButton(".", func() {
		typingField.Text = typingField.Text + "."
		typingField.Refresh()
	})
	decPointButton.Resize(standardSize)
	decPointButton.Move(fyne.NewPos(buttonXPos3, buttonYPos5))

	// Parenthesis
	parenthesisLeftButton := widget.NewButton("(", func() {
		// Add something here
	})
	parenthesisLeftButton.Resize(standardSize)
	parenthesisLeftButton.Move(fyne.NewPos(buttonXPos5, buttonYPos3))

	parenthesisRightButton := widget.NewButton(")", func() {
		// Add something here
	})
	parenthesisRightButton.Resize(standardSize)
	parenthesisRightButton.Move(fyne.NewPos(buttonXPos5, buttonYPos4))

	// Switching + and -
	togglePosNegButton := widget.NewButton("+/-", func() {
		if typingField.Text != "" {
			if typingField.Text[0] == '-' {
				typingField.Text = typingField.Text[1:]
			} else {
				typingField.Text = "-" + typingField.Text
			}
			typingField.Refresh()
		}
	})
	togglePosNegButton.Resize(standardSize)
	togglePosNegButton.Move(fyne.NewPos(buttonXPos5, buttonYPos5))

	// Button to execute operation
	exeButton := widget.NewButton("=", func() {
		initOptionsWindow(a)
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
		typingField.Text = ""
		typingField.Refresh()
	})
	clearButton.Resize(standardSize)
	clearButton.Move(fyne.NewPos(buttonXPos5, buttonYPos2))

	// Digits
	zeroButton := widget.NewButton("0", func() {
		typingField.Text = typingField.Text + "0"
		typingField.Refresh()
	})
	zeroButton.Resize(doubleWideSize)
	zeroButton.Move(fyne.NewPos(buttonXPos1, buttonYPos5))

	oneButton := widget.NewButton("1", func() {
		typingField.Text = typingField.Text + "1"
		typingField.Refresh()
	})
	oneButton.Resize(standardSize)
	oneButton.Move(fyne.NewPos(buttonXPos1, buttonYPos4))

	twoButton := widget.NewButton("2", func() {
		typingField.Text = typingField.Text + "2"
		typingField.Refresh()
	})
	twoButton.Resize(standardSize)
	twoButton.Move(fyne.NewPos(buttonXPos2, buttonYPos4))

	threeButton := widget.NewButton("3", func() {
		typingField.Text = typingField.Text + "3"
		typingField.Refresh()
	})
	threeButton.Resize(standardSize)
	threeButton.Move(fyne.NewPos(buttonXPos3, buttonYPos4))

	fourButton := widget.NewButton("4", func() {
		typingField.Text = typingField.Text + "4"
		typingField.Refresh()
	})
	fourButton.Resize(standardSize)
	fourButton.Move(fyne.NewPos(buttonXPos1, buttonYPos3))

	fiveButton := widget.NewButton("5", func() {
		typingField.Text = typingField.Text + "5"
		typingField.Refresh()
	})
	fiveButton.Resize(standardSize)
	fiveButton.Move(fyne.NewPos(buttonXPos2, buttonYPos3))

	sixButton := widget.NewButton("6", func() {
		typingField.Text = typingField.Text + "6"
		typingField.Refresh()
	})
	sixButton.Resize(standardSize)
	sixButton.Move(fyne.NewPos(buttonXPos3, buttonYPos3))

	sevenButton := widget.NewButton("7", func() {
		typingField.Text = typingField.Text + "7"
		typingField.Refresh()
	})
	sevenButton.Resize(standardSize)
	sevenButton.Move(fyne.NewPos(buttonXPos1, buttonYPos2))

	eightButton := widget.NewButton("8", func() {
		typingField.Text = typingField.Text + "8"
		typingField.Refresh()
	})
	eightButton.Resize(standardSize)
	eightButton.Move(fyne.NewPos(buttonXPos2, buttonYPos2))

	nineButton := widget.NewButton("9", func() {
		typingField.Text = typingField.Text + "9"
		typingField.Refresh()
	})
	nineButton.Resize(standardSize)
	nineButton.Move(fyne.NewPos(buttonXPos3, buttonYPos2))

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
		clearButton,
		optionsButton,
		parenthesisLeftButton,
		parenthesisRightButton,
		togglePosNegButton,
		divButton,
		percentButton,
	)
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

// Window containing connection options (server IP, port, username, etc.)
func initOptionsWindow(a fyne.App) {
	optionsWindow := a.NewWindow("Network options")
	optionsWindow.Resize(fyne.NewSize(640, 480))
	// Entry for server IP
	serverAddressField := widget.NewEntry()
	serverAddressField.Resize(fyne.NewSize(50, 36))
	serverAddressField.Text = settings.ServerIP
	// Entry for server port
	serverPortField := widget.NewEntry()
	serverPortField.Resize(fyne.NewSize(50, 36))
	serverPortField.Text = settings.serverPort
	// Username field
	usernameField := widget.NewEntry()
	usernameField.Resize(fyne.NewSize(50, 36))
	usernameField.Text = settings.Username
	// Password field
	passwordField := widget.NewEntry()
	passwordField.Resize(fyne.NewSize(50, 36))
	passwordField.Text = settings.Password
	// Checkbox for setting wether client does the calculations locally or uses a server
	localCalcCheckbox := widget.NewCheck("Calculate locally", func(checked bool) {})
	localCalcCheckbox.Checked = settings.CalculateLocally
	// Button to apply settings
	applyButton := widget.NewButton("ok", func() {
		settings.ServerIP = serverAddressField.Text
		settings.serverPort = serverPortField.Text
		settings.Username = usernameField.Text
		settings.Password = passwordField.Text
		optionsWindow.Close()
	})
	applyButton.Resize(fyne.NewSize(50, 36))
	// Button to cancel settings
	cancelButton := widget.NewButton("cancel", func() {
		optionsWindow.Close()
	})
	cancelButton.Resize(fyne.NewSize(50, 36))

	content := container.NewWithoutLayout(
		serverAddressField,
		serverPortField,
		usernameField,
		passwordField,
		localCalcCheckbox,
		applyButton,
		cancelButton,
	)
	optionsWindow.SetContent(content)
	optionsWindow.Show()
}
