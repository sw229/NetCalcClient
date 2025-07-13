package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// "Enum" containing types of operations
const (
	Add = iota + 1
	Sub
	Mul
	Div
	Pow
	Root
)

// Struct containing info about server
type CalcClient struct {
	Protocol      string
	ServerAddress string
	ServerPort    string
}

// Struct containing info about current operation
type Calc struct {
	Op         int
	OperandX   float64
	OperandY   float64
	LastResult float64
	FirstOp    bool
}

// Constructor for CalcClient struct
func newCalcClient(protocol, serverAddress, serverPort string) CalcClient {
	return CalcClient{protocol, serverAddress, serverPort}
}

// Function sends request to calculator server
// Function requires the operands and operation type listed in the enum
func (c CalcClient) operation(op int, x, y float64) (float64, error) {
	var opStr string
	switch op {
	case 1:
		opStr = "add"
	case 2:
		opStr = "sub"
	case 3:
		opStr = "mul"
	case 4:
		opStr = "div"
	case 5:
		opStr = "pow"
	case 6:
		opStr = "root"
	default:
		return 0, errors.New("Invalid operation")
	}
	resp, err := http.Get(fmt.Sprintf("%s://%s:%s/%s?x=%g&y=%g", c.Protocol, c.ServerAddress, c.ServerPort, opStr, x, y))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseFloat(string(body), 64)
	if err != nil {
		return 0, errors.New(string(body))
	}
	return result, nil
}
