package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	file, err := os.Open("./fonts.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Println(scanner.Text())
		myFigure := figure.NewFigure("DOS Fight", scanner.Text(), true)
		myFigure.Print()
		time.Sleep(2 * time.Second)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
