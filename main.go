package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type CubeEnv struct {
	cube              [6]int32
	cost              int
	heuristic         int
	internationalMove string
}

// Game environnement
type Env struct {
	mix         []string //shuffling list
	currentCube CubeEnv  //current cube
	solvedCube  [6]int32 //finished cube (const)
	res         string   //result list
}

func (env *Env) parseArgs(arg string) error {
	arg = strings.Replace(arg, "\n", "", -1)
	if len(arg) != 0 {
		if arg[len(arg)-1] == ' ' {
			arg = arg[0 : len(arg)-1]
		}
		steps := strings.Split(arg, " ")
		for step := range steps {
			steps[step] = strings.ReplaceAll(steps[step], "’", "'")
			if len(steps[step]) == 0 || len(steps[step]) > 3 || (len(steps[step]) > 0 &&
				!strings.Contains("FRUBLD", steps[step][0:1])) {
				return errors.New("Error : Invalid step name")
			} else if len(steps[step]) == 2 && !strings.Contains("'2", steps[step][1:2]) {
				return errors.New("Error : Invalid step arg")
			} else if len(steps[step]) == 3 &&
				(!strings.Contains("'2", steps[step][1:2]) || !strings.Contains("2", steps[step][2:3])) {
				return errors.New("Error : Invalid step arg")
			}
		}
		env.mix = steps
	} else {
		return errors.New("Error : No arg")
	}
	return nil
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
	env.currentCube.cube = env.rotate(stepID, way, oldCube)
	if nb == 2 {
		env.currentCube.cube = env.rotate(stepID, way, env.currentCube.cube)
	}
	// DEBUG
	//	env.debugPrint(env.currentCube.cube)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error : No arg")
		return
	}
	arg := string(args[0])
	env := Env{}
	env.setCube()
	// parsing
	err := env.parseArgs(arg)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Shuffling
	env.shuffle()
	// Solve HERE
	//env.idAstar()
	env.cfop()
}
