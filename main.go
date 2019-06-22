package main

import (
	"github.com/William-LP/go_training/pkg/classes"
	"log"
)


func main() {
	h := classe.CreateClass("Mage")
	log.Println("Point De Vie : ", h.GetPointDeVie())
	log.Println("Point De Magie: ", h.GetPointDeMagie())
	log.Println("Force : ", h.GetForce())
	log.Println("Intelligence: ", h.GetIntelligence())
}
