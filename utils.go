package main

func (env *Env) copyCube(cube [6]int32) [6]int32 {
	var copyCube [6]int32
	for face := range cube {
		copyCube[face] = cube[face]
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
	env.currentCube.cube[0] = 0b0000000000000000
	env.currentCube.cube[1] = 0b0001000100010001
	env.currentCube.cube[2] = 0b0010001000100010
	env.currentCube.cube[3] = 0b0011001100110011
	env.currentCube.cube[4] = 0b0100010001000100
	env.currentCube.cube[5] = 0b0101010101010101
	env.solvedCube = env.copyCube(env.currentCube.cube)
}

func (env *Env) shuffle() {
	for step := range env.mix {
		// exec step
		env.execStep(env.mix[step])
	}
}
