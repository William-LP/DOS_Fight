package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/William-LP/DOS_Fight/pkg/player"
	"github.com/common-nighthawk/go-figure"
)

func getInput(textToPrint string) string {
	println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(textToPrint)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\r\n")

	return userInput
}

func mainMenu(p player.Player) {
	clearScreen()
	myFigure := figure.NewFigure("DOS Fight", "rectangles", true)
	myFigure.Print()

	fmt.Println("Start a game (1)")
	fmt.Println("Open the shop (2)")
	fmt.Println("See my player's stats (3)")
	fmt.Println("Exit the game (4)")
	switch getInput("Choice : ") {
	case "1":
		game(p)
	case "2":
		shop(p)
	case "3":
		playerStats(p)
	case "4":
		exitGame()
	default:
		fmt.Println("Error, please make a correct choice")
		mainMenu(p)
	}
}

func exitGame() {
	fmt.Println("Thanks for playing, adios !")
	pause(2)
	os.Exit(0)
}

func main() {
	clearScreen()
	myFigure := figure.NewFigure("DOS Fight", "rectangles", true)
	myFigure.Print()

	playerName := getInput("Player's Name: ")

	p := player.NewPlayer(playerName)

	mainMenu(p)
}
