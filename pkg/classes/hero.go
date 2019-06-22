package main

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
	BonusPointDeVie int
	BonusForce      int
}

// Mage ; hérite des Carac + possede un bonus de Point de magie et d'Intélligence
type Mage struct {
	Carac
	BonusPointDeMagie int
	BonusIntelligence int
}
