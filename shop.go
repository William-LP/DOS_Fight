package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/William-LP/DOS_Fight/pkg/player"
	"github.com/common-nighthawk/go-figure"
)

type Items struct {
	XMLName xml.Name      `xml:"Items"`
	Item    []player.Item `xml:"Item"`
}

/* type Item struct {
	XMLName     xml.Name `xml:"Item"`
	Type        string   `xml:"type,attr"`
	Id          string   `xml:"Id"`
	Name        string   `xml:"Name"`
	Price       int      `xml:"Price"`
	HealthPoint int      `xml:"HealthPoint"`
	MagicPoint  int      `xml:"MagicPoint"`
	Strength    int      `xml:"Strength"`
	Intelect    int      `xml:"Intelect"`
} */

func shop(p player.Player) {
	clearScreen()
	var items Items
	var stuffs, potions []player.Item
	/* Using a XML file
	xmlFile, err := os.Open("../xml/items.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// xml.Unmarshal(byteValue, &items)

	*/

	// Using a XML var
	itemsXML := `<?xml version="1.0" encoding="UTF-8"?><Items><Item type="stuff"><Id>1</Id><Name>Sword of Truth</Name><Price>50</Price><HealthPoint>10</HealthPoint><MagicPoint>5</MagicPoint><Strength>12</Strength><Intelect>6</Intelect></Item><Item type="stuff"><Id>2</Id><Name>Helmet of Wisdom</Name><Price>50</Price><HealthPoint>5</HealthPoint><MagicPoint>10</MagicPoint><Strength>6</Strength><Intelect>12</Intelect></Item><Item type="potion"><Id>3</Id><Name>Second Breath</Name><Price>10</Price><HealthPoint>20</HealthPoint><MagicPoint>0</MagicPoint><Strength>0</Strength><Intelect>0</Intelect></Item><Item type="potion"><Id>4</Id><Name>Nature's call</Name><Price>10</Price><HealthPoint>0</HealthPoint><MagicPoint>20</MagicPoint><Strength>0</Strength><Intelect>0</Intelect></Item></Item`
	xml.Unmarshal([]byte(itemsXML), &items)

	for i := 0; i < len(items.Item); i++ {
		switch items.Item[i].Type {
		case "stuff":
			stuffs = append(stuffs, items.Item[i])
		case "potion":
			potions = append(potions, items.Item[i])
		}
	}

	ShopFigure := figure.NewFigure("Shop", "small", true)
	ShopFigure.Print()
	fmt.Println()
	fmt.Printf("Welcome to the shop %s ! You´ve got %d gold(s).\n", p.Name, p.Gold)
	fmt.Println()
	fmt.Println("See the stuffs (1)")
	fmt.Println("See the potions (2)")
	fmt.Println("Sell items (3)")
	fmt.Println("Exit shop (4)")
	switch getInput("Choice : ") {
	case "1":
		SubFigure := figure.NewFigure("Stuff", "cybermedium", true)
		SubFigure.Print()
		showItems(stuffs, p)
	case "2":
		SubFigure := figure.NewFigure("Potions", "cybermedium", true)
		SubFigure.Print()
		showItems(potions, p)
	case "3":
		sellItems(p)
	case "4":
		mainMenu(p)
	default:
		fmt.Println("Error, please make a correct choice")
		shop(p)
	}

}

func sellItems(p player.Player) {

}

func showItems(items []player.Item, p player.Player) {
	for i := 0; i < len(items); i++ {
		fmt.Println("\n+------------------------------+")
		fmt.Println(items[i].Id + " - " + items[i].Name + "	(" + strconv.Itoa(items[i].Price) + " golds)")
		if items[i].HealthPoint > 0 {
			fmt.Printf("+%d hp ", items[i].HealthPoint)
		}
		if items[i].MagicPoint > 0 {
			fmt.Printf("+%d mp ", items[i].MagicPoint)
		}
		if items[i].Strength > 0 {
			fmt.Printf("+%d str ", items[i].Strength)
		}
		if items[i].Intelect > 0 {
			fmt.Printf("+%d int ", items[i].Intelect)
		}
	}
	fmt.Println("\n+------------------------------+")
	fmt.Println()
	choice := getInput("Item to buy (press 'c' ton cancel) : ")
	if choice == "c" {
		shop(p)
	}
	var found = false

	for _, i := range items {
		if i.Id == choice {
			found = true
			buyItem(p, i)
		}
	}
	if found == false {
		fmt.Println()
		fmt.Println("Impossible de trouver l'item a buy")
		pause(2)
		shop(p)

	}

}

func pause(x int) {
	duration := time.Duration(x) * time.Second // Pause for 10 seconds
	time.Sleep(duration)
}

func buyItem(p player.Player, i player.Item) {
	if p.Gold >= i.Price {
		p.Gold = p.Gold - i.Price
		if i.Type == "stuff" {
			p.BonusHealthPoint = p.BonusHealthPoint + i.HealthPoint
			p.BonusMagicPoint = p.BonusMagicPoint + i.MagicPoint
			p.BonusStrength = p.BonusStrength + i.Strength
			p.BonusIntelect = p.BonusIntelect + i.Intelect
		}
		p.Inventaire = append(p.Inventaire, i)
		shop(p)
	} else {
		fmt.Println()
		fmt.Println("Not enough golds !")
		pause(2)
		shop(p)
	}
}
