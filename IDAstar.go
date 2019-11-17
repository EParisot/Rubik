package main

import (
	"fmt"
	"math/rand"
	"os"
)

var moves []string
var max []int
var result string

func (env *Env) idAstar() {
	moves = []string{"F", "R", "L", "U", "D", "B"}
	max = []int{7, 13, 15, 17} // max dept limits
	var closedList []CubeEnv
	closedList = append(closedList, env.currentCube)
	phase := 0
	threshold := globalHeuristic(env.currentCube, phase)
	for {
		tmpThres, _ := env.search(threshold, &closedList, &phase, 0)
		if tmpThres == -1 {
			result = parseOutput(result)
			fmt.Println(result)
			return
		} else if tmpThres >= 10000 {
			return
		}
		threshold = tmpThres
	}
}

func (env *Env) search(threshold int, closedList *[]CubeEnv, phase *int, depth int) (int, *[]CubeEnv) {
	// Handle dept limit
	depth++
	if depth >= max[*phase] {
		env.currentCube = env.startCube
		result = ""
		for i := 0; i < 10; i++ {
			randIdx := rand.Intn(6)
			env.execStep(moves[randIdx])
			result += moves[randIdx] + " "
		}
		if env.debug {
			fmt.Println("RESET AND MIX MORE")
			debugCube(env.currentCube.cube)
			fmt.Println(result)
		}
		env.idAstar()
		os.Exit(0)
	}
	// IDAstar threshold
	currCube := (*closedList)[len(*closedList)-1]
	if currCube.heuristic+currCube.cost > threshold {
		return currCube.heuristic + currCube.cost, closedList
	}
	// Phases transitions
	if *phase == 0 && isInG1(currCube) == 0 {
		*phase = 1
		currCube.cost = 0
		threshold = isInG2(currCube)
		for _, step := range (*closedList)[1:len(*closedList)] {
			result += step.internationalMove + " "
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
		if env.debug {
			fmt.Println("Phase0 DONE")
			debugCube(currCube.cube)
			debugPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			if len(*closedList) > 1 {
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
	}
	if *phase == 1 && isInG2(currCube) == 0 {
		*phase = 2
		currCube.cost = 0
		threshold = isInG3(currCube)
		for _, step := range (*closedList)[1:len(*closedList)] {
			result += step.internationalMove + " "
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
		if env.debug {
			fmt.Println("Phase1 DONE")
			debugCube(currCube.cube)
			debugPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			if len(*closedList) > 1 {
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
	}
	if *phase == 2 && isInG3(currCube) == 0 {
		*phase = 3
		currCube.cost = 0
		threshold = isInGc(currCube)
		for _, step := range (*closedList)[1:len(*closedList)] {
			result += step.internationalMove + " "
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
		if env.debug {
			fmt.Println("Phase2 DONE")
			debugCube(currCube.cube)
			debugPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			if len(*closedList) > 1 {
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
	}
	if isFinished(currCube) {
		for i, step := range (*closedList)[1:len(*closedList)] {
			result += step.internationalMove
			if i < len(*closedList)-2 {
				result += " "
			}
		}
		if env.debug {
			fmt.Println("ALL DONE")
			debugCube(currCube.cube)
			debugPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			if len(*closedList) > 1 {
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
		return -1, closedList
	}
	//IDAstar
	min := 100000
	childsList := getMoves(currCube, *phase)
	for _, child := range childsList {
		if !existInClosedList(child, *closedList) {
			*closedList = append(*closedList, child)
			result, closedList := env.search(threshold, closedList, phase, depth)
			if result == -1 {
				return -1, closedList
			}
			if result < min {
				min = result
			}
			if len(*closedList) > 1 {
				*closedList = (*closedList)[:len(*closedList)-1]
			} else {
				return result, closedList
			}
		}
	}
	return min, closedList
}
