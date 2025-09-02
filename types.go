package main

// Enum containing types of operations
const (
	Add = iota + 1
	Sub
	Mul
	Div
	Pow
	Root
)

// Struct containing calculator settings
type Settings struct {
	ServerProtocol   string
	ServerIP         string
	serverPort       string
	Username         string
	Password         string
	CalculateLocally bool
}

// Struct containing info about current operation
// Calculated field is set to true when exeButton (=) is clicked
// If Calculated is true the next typed number overwrites expressionField.Text
type CalcState struct {
	Calculated bool
	Expression string
	Result     string
	CurrentOp  int
}
