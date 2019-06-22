package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	classe "github.com/William-LP/go_training/pkg/classes"
)

func main() {
	h := classe.CreateClass("Mage")
	log.Println("Health Point : ", h.GetHealthPoint())
	log.Println("Magic Point: ", h.GetMagicPoint())
	log.Println("Strength : ", h.GetStrength())
	log.Println("Intelect: ", h.GetIntelect())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Player's Name: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
