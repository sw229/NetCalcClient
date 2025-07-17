package main

import "embed"

//go:embed calc_icon.png
//go:embed settings_icon.png
var embedFS embed.FS

var (
	settings   Settings
	calculator Calc
)

func main() {
	settings = Settings{
		ServerIP:         "127.0.0.1",
		serverPort:       "8080",
		CalculateLocally: true,
	}
	initMainWindow()
}
