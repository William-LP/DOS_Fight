package main

import (
	"fmt"

	"github.com/William-LP/DOS_Fight/pkg/classes"
	"github.com/William-LP/DOS_Fight/pkg/player"
	"github.com/common-nighthawk/go-figure"
)

func game(p player.Player) {
	clearScreen()
	var classe string
	startGameFigure := figure.NewFigure("Lets fight !", "cybermedium", true)
	startGameFigure.Print()
	fmt.Println()
	fmt.Printf("Your player's bonus stats are +%d HP, +%d MP, +%d Strenght and +%d Intelect.", p.BonusHealthPoint, p.BonusMagicPoint, p.BonusStrength, p.BonusIntelect)
	fmt.Println()
	fmt.Println()
	fmt.Println("(1) Warrior | (2) Mage")
	fmt.Println()
	switch getInput("Pick a class (press 'c' ton cancel) : ") {
	case "1":
		classe = "Warrior"
	case "2":
		classe = "Mage"
	case "c":
		mainMenu(p)
	default:
		fmt.Println("Error, please make a correct choice")
		game(p)
	}

	h := classes.CreateClass(classe, p.BonusHealthPoint, p.BonusMagicPoint, p.BonusStrength, p.BonusIntelect)
	fmt.Println()
	fmt.Println("Rolling stats...")
	fmt.Println()
	pause(2)
	fmt.Println(p.Name + " is a " + classe + " with the following stats :")
	fmt.Println("Health Point : ", h.GetHealthPoint())
	fmt.Println("Magic Point: ", h.GetMagicPoint())
	fmt.Println("Strength : ", h.GetStrength())
	fmt.Println("Intelect: ", h.GetIntelect())

}
