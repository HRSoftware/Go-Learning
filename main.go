package main

import (
	"Learning/games"
	"Learning/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		outputOptions()

		num, _ := utils.GetUserInputNumber()
		switch num {
		case 1:
			games.Play()
		case 2:
			utils.StartWebServer()
		case 3:
			utils.StopWebServer()
		case 4:
			utils.OpenWindow()
		case 5:
			fmt.Println("Exiting...")
			return
		}

		//select {
		//case <-stopChan:
		//	utils.StopWebServer()
		//	return
		//}
	}
}

func outputOptions() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Guessing game")
	fmt.Println("2. Starting Server")
	fmt.Println("3. Stop server")
	fmt.Println("4. Open a window")
	fmt.Println("5. Quit")
}

func webserver() {

}

func setName(name string) {}
func expriment() {

}
