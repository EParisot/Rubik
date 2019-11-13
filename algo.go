package main

import (
	"fmt"
)

var moves []string

func (env *Env) idAstar() {
	moves = []string{"F", "R", "L", "U", "D", "B"}
	var closedList []CubeEnv
	closedList = append(closedList, env.currentCube)
	phase := 1
	threshold := env.globalHeuristic(env.currentCube, phase)
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
	if currCube.heuristic+currCube.cost > threshold {
		return currCube.heuristic + currCube.cost, closedList
	}
	if *phase == 1 && isInG1(currCube) == 0 {
		*phase = 2
		currCube.cost = 0
		threshold = isInG2(currCube)
		if env.debug {
			fmt.Println("Phase1 DONE")
			env.debugPrint(currCube.cube)
		}
		env.reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
		if len(*closedList) > 1 {
			fmt.Print(" ")
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
	}
	if *phase == 2 && isInG2(currCube) == 0 {
		*phase = 3
		currCube.cost = 0
		threshold = isInG3(currCube)
		if env.debug {
			fmt.Println("\nPhase2 DONE")
			env.debugPrint(currCube.cube)
		}
		env.reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
		if len(*closedList) > 1 {
			fmt.Print(" ")
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
	}
	if *phase == 3 && isInG3(currCube) == 0 {
		*phase = 4
		currCube.cost = 0
		threshold = isInGc(currCube)
		if env.debug {
			fmt.Println("\nPhase3 DONE")
			env.debugPrint(currCube.cube)
		}
		env.reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
		if len(*closedList) > 1 {
			fmt.Print(" ")
		}
		*closedList = (*closedList)[len(*closedList)-1 : len(*closedList)]
	}
	if env.isFinished(currCube) {
		if env.debug {
			fmt.Println("\nALL DONE")
			env.debugPrint(currCube.cube)
		}
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
	var cubeList []CubeEnv
	for rotate := 0; rotate <= 5; rotate++ {
		way := 0
		var newEnvCube CubeEnv
		newEnvCube.cube = env.rotate(rotate, way, currCube.cube)
		var nb string
		if (phase == 2 && (rotate == 0 || rotate == 5)) || // <F2 R L U D B2>
			(phase == 3 && (rotate == 1 || rotate == 2 || rotate == 0 || rotate == 5)) || // <F2 R2 L2 U D B2>
			phase == 4 { // <F2 R2 L2 U2 D2 B2>
			newEnvCube.cube = env.rotate(rotate, way, newEnvCube.cube)
			nb = "2"
		}
		newEnvCube.internationalMove = moves[rotate]
		newEnvCube.internationalMove = newEnvCube.internationalMove + nb
		newEnvCube.cost = currCube.cost + 1
		newEnvCube.heuristic = env.globalHeuristic(newEnvCube, phase)
		cubeList = append(cubeList, newEnvCube)
	}
	return cubeList
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
	} else if phase == 4 {
		gHeur = isInGc(currCube)
	}
	return gHeur
}

// fixes RL and UD Edges orientation
func isInG1(currCube CubeEnv) int {
	var latFacelets int
	var topDownFacelets int
	for _, face := range []int{1, 2} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 3 && int(currCube.cube[face]>>uint(facelet*4))&15 != 4 {
				latFacelets++
			}
		}
	}
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 1 && int(currCube.cube[face]>>uint(facelet*4))&15 != 2 {
				topDownFacelets++
			}
		}
	}
	return 16 - int(latFacelets/2+topDownFacelets/2)
	/*	var latEdges int
		var topDownEdges int
		var faceEdges int
		for _, face := range []int{1, 2} {
			var oppositeFace int
			var oppositeNextFace int
			if face == 1 {
				oppositeFace = 2
			} else {
				oppositeFace = 1
			}
			for _, facelet := range []int{0, 4} {
				var nextFace int
				if face == 1 && facelet == 0 {
					nextFace = 0
					oppositeNextFace = 5
				} else if face == 1 && facelet == 4 {
					nextFace = 5
					oppositeNextFace = 0
				} else if face == 2 && facelet == 0 {
					nextFace = 5
					oppositeNextFace = 0
				} else if face == 2 && facelet == 4 {
					nextFace = 0
					oppositeNextFace = 5
				}
				if int(currCube.cube[face]>>uint(facelet*4))&15 == face || int(currCube.cube[face]>>uint(facelet*4))&15 == nextFace ||
					int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeFace || int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeNextFace {
					latEdges++
				}
			}
		}
		for _, face := range []int{3, 4} {
			var oppositeFace int
			var oppositeNextFace int
			if face == 3 {
				oppositeFace = 4
			} else {
				oppositeFace = 3
			}
			for _, facelet := range []int{0, 4} {
				var nextFace int
				if face == 3 && facelet == 0 {
					nextFace = 2
					oppositeNextFace = 1
				} else if face == 3 && facelet == 4 {
					nextFace = 1
					oppositeNextFace = 2
				} else if face == 4 && facelet == 0 {
					nextFace = 2
					oppositeNextFace = 1
				} else if face == 4 && facelet == 4 {
					nextFace = 1
					oppositeNextFace = 2
				}
				if int(currCube.cube[face]>>uint(facelet*4))&15 == face || int(currCube.cube[face]>>uint(facelet*4))&15 == nextFace ||
					int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeFace || int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeNextFace {
					topDownEdges++
				}
			}
		}
		for _, face := range []int{0, 5} {
			var oppositeFace int
			var oppositeNextFace int
			if face == 5 {
				oppositeFace = 0
			} else {
				oppositeFace = 5
			}
			for _, facelet := range []int{0, 4} {
				var nextFace int
				if face == 0 && facelet == 0 {
					nextFace = 2
					oppositeNextFace = 1
				} else if face == 0 && facelet == 4 {
					nextFace = 1
					oppositeNextFace = 2
				} else if face == 5 && facelet == 0 {
					nextFace = 2
					oppositeNextFace = 1
				} else if face == 5 && facelet == 4 {
					nextFace = 1
					oppositeNextFace = 2
				}
				if int(currCube.cube[face]>>uint(facelet*4))&15 == face || int(currCube.cube[face]>>uint(facelet*4))&15 == nextFace ||
					int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeFace || int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeNextFace {
					faceEdges++
				}
			}
		}
		return 6 - int(latEdges/2+topDownEdges/2+faceEdges/2)*/
}

