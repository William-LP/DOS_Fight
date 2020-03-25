package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/William-LP/DOS_Fight/pkg/classes"
	"github.com/William-LP/DOS_Fight/pkg/player"
	"github.com/common-nighthawk/go-figure"
	petname "github.com/dustinkirkland/golang-petname"
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

	hPlayer := classes.CreateClass(classe, p.BonusHealthPoint, p.BonusMagicPoint, p.BonusStrength, p.BonusIntelect)
	fmt.Println()
	fmt.Println("Rolling stats...")
	fmt.Println()
	pause(2)
	fmt.Println(p.Name + " is a " + classe + " with the following stats :")
	fmt.Println("Health Point : ", hPlayer.GetHealthPoint())
	fmt.Println("Magic Point: ", hPlayer.GetMagicPoint())
	fmt.Println("Strength : ", hPlayer.GetStrength())
	fmt.Println("Intelect: ", hPlayer.GetIntelect())
	pause(3)
	for i := 3; i > 0; i-- {
		clearScreen()
		Countdown := figure.NewFigure(strconv.Itoa(i), "small", true)
		Countdown.Print()
		pause(1)
	}
	clearScreen()
	Countdown := figure.NewFigure("Fight !", "small", true)
	Countdown.Print()
	botName := petname.Generate(1, "-")
	vsTitle := figure.NewFigure(p.Name+" vs "+strings.Title(botName), "cybermedium", true)
	vsTitle.Print()
	bHero := generateHeroBot(p)
	fightScreenHerosStats(bHero)
}

func fightScreenHerosStats(h classes.Hero) {
	hp := h.GetHealthPoint()
	mp := h.GetMagicPoint()
	s := h.GetStrength()
	i := h.GetIntelect()
	c := h.GetClasse()
	fmt.Println(hp, mp, s, i, c)

}

func generateHeroBot(p player.Player) classes.Hero {
	var bClasse string
	switch random(0, 100) % 2 {
	case 0:
		bClasse = "Warrior"
	default:
		bClasse = "Mage"
	}

	botBonus := p.Level * 10 // Bot bonues lines up with player level
	hpBonus := random(0, botBonus)
	mpBonus := random(0, botBonus)
	sBonus := random(0, botBonus)
	iBonus := random(0, botBonus)

	hBot := classes.CreateClass(bClasse, hpBonus, mpBonus, sBonus, iBonus)
	return hBot
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
