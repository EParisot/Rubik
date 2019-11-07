package main

import (
	"fmt"

	"github.com/fatih/color"
)

func (env *Env) printcolor(str string) {
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

func (env *Env) printVerticalCube(tmpface [9]string) {
	i := 0
	for x := 0; x < 3; x++ {
		for o := 0; o < 4; o++ {
			fmt.Print(" ")
		}
		for o := 0; o < 3; o++ {
			env.printcolor(tmpface[i])
			i += 1
		}
		fmt.Println()
	}
	fmt.Println()
}

func (env *Env) printMiddleCube(tmpStockface [3][9]string) {
	for x := 0; x < 3; x++ {
		for o := 0; o < 3; o++ {
			env.printcolor(tmpStockface[x][o])
		}
		fmt.Print(" ")
	}
	fmt.Println()
	for x := 0; x < 3; x++ {
		for o := 3; o < 5; o++ {
			if o == 4 {
				if x == 0 {
					env.printcolor("O")
				}
				if x == 1 {
					env.printcolor("W")
				}
				if x == 2 {
					env.printcolor("R")
				}
			}
			env.printcolor(tmpStockface[x][o])
		}
		fmt.Print(" ")
	}
	fmt.Println()
	for x := 0; x < 3; x++ {
		for o := 5; o < 8; o++ {
			env.printcolor(tmpStockface[x][o])
		}
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println()
}

func (env *Env) debugPrint(step string, cube [6]int32) {
	str := "WROBGY"
	order := [6]int{5, 3, 2, 0, 1, 4}
	var tmpStockface [3][9]string
	x := 0

	for _, j := range order {
		face := cube[j]
		var tmpface [9]string
		for i := 8; i >= 0; i-- {
			if i == 4 {
				tmpface[4] = string(str[j])
				continue
			}
			var nbr int32
			nbr = face & 15
			tmpface[i] = string(str[nbr])
			face = face >> 4
		}
		if j == 5 || j == 3 || j == 4 {
			env.printVerticalCube(tmpface)
		}
		if j == 2 || j == 0 || j == 1 {
			tmpStockface[x] = tmpface
			x += 1
		}
		if j == 1 {
			env.printMiddleCube(tmpStockface)
		}
	}
}
