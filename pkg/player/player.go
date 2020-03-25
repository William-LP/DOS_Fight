package player

import "encoding/xml"

// Player ;
type Player struct {
	Name string
	Gold,
	BonusHealthPoint,
	BonusMagicPoint,
	BonusStrength,
	BonusIntelect,
	Level int
	Inventaire []Item
}

type Item struct {
	XMLName     xml.Name `xml:"Item"`
	Type        string   `xml:"type,attr"`
	Id          string   `xml:"Id"`
	Name        string   `xml:"Name"`
	Price       int      `xml:"Price"`
	HealthPoint int      `xml:"HealthPoint"`
	MagicPoint  int      `xml:"MagicPoint"`
	Strength    int      `xml:"Strength"`
	Intelect    int      `xml:"Intelect"`
}

// NewPlayer ;
func NewPlayer(name string) Player {
	np := Player{
		Name:             name,
		Gold:             120,
		BonusHealthPoint: 0,
		BonusMagicPoint:  0,
		BonusStrength:    0,
		BonusIntelect:    0,
		Level:            1,
	}
	return np
}
