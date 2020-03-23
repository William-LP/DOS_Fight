package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/William-LP/go_training/pkg/player"
	"github.com/common-nighthawk/go-figure"
)

type Items struct {
	XMLName       xml.Name `xml:"Items"`
	Items_stuffs  Stuffs   `xml:"Stuffs"`
	Items_potions Potions  `xml:"Potions"`
}

type Stuffs struct {
	XMLName xml.Name `xml:"Stuffs"`
	Stuff   []Stuff  `xml:"Stuff"`
}

type Potions struct {
	XMLName xml.Name `xml:"Potions"`
	Potion  []Potion `xml:"Potion"`
}

type Stuff struct {
	XMLName     xml.Name `xml:"Stuff"`
	Id          string   `xml:"Id"`
	Name        string   `xml:"Name"`
	Price       string   `xml:"Price"`
	HealthPoint int      `xml:"HealthPoint"`
	MagicPoint  int      `xml:"MagicPoint"`
	Strength    int      `xml:"Strength"`
	Intelect    int      `xml:"Intelect"`
}

type Potion struct {
	XMLName     xml.Name `xml:"Potion"`
	Id          string   `xml:"Id"`
	Name        string   `xml:"Name"`
	Price       string   `xml:"Price"`
	HealthPoint int      `xml:"HealthPoint"`
	MagicPoint  int      `xml:"MagicPoint"`
	Strength    int      `xml:"Strength"`
	Intelect    int      `xml:"Intelect"`
}

func shop(p player.Player) {
	xmlFile, err := os.Open("../xml/items.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var items Items

	xml.Unmarshal(byteValue, &items)
	ShopFigure := figure.NewFigure("Shop", "small", true)
	ShopFigure.Print()
	fmt.Println()
	fmt.Printf("Welcome to the shop %s ! You´ve got %d gold(s).\n", p.Name, p.Gold)
	fmt.Println()
	fmt.Println("See the stuffs (1)")
	fmt.Println("See the potions (2)")
	fmt.Println("Exit shop (3)")
	fmt.Println()
	switch getInput("Choice : ") {
	case "1":
		showStuff(items)
	case "2":
		showPotions(items)
	case "3":
		mainMenu(p)
	default:
		fmt.Println("Error, please make a correct choice")
		shop(p)
	}

}

func showStuff(items Items) {
	StuffsFigure := figure.NewFigure("Stuff", "cybermedium", true)
	StuffsFigure.Print()

	for i := 0; i < len(items.Items_stuffs.Stuff); i++ {
		fmt.Println("\n+------------------------------+")
		fmt.Println(items.Items_stuffs.Stuff[i].Id + " - " + items.Items_stuffs.Stuff[i].Name + "	(" + items.Items_stuffs.Stuff[i].Price + " golds)")
		if items.Items_stuffs.Stuff[i].HealthPoint > 0 {
			fmt.Printf("+%d hp ", items.Items_stuffs.Stuff[i].HealthPoint)
		}
		if items.Items_stuffs.Stuff[i].MagicPoint > 0 {
			fmt.Printf("+%d mp ", items.Items_stuffs.Stuff[i].MagicPoint)
		}
		if items.Items_stuffs.Stuff[i].Strength > 0 {
			fmt.Printf("+%d str ", items.Items_stuffs.Stuff[i].Strength)
		}
		if items.Items_stuffs.Stuff[i].Intelect > 0 {
			fmt.Printf("+%d int ", items.Items_stuffs.Stuff[i].Intelect)
		}
	}
}

func showPotions(items Items) {
	PotionsFigure := figure.NewFigure("Potions", "cybermedium", true)
	PotionsFigure.Print()

	for i := 0; i < len(items.Items_potions.Potion); i++ {
		fmt.Println("\n+------------------------------+")
		fmt.Println(items.Items_stuffs.Stuff[i].Id + " - " + items.Items_potions.Potion[i].Name + "	(" + items.Items_potions.Potion[i].Price + " golds)")
		if items.Items_potions.Potion[i].HealthPoint > 0 {
			fmt.Printf("+%d hp ", items.Items_potions.Potion[i].HealthPoint)
		}
		if items.Items_potions.Potion[i].MagicPoint > 0 {
			fmt.Printf("+%d mp ", items.Items_potions.Potion[i].MagicPoint)
		}
		if items.Items_potions.Potion[i].Strength > 0 {
			fmt.Printf("+%d str ", items.Items_potions.Potion[i].Strength)
		}
		if items.Items_potions.Potion[i].Intelect > 0 {
			fmt.Printf("+%d int ", items.Items_potions.Potion[i].Intelect)
		}
	}
}
