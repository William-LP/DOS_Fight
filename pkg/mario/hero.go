package mario

// Hero : comment
type Hero struct {
	Pointdevie,
	Pointdemagie,
	Endurance,
	Force,
	Intelligence int
}

// CreateHero allows to build a hero
func CreateHero(pointdevie, pointdemagie, endurance, force, intelligence int) Hero {
	hero := Hero{
		Pointdevie:   pointdevie,
		Pointdemagie: pointdemagie,
		Endurance:    endurance,
		Force:        force,
		Intelligence: intelligence,
	}
	return hero
}
