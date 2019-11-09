package main

import (
	"fmt"
)

func (env *Env) idAstar() {
	//fmt.Println("Begin IDAstar")
	phase := 1
	threshold := env.globalHeuristic(env.currentCube, phase)
	var closedList []CubeEnv
	closedList = append(closedList, env.currentCube)
	for {
		tmpThres, closedList := env.search(threshold, &closedList, &phase)
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

func (env *Env) search(threshold int, closedList *[]CubeEnv, phase *int) (int, *[]CubeEnv) {
	currCube := (*closedList)[len(*closedList)-1]
	if currCube.heuristic > threshold {
		return currCube.heuristic, closedList
	}
	if *phase == 1 && isInG1(currCube) == 0 &&
		(currCube.internationalMove == "R" ||
			currCube.internationalMove == "L" ||
			currCube.internationalMove == "F" ||
			currCube.internationalMove == "B") {
		*phase = 2
		currCube.cost = 0
		fmt.Println("Phase1 DONE")
		env.debugPrint(currCube.cube)
		//return -1, closedList
	} else if *phase == 2 && isInG2(currCube) == 0 &&
		(currCube.internationalMove == "U" ||
			currCube.internationalMove == "D") {
		*phase = 3
		currCube.cost = 0
		fmt.Println("Phase2 DONE")
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
		for way := 0; way < 2; way++ {
			copyCube := env.copyCube(currCube.cube) // Check if needed
			newCube := env.rotate(rotate, way, copyCube)
			var nb string
			if phase == 3 || (phase == 2 && (rotate == 0 || rotate == 1 || rotate == 5 || rotate == 2)) {
				newCube = env.rotate(rotate, way, newCube)
				nb = "2"
			}
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
			newEnvCube.internationalMove = newEnvCube.internationalMove + nb
			newEnvCube.cube = newCube
			//tmp
			newEnvCube.cost = currCube.cost + 1
			newEnvCube.heuristic = newEnvCube.cost + env.globalHeuristic(newEnvCube, phase)
			/*if phase == 2 {
				fmt.Println(phase, newEnvCube.cost, env.globalHeuristic(newEnvCube, phase))
			}*/
			//env.debugPrint(newEnvCube.internationalMove, newEnvCube.cube)
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

func (env *Env) globalHeuristic(currCube CubeEnv, phase int) int {
	var gHeur int
	if phase == 1 {
		gHeur = isInG1(currCube)
	} else if phase == 2 {
		gHeur = isInG2(currCube)
	} else {
		gHeur = isInGc(currCube)
	}
	return gHeur
}

func cornerIsOriented(currCube CubeEnv, face int, facelet int32) bool {
	var oppositeFace int
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

func midEdgeIsGood(currCube CubeEnv, facelet int32, nextFacelet int32) bool {
	faces := []int{5, 0, 1, 2}
	for _, face := range faces {
		if int(facelet) == face {
			for _, nextFace := range faces {
				if nextFace != face && int(nextFacelet) == nextFace {
					return true
				}
			}
		}
	}
	return false
}

func latEdgeIsGood(currCube CubeEnv, facelet int32, nextFacelet int32) bool {
	faces := []int{1, 2, 3, 4}
	for _, face := range faces {
		if int(facelet) == face {
			for _, nextFace := range faces {
				if nextFace != face && int(nextFacelet) == nextFace {
					return true
				}
			}
		}
	}
	return false
}

func sagEdgeIsGood(currCube CubeEnv, facelet int32, nextFacelet int32) bool {
	faces := []int{0, 3, 4, 5}
	for _, face := range faces {
		if int(facelet) == face {
			for _, nextFace := range faces {
				if nextFace != face && int(nextFacelet) == nextFace {
					return true
				}
			}
		}
	}
	return false
}

func isInG1(currCube CubeEnv) int {
	var corners int
	var midEdges int
	var topDownEdges int
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{1, 3, 5, 7} {
			if cornerIsOriented(currCube, face, (currCube.cube[face]>>uint(facelet*4))&15) {
				corners++
			}
		}
	}
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 2, 4, 6} {
			var oppositeFace int
			if face == 3 {
				oppositeFace = 4
			} else {
				oppositeFace = 3
			}
			if (int(currCube.cube[face]>>uint(facelet*4))&15) == face || (int(currCube.cube[face]>>uint(facelet*4))&15) == oppositeFace {
				topDownEdges++
			}
		}
	}
	faces := []int{0, 1, 5, 2}
	for i, face := range faces {
		for _, facelet := range []int{0, 4} {
			var nextFace int
			var nextFacelet int
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
			if midEdgeIsGood(currCube, (currCube.cube[face]>>uint(facelet*4))&15, (currCube.cube[nextFace]>>uint(nextFacelet*4))&15) {
				midEdges++
			}
		}
	}
	return 20 - (corners + midEdges/2 + topDownEdges)
}

func isInG2(currCube CubeEnv) int {
	var corners int
	var edges int
	for _, face := range []int{0, 1, 2, 5} {
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
		for _, facelet := range []int{1, 3, 5, 7} {
			if (int(currCube.cube[face]>>uint(facelet*4))&15) == face || (int(currCube.cube[face]>>uint(facelet*4))&15) == oppositeFace {
				corners++
			}
		}
	}
	for _, face := range []int{0, 1, 2, 5} {
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
		for _, facelet := range []int{2, 6} {
			if (int(currCube.cube[face]>>uint(facelet*4))&15) == face || (int(currCube.cube[face]>>uint(facelet*4))&15) == oppositeFace {
				edges++
			}
		}
	}
	/*for _, face := range []int{3, 4} {
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
		for _, facelet := range []int{0, 2, 4, 6} {
			if (int(currCube.cube[face]>>uint(facelet*4))&15) == face || (int(currCube.cube[face]>>uint(facelet*4))&15) == oppositeFace {
				edges++
			}
		}
	}*/
	//fmt.Println(corners, edges)
	return 24 - int(corners+edges)
}

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
	//fmt.Println(corners, edges)
	return 48 - int(corners+edges)
}
