package main

import (
	"log"
	"math/rand"
	"time"
)

// package classes

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
}

// Mage ; hérite des Carac + possede un bonus de Point de magie et d'Intélligence
type Mage struct {
	Carac
}

// -------------------  Getters & Setters

// GetPointDeVie ;
func (g Guerrier) GetPointDeVie() int {
	return g.PointDeVie
}

// GetPointDeVie ;
func (m Mage) GetPointDeVie() int {
	return m.PointDeVie
}

// GetPointDeMagie ;
func (g Guerrier) GetPointDeMagie() int {
	return g.PointDeMagie
}

// GetPointDeMagie ;
func (m Mage) GetPointDeMagie() int {
	return m.PointDeMagie
}

// GetForce ;
func (g Guerrier) GetForce() int {
	return g.Force
}

// GetForce ;
func (m Mage) GetForce() int {
	return m.Force
}

// GetIntelligence ;
func (g Guerrier) GetIntelligence() int {
	return g.Intelligence
}

// GetIntelligence ;
func (m Mage) GetIntelligence() int {
	return m.Intelligence
}

// ------------------- Fin de Getters & Setters

// CreateClass ;
func CreateClass(class string) Hero {

	pdv := 100
	pdm := 100
	f := 20
	i := 20

	rand.Seed(time.Now().UnixNano())
	amplifierPts := (rand.Intn(30-10) + 10)
	amplifierStat := (rand.Intn(30-10) + 10)
	log.Println(amplifierStat)

	switch class {
	case "Guerrier":
		pdv = pdv + (pdv * amplifierPts / 100)
		f = f + (f * amplifierStat / 100)
		c := Carac{
			PointDeVie:   pdv,
			PointDeMagie: pdm,
			Force:        f,
			Intelligence: i,
		}
		h := Guerrier{c}
		return h

	case "Mage":
		pdm = pdm + (pdv * amplifierPts / 100)
		i = i + (f * amplifierStat / 100)
		c := Carac{
			PointDeVie:   pdv,
			PointDeMagie: pdm,
			Force:        f,
			Intelligence: i,
		}
		h := Mage{c}
		return h
	}
	return nil
}

func main() {
	h := CreateClass("Mage")
	log.Println("Point De Vie : ", h.GetPointDeVie())
	log.Println("Point De Magie: ", h.GetPointDeMagie())
	log.Println("Force : ", h.GetForce())
	log.Println("Intelligence: ", h.GetIntelligence())
}
