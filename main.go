package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/William-LP/go_training/pkg/player"

	"github.com/common-nighthawk/go-figure"
)

func getInput(textToPrint string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(textToPrint)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\r\n")

	return userInput
}

func mainMenu(p player.Player) {
	fmt.Println()
	fmt.Println("Start a game (1)")
	fmt.Println("Open the shop (2)")
	fmt.Println("Exit the game (3)")
	fmt.Println()
	switch getInput("Choice : ") {
	case "1":
		startGame()
	case "2":
		openShop(p)
	case "3":
		exitGame()
	default:
		fmt.Println("Error, please make a correct choice")
		mainMenu(p)
	}
}

func exitGame() {
	fmt.Println("Thanks for playing, adios !")
	os.Exit(0)
}

func startGame() {
	fmt.Println("Lets start a game !")
}

func openShop(p player.Player) {
	fmt.Println()
	myFigure := figure.NewFigure("Shop", "rectangles", true)
	myFigure.Print()
	fmt.Println()
	fmt.Println("Welcome to the shop", p.Name, "!")
	fmt.Println("You've got", p.Gold, "gold(s) available.")

}

func main() {
	myFigure := figure.NewFigure("DOS Fight", "rectangles", true)
	myFigure.Print()

	playerName := getInput("Player's Name: ")

	p := player.NewPlayer(playerName)

	mainMenu(p)
	/*
		h := classes.CreateClass("Mage")

		fmt.Println("Health Point : ", h.GetHealthPoint())
		fmt.Println("Magic Point: ", h.GetMagicPoint())
		fmt.Println("Strength : ", h.GetStrength())
		fmt.Println("Intelect: ", h.GetIntelect())
	*/
}
