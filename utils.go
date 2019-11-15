package main

var solvedCube [6]int32

func isFinished(currCube CubeEnv) bool {
	if currCube.cube == solvedCube {
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
	env.currentCube.cube[0] = 0b00000000000000000000000000000000 // Orange (appear Magenta)
	env.currentCube.cube[1] = 0b00010001000100010001000100010001 // Green
	env.currentCube.cube[2] = 0b00100010001000100010001000100010 // Blue
	env.currentCube.cube[3] = 0b00110011001100110011001100110011 // White
	env.currentCube.cube[4] = 0b01000100010001000100010001000100 // Yellow
	env.currentCube.cube[5] = 0b01010101010101010101010101010101 // Red
	solvedCube = env.currentCube.cube
}

func (env *Env) shuffle(steps []string) {
	for step := range steps {
		// exec step
		env.execStep(steps[step])
	}
	env.startCube = env.currentCube
}
