package main

// package classes

import (
	"log"
)

// Hero ; interface de base
type Hero interface {
	GetPointDeVie() int
	GetPointDeMagie() int
	GetForce() int
	GetIntelligence() int
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
	BonusPointDeVie int
	BonusForce      int
}

// Mage ; hérite des Carac + possede un bonus de Point de magie et d'Intélligence
type Mage struct {
	Carac
	BonusPointDeMagie int
	BonusIntelligence int
}

// CreateHero ;
func CreateHero(classe string, pdv, pdm, f, i int) Hero {
	carac := Carac{
		PointDeVie:   pdv,
		PointDeMagie: pdm,
		Force:        f,
		Intelligence: i,
	}
	switch classe {
	case "Guerrier":
		h := Guerrier{
			Carac:           carac,
			BonusPointDeVie: 150,
			BonusForce:      100,
		}
		h.PointDeVie = (h.BonusPointDeVie + h.PointDeVie)
		return h
	case "Mage":
		h := Mage{
			Carac:             carac,
			BonusPointDeMagie: 150,
			BonusIntelligence: 100,
		}
		return h
	}
	return nil
}

func main() {
	monHero := CreateHero("Guerrier", 10, 10, 10, 10)
	log.Println(monHero.GetPointDeVie)
}
