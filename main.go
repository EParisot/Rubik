package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Game environnement
type Env struct {
	mix        []string     //shuffling list
	cube       [6][3][3]int //current cube
	solvedCube [6][3][3]int //finished cube (const)
	res        string       //result list
}

func (env *Env) parseArgs(arg string) error {
	arg = strings.Replace(arg, "\n", "", -1)
	if arg[len(arg)-1] == ' ' {
		arg = arg[0 : len(arg)-1]
	}
	steps := strings.Split(arg, " ")
	for step := range steps {
		if len(steps[step]) == 0 || len(steps[step]) > 3 || len(steps[step]) > 0 &&
			!strings.Contains("FRUBLD", steps[step][0:1]) {
			return errors.New("Error : Invalid step name")
		} else if len(steps[step]) == 2 && !strings.Contains("'’2", steps[step][1:2]) {
			return errors.New("Error : Invalid step arg")
		} else if len(steps[step]) == 3 &&
			(!strings.Contains("'’2", steps[step][1:2]) || !strings.Contains("2", steps[step][2:3])) {
			return errors.New("Error : Invalid step arg")
		}
	}
	env.mix = steps
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
	env.cube = env.rotate(stepID, way)
	if nb == 2 {
		env.cube = env.rotate(stepID, way)
	}
	// DEBUG
	//fmt.Println(env.cube)
}

func (env *Env) shuffle() {
	for step := range env.mix {
		// exec step
		env.execStep(env.mix[step])
	}
}

func (env *Env) setCube() {
	i := 0
	for face := range env.cube {
		for line := range env.cube[face] {
			for col := range env.cube[face][line] {
				env.cube[face][line][col] = i
				env.solvedCube[face][line][col] = i
				i++
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		arg := string(args[0])
		env := Env{}
		env.setCube()
		// parsing
		err := env.parseArgs(arg)
		if err != nil {
			fmt.Println(err)
		} else {
			// Shuffling
			env.shuffle()
			// Solve HERE
		}
	} else {
		fmt.Println("Error : No args")
	}
	// TEST
	fmt.Println("U U'2 U")
}
