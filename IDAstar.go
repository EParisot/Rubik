package main

import (
	"fmt"
	"math/rand"
)

var moves [6]string
var max [3]int
var result [5]string

func (env *Env) idAstar() {
	moves = [6]string{"F", "R", "L", "U", "D", "B"}
	max = [3]int{13, 13, 15} // max dept limits
	var closedList []CubeEnv
	closedList = append(closedList, env.currentCube)
	phase := 0
	threshold := globalHeuristic(env.currentCube, phase)
	for {
		tmpThres, _ := env.search(threshold, &closedList, &phase, 0)
		if tmpThres == -1 {
			var finalResult string
			for _, res := range result {
				finalResult += res
			}
			finalResult = parseOutput(finalResult)
			fmt.Println(finalResult)
			return
		} else if tmpThres == -2 || tmpThres >= 10000 {
			return
		}
		threshold = tmpThres
	}
}

func (env *Env) search(threshold int, closedList *[]CubeEnv, phase *int, depth int) (int, *[]CubeEnv) {
	// Handle dept limit
	depth++
	if *phase < 3 && depth >= max[*phase] {
		// reset cube and result
		env.currentCube = env.startCube
		result = [5]string{}
		// generate new random steps
		for i := 0; i < 5; i++ {
			randIdx := rand.Intn(6)
			env.execStep(moves[randIdx])
			result[0] += moves[randIdx] + " "
		}
		if env.debug {
			fmt.Println("RESET AND MIX MORE")
			debugCube(env.currentCube.cube)
			fmt.Println(result)
		}
		env.idAstar()
		return -2, closedList
	}
	// IDAstar threshold
	currCube := (*closedList)[len(*closedList)-1]
	if currCube.heuristic+currCube.cost > threshold {
		return currCube.heuristic + currCube.cost, closedList
	}
	// Phases transitions
	if *phase == 0 && isInG1(currCube) == 0 {
		for _, step := range (*closedList)[1:len(*closedList)] {
			result[*phase] += step.internationalMove + " "
		}
		*phase = 1
		currCube.cost = 0
		threshold = isInG2(currCube)
		if env.debug {
			fmt.Println("Phase0 DONE")
			debugCube(currCube.cube)
			debugPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			fmt.Print("\n")
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
	}
	if *phase == 1 && isInG2(currCube) == 0 {
		for _, step := range (*closedList)[1:len(*closedList)] {
			result[*phase] += step.internationalMove + " "
		}
		*phase = 2
		currCube.cost = 0
		threshold = isInG3(currCube)
		if env.debug {
			fmt.Println("Phase1 DONE")
			debugCube(currCube.cube)
			debugPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			fmt.Print("\n")
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
	}
	if *phase == 2 && isInG3(currCube) == 0 {
		for _, step := range (*closedList)[1:len(*closedList)] {
			result[*phase] += step.internationalMove + " "
		}
		*phase = 3
		currCube.cost = 0
		threshold = isInGc(currCube)
		if env.debug {
			fmt.Println("Phase2 DONE")
			debugCube(currCube.cube)
			debugPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			fmt.Print("\n")
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
	}
	if isFinished(currCube) {
		for i, step := range (*closedList)[1:len(*closedList)] {
			result[*phase] += step.internationalMove
			if i < len(*closedList)-2 {
				result[*phase] += " "
			}
		}
		*phase = 4
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
	for _, child := range getMoves(currCube, *phase) {
		if !existInClosedList(child, *closedList) {
			*closedList = append(*closedList, child)
			result, closedList := env.search(threshold, closedList, phase, depth)
			if result == -1 || result == -2 {
				return result, closedList
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
