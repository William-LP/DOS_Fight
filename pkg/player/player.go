package player

// Player ;
type Player struct {
	Name string
	Gold int
}

// NewPlayer ;
func NewPlayer(name string) Player {
	np := Player{
		Name: name,
		Gold: 50,
	}
	return np
}
