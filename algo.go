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
	steps := strings.Split(arg, " ")
	for step := range steps {
		if len(steps[step]) == 0 || len(steps[step]) > 3 ||
			!strings.Contains("FRUBLD", steps[step][:1]) {
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

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		arg := string(args[0])
		env := Env{}
		err := env.parseArgs(arg)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error : No args")
	}
}
