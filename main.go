package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/profile"
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

func parseDoubles(output string) string {
	outTab := strings.Split(output, " ")
	for i := 0; i < len(outTab)-1; i++ {
		if outTab[i] != "" {
			if outTab[i] == outTab[i+1] {
				if !strings.Contains(outTab[i], "2") {
					outTab[i] = outTab[i] + "2"
					outTab[i+1] = ""
				} else {
					outTab[i] = ""
					outTab[i+1] = ""
				}
			} else if outTab[i][0] == outTab[i+1][0] {
				if !strings.Contains(outTab[i], "2") && !strings.Contains(outTab[i+1], "2") {
					if !strings.Contains(outTab[i], "'") && strings.Contains(outTab[i+1], "'") {
						outTab[i] = ""
						outTab[i+1] = ""
					} else if strings.Contains(outTab[i], "'") && !strings.Contains(outTab[i+1], "'") {
						outTab[i] = ""
						outTab[i+1] = ""
					}
				} else if strings.Contains(outTab[i], "2") && strings.Contains(outTab[i+1], "2") {
					if !strings.Contains(outTab[i], "'") && strings.Contains(outTab[i+1], "'") {
						outTab[i] = ""
						outTab[i+1] = ""
					} else if strings.Contains(outTab[i], "'") && !strings.Contains(outTab[i+1], "'") {
						outTab[i] = ""
						outTab[i+1] = ""
					}
				} else if !strings.Contains(outTab[i], "2") && strings.Contains(outTab[i+1], "2") {
					if !strings.Contains(outTab[i], "'") && !strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i] + "'"
						outTab[i+1] = ""
					} else if strings.Contains(outTab[i], "'") && strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i][:1]
						outTab[i+1] = ""
					} else if !strings.Contains(outTab[i], "'") && strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i] + "'"
						outTab[i+1] = ""
					} else if strings.Contains(outTab[i], "'") && !strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i][:1]
						outTab[i+1] = ""
					}
				} else if strings.Contains(outTab[i], "2") && !strings.Contains(outTab[i+1], "2") {
					if !strings.Contains(outTab[i], "'") && !strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i][:1] + "'"
						outTab[i+1] = ""
					} else if strings.Contains(outTab[i], "'") && strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i][:1]
						outTab[i+1] = ""
					} else if !strings.Contains(outTab[i], "'") && strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i][:1]
						outTab[i+1] = ""
					} else if strings.Contains(outTab[i], "'") && !strings.Contains(outTab[i+1], "'") {
						outTab[i] = outTab[i][:1] + "'"
						outTab[i+1] = ""
					}
				}
			}
		}
	}
	output = ""
	for i, step := range outTab {
		if step != "" {
			output += step
			if i < len(outTab)-1 {
				output += " "
			}
		}
	}

	return output
}

func checkOver(output string) bool {
	outTab := strings.Split(output, " ")
	for i := 0; i < len(outTab)-1; i++ {
		if outTab[i][0] == outTab[i+1][0] {
			return false
		}
	}
	return true
}

// TODO improve this
func parseOutput(output string) string {
	output = parseDoubles(output)
	if checkOver(output) == false {
		return parseOutput(output)
	}
	return output
}

func main() {
	var mix string
	var debug bool
	var idaStar bool
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error : No arg")
		return
	}
	for _, arg := range args {
		if arg == "-d" {
			debug = true
		} else if arg == "-ida" {
			idaStar = true
		} else {
			mix = string(arg)
		}
	}
	if debug {
		defer profile.Start(profile.ProfilePath(".")).Stop()
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
	// Solve
	if !idaStar {
		env.beginner()
	} else {
		env.idAstar()
	}
}
