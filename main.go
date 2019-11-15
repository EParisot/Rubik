package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// CubeEnv : Cube representation
type CubeEnv struct {
	cube              [6]int32
	cost              int
	heuristic         int
	internationalMove string
}

// Env : Game environnement
type Env struct {
	startCube   CubeEnv //start cube
	currentCube CubeEnv //current cube
	res         string  //result list
	debug       bool
}

func parseArgs(arg string) ([]string, error) {
	var steps []string
	arg = strings.Replace(arg, "\n", "", -1)
	if len(arg) != 0 {
		if arg[len(arg)-1] == ' ' {
			arg = arg[0 : len(arg)-1]
		}
		steps = strings.Split(arg, " ")
		for step := range steps {
			steps[step] = strings.ReplaceAll(steps[step], "’", "'")
			if len(steps[step]) == 0 || len(steps[step]) > 3 || (len(steps[step]) > 0 &&
				!strings.Contains("FRUBLD", steps[step][0:1])) {
				return nil, errors.New("Error : Invalid step name")
			} else if len(steps[step]) == 2 && !strings.Contains("'2", steps[step][1:2]) {
				return nil, errors.New("Error : Invalid step arg")
			} else if len(steps[step]) == 3 &&
				(!strings.Contains("'2", steps[step][1:2]) || !strings.Contains("2", steps[step][2:3])) {
				return nil, errors.New("Error : Invalid step arg")
			}
		}
	} else {
		return nil, errors.New("Error : No arg")
	}
	return steps, nil
}

func parseOutput(rawOutput string) string {
	var output string
	outTab := strings.Split(rawOutput, " ")
	for i := 0; i < len(outTab); i++ {
		if i < len(outTab)-1 && outTab[i] == outTab[i+1] {
			if strings.Contains(outTab[i], "2") {
				i++
				continue
			} else {
				i++
				outTab[i] += "2"
			}
		}
		output += outTab[i]
		if i < len(outTab)-1 {
			output += " "
		}
	}
	return output
}

func main() {
	var mix string
	var debug bool
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error : No arg")
		return
	}
	if len(args) == 2 {
		if args[0] == "-d" {
			debug = true
			mix = string(args[1])
		} else if args[1] == "-d" {
			debug = true
			mix = string(args[0])
		}
	} else {
		debug = false
		mix = string(args[0])
	}
	env := Env{debug: debug}
	env.setCube()
	// parsing
	steps, err := parseArgs(mix)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Shuffling
	env.shuffle(steps)
	// Solve HERE
	env.idAstar()
}
