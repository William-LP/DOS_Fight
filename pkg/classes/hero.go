package classe

import (
	"math/rand"
	"time"
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

	// Source d'aleatoire
	rand.Seed(time.Now().UnixNano())

	pdv := 100 // points de vie
	pdm := 100 // points de magie
	f := 20    // force
	i := 20    // intelligence

	// valeur max des points bonus ajouté aux stats
	CapBonus := 50
	// pourcentage maximum dont peut etre augmenter le CapBonus
	MaxPourcentageBonus := 30

	BonusEffectif := CapBonus + CapBonus*rand.Intn(MaxPourcentageBonus)/100

	BonusPoint := (rand.Intn(BonusEffectif))
	BonusStat := BonusEffectif - BonusPoint

	switch class {
	case "Guerrier":
		pdv = pdv + BonusPoint
		f = f + BonusStat
		c := Carac{
			PointDeVie:   pdv,
			PointDeMagie: pdm,
			Force:        f,
			Intelligence: i,
		}
		h := Guerrier{c}
		return h

	case "Mage":
		pdm = pdm + BonusPoint
		i = i + BonusStat
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
