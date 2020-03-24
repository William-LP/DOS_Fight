package main

import (
	"fmt"

	"github.com/William-LP/go_training/pkg/player"
	"github.com/common-nighthawk/go-figure"
)

func playerStats(p player.Player) {
	clearScreen()
	ShopFigure := figure.NewFigure("Player's Stats", "small", true)
	ShopFigure.Print()
	fmt.Println()
	fmt.Println("Player Name :", p.Name)
	fmt.Println("Gold :", p.Gold)
	fmt.Println("Bonus HP :", p.BonusHealthPoint)
	fmt.Println("Bonus MP :", p.BonusMagicPoint)
	fmt.Println("Bonus Strenght :", p.BonusStrength)
	fmt.Println("Bonus Intelect :", p.BonusIntelect)
	fmt.Println()
	fmt.Println("Press 'Enter' to continue")
	fmt.Scanln()
	mainMenu(p)
}