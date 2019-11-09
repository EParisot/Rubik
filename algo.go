package main

import (
	"fmt"
)

func (env *Env) idAstar() {
	//fmt.Println("Begin IDAstar")
	threshold := env.globalHeuristic(env.currentCube)
	var closedList []CubeEnv
	closedList = append(closedList, env.currentCube)
	for {
		tmpThres, closedList := env.search(threshold, &closedList)
		if tmpThres == -1 {
			//fmt.Println("IDAstar Done")
			env.reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			return
		} else if tmpThres >= 10000 {
			//fmt.Println("IDAstar returned no solution")
			return
		}
		threshold = tmpThres
	}
}

func (env *Env) search(threshold int, closedList *[]CubeEnv) (int, *[]CubeEnv) {
	currCube := (*closedList)[len(*closedList)-1]
	if currCube.heuristic > threshold {
		return currCube.heuristic, closedList
	}
	if isInG1(currCube) == 0 { //env.isFinished(currCube) {
		return -1, closedList
	}
	min := 100000
	childsList := env.getMoves(currCube)
	for _, child := range childsList {
		if !existInClosedList(child, *closedList) {
			*closedList = append(*closedList, child)
			result, closedList := env.search(threshold, closedList)
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

func (env *Env) getMoves(currCube CubeEnv) []CubeEnv {
	var gridList []CubeEnv
	for rotate := 0; rotate <= 5; rotate++ {
		for way := 0; way < 2; way++ {
			copyCube := env.copyCube(currCube.cube) // Check if needed
			newCube := env.rotate(rotate, way, copyCube)
			var newEnvCube CubeEnv
			if rotate == 0 {
				newEnvCube.internationalMove = "F"
			} else if rotate == 1 {
				newEnvCube.internationalMove = "R"
			} else if rotate == 3 {
				newEnvCube.internationalMove = "U"
			} else if rotate == 5 {
				newEnvCube.internationalMove = "B"
			} else if rotate == 2 {
				newEnvCube.internationalMove = "L"
			} else if rotate == 4 {
				newEnvCube.internationalMove = "D"
			}
			if way == 1 {
				newEnvCube.internationalMove = newEnvCube.internationalMove + "'"
			}
			newEnvCube.cube = newCube
			//tmp
			newEnvCube.cost = currCube.cost + 1
			newEnvCube.heuristic = newEnvCube.cost + env.globalHeuristic(newEnvCube)

			//	env.debugPrint(newEnvCube.internationalMove, newEnvCube.cube)
			gridList = append(gridList, newEnvCube)
		}
	}
	return gridList
}

func (env *Env) reconstructPathIDA(closedList []CubeEnv, endGrid CubeEnv) {
	//fmt.Println("Ordered sequence of states that make up the solution : ")
	for i, step := range closedList[1:len(closedList)] {
		fmt.Print(step.internationalMove)
		if i < len(closedList)-2 {
			fmt.Print(" ")
		}
	}
	//fmt.Println("Number of moves required : ", len(closedList)-1)
}

func (env *Env) globalHeuristic(currCube CubeEnv) int {
	gHeur := isInG1(currCube)
	return gHeur
}

func cornerIsOriented(currCube CubeEnv, face int, facelet int32) bool {
	oppositeFace := 0
	if face == 3 {
		oppositeFace = 4
	} else {
		oppositeFace = 3
	}
	if int(facelet) == face || int(facelet) == oppositeFace {
		return true
	}
	return false
}

func edgeIsGood(currCube CubeEnv, facelet int32, nextFacelet int32) bool {
	for _, face := range []int{5, 0, 1, 2} {
		if int(facelet) == face {
			for _, nextFace := range []int{5, 0, 1, 2} {
				if nextFace != face && int(nextFacelet) == nextFace {
					return true
				}
			}
		}
	}
	return false
}

func isInG1(currCube CubeEnv) int {
	corners := 0
	midEdges := 0
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{1, 3, 5, 7} {
			if cornerIsOriented(currCube, face, (currCube.cube[face]>>uint(facelet*4))&15) {
				corners++
			}
		}
	}
	faces := []int{0, 1, 5, 2}
	for i, face := range faces {
		for _, facelet := range []int{0, 4} {
			var nextFace int
			nextFacelet := 0
			if face == 5 {
				if facelet == 0 {
					nextFace = 2
					nextFacelet = 0
				} else {
					nextFace = 1
					nextFacelet = 4
				}
			} else {
				if facelet == 0 {
					idx := i - 1
					if idx == -1 {
						idx = 3
					}
					nextFace = faces[idx]
					if nextFace == 5 {
						nextFacelet = 0
					} else {
						nextFacelet = 4
					}
				} else {
					idx := i + 1
					if idx == 4 {
						idx = 0
					}
					nextFace = faces[idx]
					if nextFace == 5 {
						nextFacelet = 4
					} else {
						nextFacelet = 0
					}
				}
			}
			if edgeIsGood(currCube, (currCube.cube[face]>>uint(facelet*4))&15, (currCube.cube[nextFace]>>uint(nextFacelet*4))&15) {
				midEdges++
			}
		}
	}
	return 16 - (corners + midEdges)
}
