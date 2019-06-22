package main

import (
	"log"

	"github.com/William-LP/go_training/pkg/classes"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()

	h := classes.CreateClass("Mage")

	log.Println("Health Point : ", h.GetHealthPoint())
	log.Println("Magic Point: ", h.GetMagicPoint())
	log.Println("Strength : ", h.GetStrength())
	log.Println("Intelect: ", h.GetIntelect())
}
