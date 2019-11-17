package main

import (
	//	"fmt"
	"math/rand"
	"strings"
)

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

func (env *Env) generaterandom() string {

	var str string
	possible := [5]string{"U", "R", "L", "D", "B"}
	var reverse int

	for i := 0; i < 25; i++ {
		str = str + possible[rand.Intn(5)]
		reverse = rand.Intn(2)
		if reverse == 1 {
			str = str + "'"
		}
		str = str + " "
	}
	str = str[:len(str)-1]
	return str
}

func (env *Env) shuffle(steps []string) {
	for step := range steps {
		// exec step
		env.execStep(steps[step])
	}
	env.startCube = env.currentCube
}

func (env *Env) execStep(step string) {
	stepID := 0
	way := 0
	nb := 1
	if len(step) == 2 {
		if string(step[1]) == "'" || string(step[1]) == "’" {
			way = 1
		} else if string(step[1]) == "2" {
			nb = 2
		}
	} else if len(step) == 3 {
		if string(step[1]) == "'" || string(step[1]) == "’" {
			way = 1
		}
		if string(step[2]) == "2" {
			nb = 2
		}
	}
	if step[0] == 'F' {
		stepID = 0
	} else if step[0] == 'R' {
		stepID = 1
	} else if step[0] == 'U' {
		stepID = 3
	} else if step[0] == 'B' {
		stepID = 5
	} else if step[0] == 'L' {
		stepID = 2
	} else if step[0] == 'D' {
		stepID = 4
	}
	// exec rotations
	oldCube := env.currentCube.cube
	env.currentCube.cube = rotate(stepID, way, oldCube)
	if nb == 2 {
		env.currentCube.cube = rotate(stepID, way, env.currentCube.cube)
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
