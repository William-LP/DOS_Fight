package classes

// Hero ;
type Hero interface {
	CreateHero() Hero
}

// Carac ;
type Carac struct {
	PointDeVie,
	PointDeMagie,
	Force,
	Intelligence int
}

// Guerrier ;
type Guerrier struct {
	Carac
	BonusPointDeVie int
	BonusForce      int
}

// Mage ;
type Mage struct {
	Carac
	BonusPointDeMagie int
	BonusIntelligence int
}

// GetPointDeVie Guerrier
func (h Guerrier) GetPointDeVie() int {
	return h.PointDeVie + h.BonusPointDeVie
}

// GetPointDeVie Mage
func (h Mage) GetPointDeVie() int {
	return h.PointDeVie
}

// GetPointDeMagie Guerrier
func (h Guerrier) GetPointDeMagie() int {
	return h.PointDeMagie
}

// GetPointDeMagie Mage
func (h Mage) GetPointDeMagie() int {
	return h.PointDeMagie + h.BonusPointDeMagie
}

// GetForce Guerrier
func (h Guerrier) GetForce() int {
	return h.Force + h.BonusForce
}

// GetForce Mage
func (h Mage) GetForce() int {
	return h.Force
}

// GetIntelligence Guerrier
func (h Guerrier) GetIntelligence() int {
	return h.Intelligence
}

// GetIntelligence Mage
func (h Mage) GetIntelligence() int {
	return h.Intelligence + h.BonusIntelligence
}
