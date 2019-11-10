package main

import (
	"fmt"
)

const FRONT = 0 //Orange witout rotation
const RIGHT = 1 //Green witout rotation
const LEFT = 2  //Blue witout rotation
const UP = 3    //White witout rotation
const DOWN = 4  //Yellow witout rotation
const BACK = 5  //Red witout rotation

// 0 1 2
// 7 8 3
// 6 5 4

// https://www.francocube.com/cyril/step_1
//Algo Arrete :
// D'MDM'
// MD2M'
// U'B'EB

func (env *Env) PutYellowCubieontheirface() {
	for i := 0; i < 3; i++ {
		if i == 1 { // bring out back-right edge
			env.exec("R' U R F2")
		} else if i == 2 {
			env.exec("L U' L' F2") //bring out back-left edge
		}
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if 4 == ((env.currentCube.cube[DOWN]>>24)&15) || 4 == ((env.currentCube.cube[FRONT]>>8)&15) {
					return
				}
				env.exec("F")
			}
			env.exec("U")
		}
	}
}

func (env *Env) first_cross() {
	for i := 0; i < 4; i++ {
		env.PutYellowCubieontheirface()
		if 4 == ((env.currentCube.cube[FRONT] >> 8) & 15) {
			env.exec("F' R U R' F2") // Algo for reverse cubie
		}
		env.exec("D'")
	}
	env.debugPrint(env.currentCube.cube)
	// next align yellow aretes cubie every with the correct color of their side
	var CubiefaceOrange bool
	var CubiefaceBlue bool
	var CubiefaceRed bool
	var CubiefaceGreen bool
	for true {
		//env.debugPrint(env.currentCube.cube)
		var s []bool
		CubiefaceOrange = ((env.currentCube.cube[FRONT] >> 8) & 15) == 0
		CubiefaceBlue = ((env.currentCube.cube[LEFT] >> 8) & 15) == 2
		CubiefaceRed = ((env.currentCube.cube[BACK] >> 24) & 15) == 5
		CubiefaceGreen = ((env.currentCube.cube[RIGHT] >> 8) & 15) == 1
		if CubiefaceOrange == true {
			s = append(s, CubiefaceOrange)
		}
		if CubiefaceBlue == true {
			s = append(s, CubiefaceBlue)
		}
		if CubiefaceRed == true {
			s = append(s, CubiefaceRed)
		}
		if CubiefaceGreen == true {
			s = append(s, CubiefaceGreen)
		}
		if len(s) >= 2 {
			break
		}
		env.exec("D")
	}

	env.debugPrint(env.currentCube.cube)
	if !CubiefaceOrange && !CubiefaceRed {
		env.exec("F2 U2 B2 U2 F2") // swap front-back
	} else if !CubiefaceGreen && !CubiefaceBlue {
		fmt.Println("Here")
		env.exec("R2 U2 L2 U2 R2") // swap right-left
	} else if !CubiefaceOrange && !CubiefaceGreen {
		env.exec("F2 U' R2 U F2") //swap front-right
	} else if !CubiefaceGreen && !CubiefaceRed {
		env.exec("R2 U' B2 U R2") //swap right-back
	} else if !CubiefaceRed && !CubiefaceBlue {
		env.exec("B2 U' L2 U B2") //swap back-left
	} else if !CubiefaceBlue && !CubiefaceOrange {
		env.exec("L2 U' F2 U L2") //swap left-front
	}
	env.debugPrint(env.currentCube.cube)
}

