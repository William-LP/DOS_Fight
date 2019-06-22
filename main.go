package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/William-LP/go_training/pkg/classes"
	"github.com/William-LP/go_training/pkg/player"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewFigure("DOS Fight", "rectangles", true)
	myFigure.Print()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Player's Name: ")
	playerName, _ := reader.ReadString('\n')
	p := player.NewPlayer(playerName)
	log.Println(p.Gold)

	h := classes.CreateClass("Mage")

	log.Println("Health Point : ", h.GetHealthPoint())
	log.Println("Magic Point: ", h.GetMagicPoint())
	log.Println("Strength : ", h.GetStrength())
	log.Println("Intelect: ", h.GetIntelect())

}
