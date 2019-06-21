package main

import (
	"log"

	"github.com/William-LP/go_training/pkg/mario"
)

func main() {
	h := mario.CreateHero(10, 10, 10, 10, 10)
	log.Println(h)
}
