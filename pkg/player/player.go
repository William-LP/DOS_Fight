package player

type Player struct {
	Name string,
	Gold int
}

func NewPlayer(name string) Player {
	np := {
		Name : name,
		Gold : 50
	}
	return np
}



