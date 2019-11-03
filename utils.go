package main

func (env *Env) copyCube(cube [6][3][3]int) [6][3][3]int {
	var copyCube [6][3][3]int
	for face := range cube {
		for line := range cube[face] {
			for col := range cube[face][line] {
				copyCube[face][line][col] = cube[face][line][col]
			}
		}
	}
	return copyCube
}

func (env *Env) isFinished(currCube CubeEnv) bool {
	if currCube.cube == env.solvedCube {
		return true
	}
	return false
}

func existInClosedList(currCube CubeEnv, closedList []CubeEnv) bool {
	for i := range closedList {
		if closedList[i].cube == currCube.cube {
			return true
		}
	}
	return false
}

func (env *Env) setCube() {
	i := 0
	for face := range env.currentCube.cube {
		for line := range env.currentCube.cube[face] {
			for col := range env.currentCube.cube[face][line] {
				env.currentCube.cube[face][line][col] = i
				env.solvedCube[face][line][col] = i
				i++
			}
		}
	}
}

func (env *Env) shuffle() {
	for step := range env.mix {
		// exec step
		env.execStep(env.mix[step])
	}
}
