package main

import (
	"fmt"
)

// package classes

// Hero ; interface de base
type Hero interface {
	sayClass() string
}

// Carac ; structure étendu aux structures des différentes classes
type Carac struct {
	PointDeVie,
	PointDeMagie,
	Force,
	Intelligence int
}

// Guerrier ; hérite des Carac + possede un bonus de Point de vie et de Force
type Guerrier struct {
	Carac
}

// Mage ; hérite des Carac + possede un bonus de Point de magie et d'Intélligence
type Mage struct {
	Carac
}

func (g Guerrier) sayClass() string {
	return "jsuis un war"
}

func (m Mage) sayClass() string {
	return "jsuis un mago"
}

func main() {
	c := Carac{10, 10, 10, 10}
	h := Guerrier{c}
	fmt.Printf(h.sayClass())
}
