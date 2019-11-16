package main

import (
//	"fmt"
	"strings"
)

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
	env.currentCube.cube[0] = 0b00000000000000000000000000000000 // Orange
	env.currentCube.cube[1] = 0b00010001000100010001000100010001 // Green
	env.currentCube.cube[2] = 0b00100010001000100010001000100010 // Blue
	env.currentCube.cube[3] = 0b00110011001100110011001100110011 // White
	env.currentCube.cube[4] = 0b01000100010001000100010001000100 // Yellow
	env.currentCube.cube[5] = 0b01010101010101010101010101010101 // Red
	env.solvedCube = env.copyCube(env.currentCube.cube)
}

func (env *Env) shuffle() {
	for step := range env.mix {
		// exec step
		env.execStep(env.mix[step])
	}
}

func (env *Env) exec(str string) {
	okay := strings.Split(str, " ")
	for _, i := range okay {
		// exec step
		env.res += string(i) + " "
		env.execStep(string(i))
	}
}

func (env *Env) execFace(str string, face int32) {
	///fmt.Println("Before :" + str)
	if face == GREEN {
		r := strings.NewReplacer("F", "R",
            "R", "B",
			"B", "L",
			"L", "F")
		str = r.Replace(str)
	} else if face == RED {
		r := strings.NewReplacer("F", "B",
            "R", "L",
			"B", "F",
			"L", "R")
		str = r.Replace(str)
	} else if face == BLUE {
		r := strings.NewReplacer("F", "L",
            "R", "F",
			"B", "R",
			"L", "B")
		str = r.Replace(str)
	}
	//fmt.Println("After :" + str)
	okay := strings.Split(str, " ")
	for _, i := range okay {
		// exec step
		env.res += string(i) + " "
		env.execStep(string(i))
	}
}