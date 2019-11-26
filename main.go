package main

import (
	"errors"
	"fmt"
	"strconv"
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
			} else if len(outTab[i]) > 0 && len(outTab[i+1]) > 0 && outTab[i][0] == outTab[i+1][0] {
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
		if len(outTab[i]) > 0 && len(outTab[i+1]) > 0 && outTab[i][0] == outTab[i+1][0] {
			return false
		}
	}
	return true
}

func parseOutput(output string) string {
	output = parseDoubles(output)
	if checkOver(output) == false {
		return parseOutput(output)
	}
	return output
}

func main() {
	//var mix string
	var debug bool
	/*var human bool
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("./Rubik[.exe] [-d] [-h] \"R2 D’ B’ D F2 R F2 R2 U L’ F2 U’ B’ L2 R D B’ R’ B2 L2 F2 L2 R2 U2 D2\"")
		return
	}
	for _, arg := range args {
		if arg == "-d" {
			debug = true
		} else if arg == "-h" {
			human = true
		} else {
			mix = string(arg)
		}
	}
	if debug {
		defer profile.Start(profile.ProfilePath(".")).Stop()
	}*/
	env := Env{debug: debug}
	env.setCube()
	/*// parsing
	steps, err := parseArgs(mix)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Shuffling
	env.shuffle(steps)
	// Solve
	if human {
		env.beginner()
	} else {
		env.idAstar()
	}*/
	moves = [6]string{"F", "R", "L", "U", "D", "B"}
	debugCube(env.currentCube.cube)
	buildTableG1(env.currentCube.cube, 0)
	fmt.Println(count)
}

var count int
var resMap map[string]int

func buildTableG1(currCube [6]int32, depth int) {
	if depth < 18 {
		var newCube [6]int32
		for i := 0; i < 6; i++ {
			newCube = rotate(i, 0, currCube)
			newCube = rotate(i, 0, newCube)
			//fmt.Println(moves[i] + "2")
			//fmt.Println(depth)
			//debugCube(newCube)
			var cubeStr string
			for _, val := range newCube {
				cubeStr += strconv.FormatInt(int64(val), 10)
			}
			//fmt.Println(cubeStr)
			if _, ok := resMap[cubeStr]; ok == false {
				resMap[cubeStr] = depth + 1
				count++
				buildTableG1(newCube, depth+1)
			}
		}
	}
}
