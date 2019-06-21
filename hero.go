package mario

// Hero : comment
type Hero struct {
	pointdevie,
	pointdemagie,
	endurance,
	force,
	intelligence int
}

// ------------- POINT DE VIE ------------------------
func (hero Hero) getPointdevie() int {
	return hero.pointdevie
}

func (hero *Hero) setPointdevie(pointdevie int) {
	hero.pointdevie = pointdevie
}

// ------------- POINT DE MAGIE  ------------------------
func (hero Hero) getPointdemagie() int {
	return hero.pointdemagie
}

func (hero *Hero) setPointdemagie(pointdemagie int) {
	hero.pointdemagie = pointdemagie
}

// ------------- ENDURANCE ------------------------
func (hero Hero) getEndurance() int {
	return hero.endurance
}

func (hero *Hero) setEndurance(endurance int) {
	hero.endurance = endurance
}

// ------------- FORCE ------------------------
func (hero Hero) getForce() int {
	return hero.force
}

func (hero *Hero) setForce(force int) {
	hero.force = force
}

// ------------- INTELLIGENCE ------------------------
func (hero Hero) getIntelligence() int {
	return hero.intelligence
}

func (hero *Hero) setIntelligence(intelligence int) {
	hero.intelligence = intelligence
}

func createHero(pointdevie, pointdemagie, endurance, force, intelligence int) Hero {
	hero := Hero{
		pointdevie:   pointdevie,
		pointdemagie: pointdemagie,
		endurance:    endurance,
		force:        force,
		intelligence: intelligence,
	}
	return hero
}
