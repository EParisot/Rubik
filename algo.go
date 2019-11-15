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
	max = []int{7, 13, 15, 17}
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
			debugPrint(env.currentCube.cube)
			fmt.Println(result)
		}
		env.idAstar()
		os.Exit(0)
	}
	// IDAstar condition
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
			debugPrint(currCube.cube)
			reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
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
			debugPrint(currCube.cube)
			reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
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
			debugPrint(currCube.cube)
			reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
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
			debugPrint(currCube.cube)
			reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			if len(*closedList) > 1 {
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
		return -1, closedList
	}
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

func getMoves(currCube CubeEnv, phase int) []CubeEnv {
	var cubeList []CubeEnv
	for face := 0; face <= 5; face++ {
		var newEnvCube CubeEnv
		newEnvCube.cube = rotate(face, 0, currCube.cube)
		var nb string
		if (phase == 1 && (face == 0 || face == 5)) || // <F2 R L U D B2>
			(phase == 2 && (face == 1 || face == 2 || face == 0 || face == 5)) || // <F2 R2 L2 U D B2>
			phase == 3 { // <F2 R2 L2 U2 D2 B2>
			newEnvCube.cube = rotate(face, 0, newEnvCube.cube)
			nb = "2"
		}
		newEnvCube.internationalMove = moves[face] + nb
		newEnvCube.cost = currCube.cost + 1
		newEnvCube.heuristic = globalHeuristic(newEnvCube, phase)
		cubeList = append(cubeList, newEnvCube)
	}
	return cubeList
}

func reconstructPathIDA(closedList []CubeEnv, endGrid CubeEnv) {
	for i, step := range closedList[1:len(closedList)] {
		fmt.Print(step.internationalMove)
		if i < len(closedList)-2 {
			fmt.Print(" ")
		}
	}
}

func globalHeuristic(currCube CubeEnv, phase int) int {
	var gHeur int
	if phase == 0 {
		gHeur = isInG1(currCube)
	} else if phase == 1 {
		gHeur = isInG2(currCube)
	} else if phase == 2 {
		gHeur = isInG3(currCube)
	} else if phase == 3 {
		gHeur = isInGc(currCube)
	}
	return gHeur
}

// fixes RL and UD Edges orientation
func isInG1(currCube CubeEnv) int {
	var latFacelets int
	var topDownFacelets int
	for _, face := range []int{1, 2} {
		for _, facelet := range []int{0, 2, 4, 6} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 3 && int(currCube.cube[face]>>uint(facelet*4))&15 != 4 {
				latFacelets++
			}
		}
	}
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 2, 4, 6} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 1 && int(currCube.cube[face]>>uint(facelet*4))&15 != 2 {
				topDownFacelets++
			}
		}
	}
	return 4 - int((latFacelets+topDownFacelets)/4)
}

// Fixes UD facelets orientations and midEdges in midLayer
func isInG2(currCube CubeEnv) int {
	var topDownFacelets int
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 == 3 || int(currCube.cube[face]>>uint(facelet*4))&15 == 4 {
				topDownFacelets++
			}
		}
	}
	var midEdges int
	for _, face := range []int{0, 1, 2, 5} {
		for _, facelet := range []int{0, 4} {
			var oppositeFace int
			if face == 0 {
				oppositeFace = 5
			} else if face == 1 {
				oppositeFace = 2
			} else if face == 2 {
				oppositeFace = 1
			} else {
				oppositeFace = 0
			}
			if int(currCube.cube[face]>>uint(facelet*4))&15 == face || int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeFace {
				midEdges++
			}
		}
	}
	return 6 - int((topDownFacelets+midEdges)/4)
}

// Fixed all topDown corners and edges orientation and corners parity
func isInG3(currCube CubeEnv) int {
	var facelets int
	var parity int
	for _, face := range []int{0, 1, 2, 3, 4, 5} {
		var oppositeFace int
		if face%2 == 0 {
			oppositeFace = face - 1
		} else {
			oppositeFace = face + 1
		}
		if oppositeFace == 6 {
			oppositeFace = 0
		} else if oppositeFace == -1 {
			oppositeFace = 5
		}
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if (int(currCube.cube[face]>>uint(facelet*4))&15) == face || (int(currCube.cube[face]>>uint(facelet*4))&15) == oppositeFace {
				facelets++
			}
		}
	}
	for _, face := range []int{0, 1, 2, 5} {
		for _, facelet := range []int{1, 5} {
			var nextFacelet int
			if facelet == 1 {
				nextFacelet = 3
			} else {
				nextFacelet = 7
			}
			if (int(currCube.cube[face]>>uint(facelet*4)) & 15) == (int(currCube.cube[face]>>uint(nextFacelet*4)) & 15) {
				parity++
			}
		}
	}
	if parity == 0 || parity == 8 {
		parity = 8
	} else {
		parity = 0
	}
	return 14 - int((facelets+parity)/4)
}

// Restore solved cube
func isInGc(currCube CubeEnv) int {
	var corners int
	var edges int
	for _, face := range []int{0, 1, 2, 3, 4, 5} {
		for _, facelet := range []int{1, 3, 5, 7} {
			if int((currCube.cube[face]>>uint(facelet*4))&15) == face {
				corners++
			}
		}
	}
	for _, face := range []int{0, 1, 2, 3, 4, 5} {
		for _, facelet := range []int{0, 2, 4, 6} {
			if (int(currCube.cube[face]>>uint(facelet*4)) & 15) == face {
				edges++
			}
		}
	}
	return 12 - int((corners+edges)/4)
}
