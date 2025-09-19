package main

import "embed"

//go:embed calc_icon.png
//go:embed settings_icon.png
var embedFS embed.FS

func main() {
	var (
		settings        Settings
		calculatorState CalcState
	)

	settings = Settings{
		ServerIP:         "127.0.0.1",
		serverPort:       "8080",
		ServerProtocol:   "http",
		CalculateLocally: true,
	}
	initMainWindow(&settings, &calculatorState)
}
