package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printcolor(str string) {
	if str == "W" {
		c := color.New(color.FgWhite)
		c.Print("W")
	} else if str == "R" {
		c := color.New(color.FgRed)
		c.Print("R")
	} else if str == "O" {
		c := color.New(color.FgMagenta)
		c.Print("O")
	} else if str == "B" {
		c := color.New(color.FgBlue)
		c.Print("B")
	} else if str == "G" {
		c := color.New(color.FgGreen)
		c.Print("G")
	} else if str == "Y" {
		c := color.New(color.FgYellow)
		c.Print("Y")
	}
}

func printVerticalCube(tmpface [9]string) {
	i := 0
	order := [9]int{0, 1, 2, 7, 8, 3, 6, 5, 4}
	for x := 0; x < 3; x++ {
		for o := 0; o < 4; o++ {
			fmt.Print(" ")
		}
		for o := 0; o < 3; o++ {
			printcolor(tmpface[order[i]])
			i++
		}
		fmt.Println()
	}
	fmt.Println()
}

func printMiddleCube(tmpStockface [3][9]string) {
	order := [9]int{0, 1, 2, 7, 8, 3, 6, 5, 4}
	for x := 0; x < 3; x++ {
		for o := 0; o < 3; o++ {
			printcolor(tmpStockface[x][order[o]])
		}
		fmt.Print(" ")
	}
	fmt.Println()
	for x := 0; x < 3; x++ {
		for o := 3; o < 6; o++ {
			printcolor(tmpStockface[x][order[o]])
		}
		fmt.Print(" ")
	}
	fmt.Println()
	for x := 0; x < 3; x++ {
		for o := 6; o < 9; o++ {
			printcolor(tmpStockface[x][order[o]])
		}
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println()
}

func debugPrint(cube [6]int32) {
	str := "OGBWYR" // "WROBGY"
	order := [6]int{5, 3, 2, 0, 1, 4}
	var tmpStockface [3][9]string
	x := 0

	for _, j := range order {
		face := cube[j]
		var tmpface [9]string
		tmpface[8] = string(str[j])
		for i := 7; i >= 0; i-- {
			var nbr int32
			nbr = face & 15
			tmpface[i] = string(str[nbr])
			face = face >> 4
		}
		if j == 5 || j == 3 || j == 4 {
			printVerticalCube(tmpface)
		}
		if j == 2 || j == 0 || j == 1 {
			tmpStockface[x] = tmpface
			x++
		}
		if j == 1 {
			printMiddleCube(tmpStockface)
		}
	}
}
