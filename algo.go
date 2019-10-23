package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Env struct {
	mix        []string
	cube       [6][3][3]int
	solvedCube [6][3][3]int
	res        string
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

func (env *Env) rotSide0(cube *[6][3][3]int, way int) {
	sides := [4]int{3, 2, 4, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = mem
	}
}

func (env *Env) rotSide1(cube *[6][3][3]int, way int) {
	sides := [4]int{3, 0, 4, 5}
	if way == 0 {
		mem := cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = cube[sides[2]][2][2]
		cube[sides[2]][2][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = mem
		mem = cube[sides[0]][1][2]
		cube[sides[0]][1][2] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = cube[sides[2]][1][2]
		cube[sides[2]][1][2] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = mem
		mem = cube[sides[0]][0][2]
		cube[sides[0]][0][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = mem
	} else {
		mem := cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = cube[sides[2]][2][2]
		cube[sides[2]][2][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = mem
		mem = cube[sides[0]][1][2]
		cube[sides[0]][1][2] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = cube[sides[2]][1][2]
		cube[sides[2]][1][2] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = mem
		mem = cube[sides[0]][0][2]
		cube[sides[0]][0][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = mem
	}
}

func (env *Env) rotSide2(cube *[6][3][3]int, way int) {
	sides := [4]int{3, 5, 4, 0}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = cube[sides[2]][2][0]
		cube[sides[2]][2][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = mem
		mem = cube[sides[0]][1][0]
		cube[sides[0]][1][0] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = cube[sides[2]][1][0]
		cube[sides[2]][1][0] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = mem
		mem = cube[sides[0]][0][0]
		cube[sides[0]][0][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = cube[sides[2]][2][0]
		cube[sides[2]][2][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = mem
		mem = cube[sides[0]][1][0]
		cube[sides[0]][1][0] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = cube[sides[2]][1][0]
		cube[sides[2]][1][0] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = mem
		mem = cube[sides[0]][0][0]
		cube[sides[0]][0][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = mem
	}
}

func (env *Env) rotSide3(cube *[6][3][3]int, way int) {
	sides := [4]int{5, 2, 0, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][0][1]
		cube[sides[1]][0][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][0][1]
		cube[sides[3]][0][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][0][1]
		cube[sides[3]][0][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][0][1]
		cube[sides[1]][0][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = mem
	}
}

func (env *Env) rotSide4(cube *[6][3][3]int, way int) {
	sides := [4]int{0, 2, 5, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][2][1]
		cube[sides[1]][2][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][2][1]
		cube[sides[3]][2][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][2][1]
		cube[sides[3]][2][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][2][1]
		cube[sides[1]][2][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = mem
	}
}

func (env *Env) rotSide5(cube *[6][3][3]int, way int) {
	sides := [4]int{4, 2, 3, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = mem
	}
}

func (env *Env) rotSides(cube *[6][3][3]int, face, way int) {
	if face == 0 {
		env.rotSide0(cube, way)
	} else if face == 1 {
		env.rotSide1(cube, way)
	} else if face == 2 {
		env.rotSide2(cube, way)
	} else if face == 3 {
		env.rotSide3(cube, way)
	} else if face == 4 {
		env.rotSide4(cube, way)
	} else if face == 5 {
		env.rotSide5(cube, way)
	}
}

func (env *Env) rotate(face, way int) [6][3][3]int {
	var cube [6][3][3]int
	for i := range env.cube {
		for j := range env.cube[i] {
			for k := range env.cube[i][j] {
				cube[i][j][k] = env.cube[i][j][k]
			}
		}
	}
	if way == 0 {
		mem := cube[face][0][0]
		cube[face][0][0] = cube[face][2][0]
		cube[face][2][0] = cube[face][2][2]
		cube[face][2][2] = cube[face][0][2]
		cube[face][0][2] = mem
		mem = cube[face][0][1]
		cube[face][0][1] = cube[face][1][0]
		cube[face][1][0] = cube[face][2][1]
		cube[face][2][1] = cube[face][1][2]
		cube[face][1][2] = mem
	} else {
		mem := cube[face][0][0]
		cube[face][0][0] = cube[face][0][2]
		cube[face][0][2] = cube[face][2][2]
		cube[face][2][2] = cube[face][2][0]
		cube[face][2][0] = mem
		mem = cube[face][0][1]
		cube[face][0][1] = cube[face][1][2]
		cube[face][1][2] = cube[face][2][1]
		cube[face][2][1] = cube[face][1][0]
		cube[face][1][0] = mem
	}
	env.rotSides(&cube, face, way)
	return cube
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
	env.cube = env.rotate(stepID, way)
	if nb == 2 {
		env.cube = env.rotate(stepID, way)
	}
	fmt.Println(env.cube)
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
	//fmt.Println("F F")
}