func (env *Env) f2l() {
	//https://ruwix.com/the-rubiks-cube/advanced-cfop-fridrich/first-two-layers-f2l/
	all_algo := [42]string{
		//1st case
		"R U R'",
		"F' U' F",
		"U' F' U F",
		"U R U' R'",

		//2nd case
		"U R U' R' U' F' U F",
		"U' F' U F U R U' R'",
		"F' U F U' F' U F",
		"R U R' U' R U R'",
		"R U' R' U R U' R'",
		"F' U' F U F' U' F",

		//3rd case
		"R U R' U' R U R' U' R U R'",
		"R U' R' D R' U R",
		"U F' U F U F' U2 F",
		"U F' U' F D' F U F'",
		"U' R U' R' U' R U2 R'",
		"U' R U R' D R' U' R",

		//4th case
		"R U' R' U D R' U' R",
		"F' U F U' D' F U F'",
		"U F' U2 F U F' U2 F",
		"U' R U2 R' U' R U2 R'",
		"U F' U' F U F' U2 F",
		"U' R U R' U' R U2 R'",
		"U' R U' R' U R U R'",
		"U F' U F U' F' U' F",
		"U' R U R' U R U R'",
		"U F' U' F U' F' U' F",
		"U F' U2 F U' R U R'",
		"U' R U2 R' U F' U' F",

		//5th case
		"R U R' U' U' R U R' U' R U R'",
		"R' U' R U U R' U' R U R' U' R", // why y' au debut ?
		"U2 R U R' U R U' R'",
		"U2 F' U' F U' F' U F",
		"U R U2 R' U R U' R'",
		"U' F' U2 F U' F' U F",
		"R U2 R' U' R U R'",
		"F' U2 F U F' U' F",

		//6th case
		"R U' R' D R' U2 R U R' U2 R",
		"R U' R' U R U2 R' U R U' R'",
		"R U' R' U' R U R' U' R U2 R'",
		"R U R' U' R U' R' U D R' U' R",
		"R U' R' D R' U' R U' R' U' R",
	}

	//1rt case :
	cornercubie := ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie := ((env.currentCube.cube[UP]>>24)&15) == FRONT &&
		((env.currentCube.cube[BACK]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[0])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == UP &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[UP])&15) == RIGHT &&
		((env.currentCube.cube[LEFT]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[1])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[2])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == UP &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[3])
	}

	// 2nd case :
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>4)&15) == RIGHT &&
		((env.currentCube.cube[DOWN]>>20)&15) == DOWN
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[4])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>4)&15) == RIGHT &&
		((env.currentCube.cube[DOWN]>>20)&15) == DOWN
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[5])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
		((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[6])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
		((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[7])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
		((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[8])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
		((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[9])
	}

	//3rd case
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
		((env.currentCube.cube[RIGHT])&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[10])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT])&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[11])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
		((env.currentCube.cube[RIGHT])&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[12])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT])&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[13])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
		((env.currentCube.cube[RIGHT])&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[14])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT])&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[15])
	}

	//4th case

	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>16)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[16])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>8)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[17])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[UP]>>24)&15) == RIGHT &&
		((env.currentCube.cube[BACK]>>8)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[18])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[UP])&15) == FRONT &&
		((env.currentCube.cube[LEFT]>>16)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[19])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[UP])&15) == RIGHT &&
		((env.currentCube.cube[LEFT]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[20])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[UP]>>24)&15) == FRONT &&
		((env.currentCube.cube[BACK]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[21])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[22])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[23])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[UP])&15) == FRONT &&
		((env.currentCube.cube[LEFT]>>16)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[24])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[UP]>>24)&15) == RIGHT &&
		((env.currentCube.cube[BACK]>>8)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[25])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
		((env.currentCube.cube[UP]>>12)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>8)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[26])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>12)&15) == FRONT
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>16)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[27])
	}

	//5th

	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>8)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[28])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>16)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[29])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[UP])&15) == FRONT &&
		((env.currentCube.cube[LEFT]>>16)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[30])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[UP]>>24)&15) == RIGHT &&
		((env.currentCube.cube[BACK]>>8)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[31])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[UP]>>24)&15) == FRONT &&
		((env.currentCube.cube[BACK]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[32])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[UP])&15) == RIGHT &&
		((env.currentCube.cube[LEFT]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[33])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
		((env.currentCube.cube[UP]>>16)&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[34])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
		((env.currentCube.cube[UP]>>12)&15) == DOWN
	sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
		((env.currentCube.cube[UP]>>8)&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[35])
	}

	//6th case

	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == FRONT &&
		((env.currentCube.cube[RIGHT]>>4)&15) == RIGHT &&
		((env.currentCube.cube[DOWN]>>20)&15) == DOWN
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT])&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[36])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
		((env.currentCube.cube[DOWN]>>20)&15) == FRONT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
		((env.currentCube.cube[RIGHT])&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[37])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
		((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
		((env.currentCube.cube[RIGHT])&15) == RIGHT
	if cornercubie && sidecubie {
		env.exec(all_algo[38])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
		((env.currentCube.cube[DOWN]>>20)&15) == FRONT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT])&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[39])
	}
	cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
		((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
		((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
	sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
		((env.currentCube.cube[RIGHT])&15) == FRONT
	if cornercubie && sidecubie {
		env.exec(all_algo[40])
	}

	fmt.Println(all_algo)
	// this function finish the first 2 stages
}

func (env *Env) cfop() {
	fmt.Println((env.currentCube.cube[5] >> 8) & 15)
	env.debugPrint(env.currentCube.cube)
	//env.first_cross() // POur gagner beaucoup de coup, possible de faire un A* en - de 10s
	//env.f2l()

	//env.debugPrint(env.currentCube.cube)
}
