package player

// Player ;
type Player struct {
	Name string
	Gold,
	BonusHealthPoint,
	BonusMagicPoint,
	BonusStrength,
	BonusIntelect int
}

// NewPlayer ;
func NewPlayer(name string) Player {
	np := Player{
		Name:             name,
		Gold:             50,
		BonusHealthPoint: 0,
		BonusMagicPoint:  0,
		BonusStrength:    0,
		BonusIntelect:    0,
	}
	return np
}
