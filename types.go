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
type Calc struct {
	Expression string
	Result     string
	CurrentOp  int
}
