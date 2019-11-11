package main

import (
	"fmt"
)

func (env *Env) idAstar() {
	phase := 1
	threshold := env.globalHeuristic(env.currentCube, phase)
	var closedList []CubeEnv
	closedList = append(closedList, env.currentCube)
	for {
		tmpThres, closedList := env.search(threshold, &closedList, &phase)
		if tmpThres == -1 {
			env.reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			return
		} else if tmpThres >= 10000 {
			return
		}
		threshold = tmpThres
	}
}

func (env *Env) search(threshold int, closedList *[]CubeEnv, phase *int) (int, *[]CubeEnv) {
	currCube := (*closedList)[len(*closedList)-1]
	if currCube.heuristic > threshold {
		return currCube.heuristic, closedList
	}
	if *phase == 1 && isInG1(currCube) == 0 {
		*phase = 2
		currCube.cost = 0
		threshold = isInG2(currCube)
		fmt.Println("Phase1 DONE")
		env.debugPrint(currCube.cube)
		//return -1, closedList
	}
	if *phase == 2 && isInG2(currCube) == 0 {
		*phase = 3
		currCube.cost = 0
		threshold = isInG3(currCube)
		fmt.Println("Phase2 DONE")
		env.debugPrint(currCube.cube)
		//return -1, closedList
	}
	if *phase == 3 && isInG3(currCube) == 0 {
		*phase = 4
		currCube.cost = 0
		threshold = isInGc(currCube)
		fmt.Println("Phase3 DONE")
		env.debugPrint(currCube.cube)
		//return -1, closedList
	}
	if env.isFinished(currCube) {
		fmt.Println("ALL DONE")
		env.debugPrint(currCube.cube)
		return -1, closedList
	}
	min := 100000
	childsList := env.getMoves(currCube, *phase)
	for _, child := range childsList {
		if !existInClosedList(child, *closedList) {
			*closedList = append(*closedList, child)
			result, closedList := env.search(threshold, closedList, phase)
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

func (env *Env) getMoves(currCube CubeEnv, phase int) []CubeEnv {
	var gridList []CubeEnv
	for rotate := 0; rotate <= 5; rotate++ {
		//for way := 0; way < 2; way++ {
		way := 0
		//copyCube := env.copyCube(currCube.cube) // Check if needed
		newCube := env.rotate(rotate, way, currCube.cube)
		var nb string
		if (phase == 2 && (rotate == 0 || rotate == 5)) ||
			(phase == 3 && (rotate == 1 || rotate == 2 || rotate == 0 || rotate == 5)) ||
			phase == 4 {
			newCube = env.rotate(rotate, way, newCube)
			nb = "2"
		}
		var newEnvCube CubeEnv
		if rotate == 0 {
			newEnvCube.internationalMove = "F"
		} else if rotate == 1 {
			newEnvCube.internationalMove = "R"
		} else if rotate == 2 {
			newEnvCube.internationalMove = "L"
		} else if rotate == 3 {
			newEnvCube.internationalMove = "U"
		} else if rotate == 4 {
			newEnvCube.internationalMove = "D"
		} else if rotate == 5 {
			newEnvCube.internationalMove = "B"
		}
		if way == 1 {
			newEnvCube.internationalMove = newEnvCube.internationalMove + "'"
		}
		newEnvCube.internationalMove = newEnvCube.internationalMove + nb
		newEnvCube.cube = newCube
		newEnvCube.cost = currCube.cost + 1
		newEnvCube.heuristic = newEnvCube.cost + env.globalHeuristic(newEnvCube, phase)
		//fmt.Println(phase, newEnvCube.cost, env.globalHeuristic(newEnvCube, phase), newEnvCube.internationalMove)
		//env.debugPrint(newEnvCube.internationalMove, newEnvCube.cube)
		gridList = append(gridList, newEnvCube)
	}
	//}
	return gridList
}

func (env *Env) reconstructPathIDA(closedList []CubeEnv, endGrid CubeEnv) {
	for i, step := range closedList[1:len(closedList)] {
		fmt.Print(step.internationalMove)
		if i < len(closedList)-2 {
			fmt.Print(" ")
		}
	}
}

func (env *Env) globalHeuristic(currCube CubeEnv, phase int) int {
	var gHeur int
	if phase == 1 {
		gHeur = isInG1(currCube)
	} else if phase == 2 {
		gHeur = isInG2(currCube)
	} else if phase == 3 {
		gHeur = isInG3(currCube)
	} else {
		gHeur = isInGc(currCube)
	}
	return gHeur
}

// fixes FB Edges orientation
func isInG1(currCube CubeEnv) int {
	var edges int
	var facelets int
	for _, face := range []int{0, 5} {
		for _, facelet := range []int{2, 6} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 3 && int(currCube.cube[face]>>uint(facelet*4))&15 != 4 {
				edges++
			}
		}
		for _, facelet := range []int{0, 4} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 1 && int(currCube.cube[face]>>uint(facelet*4))&15 != 2 {
				edges++
			}
		}
	}
	for _, face := range []int{0, 5} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 3 && int(currCube.cube[face]>>uint(facelet*4))&15 != 4 {
				facelets++
			}
		}
	}
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 0 && int(currCube.cube[face]>>uint(facelet*4))&15 != 5 {
				facelets++
			}
		}
	}
	return 40 - (edges + facelets)
}

// Fixes UD facelets orientations and midEdges in midLayer
func isInG2(currCube CubeEnv) int {
	var topDownFacelets int
	var midEdges int
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 == 3 || int(currCube.cube[face]>>uint(facelet*4))&15 == 4 {
				topDownFacelets++
			}
		}
	}
	for _, face := range []int{0, 1, 2, 5} {
		for _, facelet := range []int{0, 4} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 3 && int(currCube.cube[face]>>uint(facelet*4))&15 != 4 {
				midEdges++
			}
		}
	}
	return 24 - int(topDownFacelets+midEdges)
}

// Fixed all corners and edges orientation
func isInG3(currCube CubeEnv) int {
	var facelets int
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
	return 48 - int(facelets)
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
	return 48 - int(corners+edges)
}
