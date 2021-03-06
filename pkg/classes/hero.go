package classes

import (
	"math/rand"
	"time"
)

// Hero ; interface de base
type Hero interface {
	GetHealthPoint() int
	GetMagicPoint() int
	GetStrength() int
	GetIntelect() int
	GetClasse() string
}

// Trait ; structure étendu aux structures des différentes classes
type Trait struct {
	HealthPoint,
	MagicPoint,
	Strength,
	Intelect int
	Classe string
}

// Warrior ; hérite des Trait + possede un bonus de Point de vie et de Strength
type Warrior struct {
	Trait
}

// Mage ; hérite des Trait + possede un bonus de Point de magie et d'Intélligence
type Mage struct {
	Trait
}

// -------------------  Getters & Setters

// GetClasse ;
func (w Warrior) GetClasse() string {
	return w.Classe
}

// GetClasse ;
func (m Mage) GetClasse() string {
	return m.Classe
}

// GetHealthPoint ;
func (w Warrior) GetHealthPoint() int {
	return w.HealthPoint
}

// GetHealthPoint ;
func (m Mage) GetHealthPoint() int {
	return m.HealthPoint
}

// GetMagicPoint ;
func (w Warrior) GetMagicPoint() int {
	return w.MagicPoint
}

// GetMagicPoint ;
func (m Mage) GetMagicPoint() int {
	return m.MagicPoint
}

// GetStrength ;
func (w Warrior) GetStrength() int {
	return w.Strength
}

// GetStrength ;
func (m Mage) GetStrength() int {
	return m.Strength
}

// GetIntelect ;
func (w Warrior) GetIntelect() int {
	return w.Intelect
}

// GetIntelect ;
func (m Mage) GetIntelect() int {
	return m.Intelect
}

// ------------------- Fin de Getters & Setters

// CreateClass ;
func CreateClass(class string, bonusHp, bonusMp, bonusS, bonusI int) Hero {

	// Source d'aleatoire
	rand.Seed(time.Now().UnixNano())

	hp := 100 + bonusHp // points de vie
	mp := 100 + bonusMp // points de magie
	s := 20 + bonusS    // Strength
	i := 20 + bonusI    // Intelect

	// valeur max des points bonus ajouté aux stats
	BonusCap := 50
	// pourcentage maximum dont peut etre augmenter le BonusCap
	MaxBonusPourcentage := 30

	EffectiveBonus := BonusCap + BonusCap*rand.Intn(MaxBonusPourcentage)/100

	BonusPoint := (rand.Intn(EffectiveBonus))
	BonusTrait := EffectiveBonus - BonusPoint

	switch class {
	case "Warrior":
		hp = hp + BonusPoint
		s = s + BonusTrait
		t := Trait{
			HealthPoint: hp,
			MagicPoint:  mp,
			Strength:    s,
			Intelect:    i,
			Classe:      "warrior",
		}
		h := Warrior{t}
		return h

	case "Mage":
		mp = mp + BonusPoint
		i = i + BonusTrait
		t := Trait{
			HealthPoint: hp,
			MagicPoint:  mp,
			Strength:    s,
			Intelect:    i,
			Classe:      "mage",
		}
		h := Mage{t}
		return h
	}
	return nil
}
