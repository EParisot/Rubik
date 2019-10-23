package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Env struct {
	mix  []string
	cube [6][3][3]int
	res  string
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

func (env *Env) shuffle() {
	/*for step := range env.mix {
		fmt.Println(env.mix[step])
		TODO exec step
	}*/
	//fmt.Print(env.cube)
}

func (env *Env) setCube() {
	for face := range env.cube {
		for line := range env.cube[face] {
			for col := range env.cube[face][line] {
				env.cube[face][line][col] = face
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
		err := env.parseArgs(arg)
		if err != nil {
			fmt.Println(err)
		} else {
			env.shuffle()
		}
	} else {
		fmt.Println("Error : No args")
	}
	// TEST
	fmt.Println("D D'2 D")
}