// Fixes UD facelets orientations and midEdges in midLayer
func isInG2(currCube CubeEnv) int {
	var topDownFacelets int
	var midEdges int
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} { // TO CMP with {1, 3, 5, 7}
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
	/*var cornerParity int
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{1, 3, 5, 7} {
			var lFacelet int
			var rFacelet int
			if facelet == 1 {
				lFacelet = 7
				rFacelet = 3
			} else if facelet == 3 {
				lFacelet = 1
				rFacelet = 5
			} else if facelet == 5 {
				lFacelet = 3
				rFacelet = 7
			} else {
				lFacelet = 5
				rFacelet = 1
			}
			if ((int(currCube.cube[face]>>uint(lFacelet*4))&15) != (int(currCube.cube[face]>>uint(facelet*4))&15) &&
				(int(currCube.cube[face]>>uint(rFacelet*4))&15) == (int(currCube.cube[face]>>uint(facelet*4))&15)) ||
				(int(currCube.cube[face]>>uint(lFacelet*4))&15) != (int(currCube.cube[face]>>uint(rFacelet*4))&15) {
				cornerParity++
			}
		}
	}
	if cornerParity == 0 || cornerParity == 8 {
		cornerParity = 8
	} else {
		cornerParity = 0
	}
	return 16 - int(topDownFacelets/2+midEdges/2+cornerParity/2)
	*/return 12 - int(topDownFacelets/2+midEdges/2)
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
	return 14 - int(facelets/4+parity/4)
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
	return 24 - int(corners/2+edges/2)
}
