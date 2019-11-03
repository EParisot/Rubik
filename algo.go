package main

import (
	"fmt"
)

func (env *Env) idAstar() {
	fmt.Println("Begin IDAstar")
	threshold := env.globalHeuristic(env.currentCube)
	var closedList []CubeEnv
	closedList = append(closedList, env.currentCube)
	for {
		tmpThres, closedList := env.search(threshold, &closedList)
		if tmpThres == -1 {
			fmt.Println("IDAstar Done")
			env.reconstructPathIDA(*closedList, (*closedList)[len(*closedList)-1])
			return
		} else if tmpThres >= 10000 {
			fmt.Println("IDAstar returned no solution")
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
	if env.isFinished(currCube) {
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
	for rotate := 0; rotate < 5; rotate++ {
		for way := 0; way < 2; way++ {
			copyCube := env.copyCube(currCube.cube)
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
			newEnvCube.heuristic = newEnvCube.cost

			//env.debugPrint(newEnvCube.internationalMove, newEnvCube.cube)
			gridList = append(gridList, newEnvCube)
		}
	}
	return gridList
}

func (env *Env) reconstructPathIDA(closedList []CubeEnv, endGrid CubeEnv) {
	fmt.Println("Ordered sequence of states that make up the solution : ")
	for _, step := range closedList {
		fmt.Println(step.internationalMove)
	}
	fmt.Println("Number of moves required : ", len(closedList)-1)
}

func (env *Env) globalHeuristic(currCube CubeEnv) int {
	gHeur := 0
	return gHeur
}
