package main

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
