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
	env.currentCube.cube[0] = 0b00000000000000000000000000000000 // White
	env.currentCube.cube[1] = 0b00010001000100010001000100010001 // Red
	env.currentCube.cube[2] = 0b00100010001000100010001000100010 // Orange
	env.currentCube.cube[3] = 0b00110011001100110011001100110011 // Blue
	env.currentCube.cube[4] = 0b01000100010001000100010001000100 // Green
	env.currentCube.cube[5] = 0b01010101010101010101010101010101 // Yellow (appear Magenta)
	env.solvedCube = env.copyCube(env.currentCube.cube)
	env.debugPrint("ok", env.currentCube.cube)
}

func (env *Env) shuffle() {
	for step := range env.mix {
		// exec step
		env.execStep(env.mix[step])
	}
}