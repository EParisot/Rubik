package main

import "fmt"

func (env *Env) debugPrint(step string, cube [6][3][3]int) {
	fmt.Printf("step : %s\n", step)
	for i := range cube[5] {
		fmt.Print("\t\t")
		for _, val := range cube[5][i] {
			fmt.Printf("%.2d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := range cube[3] {
		fmt.Print("\t\t")
		for _, val := range cube[3][i] {
			fmt.Printf("%.2d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := range cube[2] {
		for _, val := range cube[2][i] {
			fmt.Printf("%.2d ", val)
		}
		fmt.Print("\t")
		for _, val := range cube[0][i] {
			fmt.Printf("%.2d ", val)
		}
		fmt.Print("\t")
		for _, val := range cube[1][i] {
			fmt.Printf("%.2d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := range cube[4] {
		fmt.Print("\t\t")
		for _, val := range cube[4][i] {
			fmt.Printf("%.2d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}
