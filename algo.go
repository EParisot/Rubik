package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Cube struct {
}

type Env struct {
	mix  []string
	cube Cube
	res  string
}

func (env *Env) parseArgs(arg string) error {
	steps := strings.Split(arg[0:len(arg)-1], " ")
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
}

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		arg := string(args[0])
		env := Env{}
		err := env.parseArgs(arg)
		if err != nil {
			fmt.Println(err)
		} else {
			env.shuffle()
		}
	} else {
		fmt.Println("Error : No args")
	}
	// testing
	fmt.Println("D D'2 D")
}
