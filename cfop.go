package main

import (
	"fmt"
)

const ORANGE = 0 //Orange witout rotation
const GREEN = 1  //Green witout rotation
const BLUE = 2   //Blue witout rotation
const WHITE = 3  //White witout rotation
const YELLOW = 4 //Yellow witout rotation
const RED = 5    //Red witout rotation

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
				if 4 == ((env.currentCube.cube[YELLOW]>>24)&15) || 4 == ((env.currentCube.cube[ORANGE]>>8)&15) {
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
		if 4 == ((env.currentCube.cube[ORANGE] >> 8) & 15) {
			env.exec("F' R U R' F2") // Algo for reverse cubie
		}
		env.exec("D'")
	}
	//env.debugPrint(env.currentCube.cube)
	// next align yellow aretes cubie every with the correct color of their side
	var CubiefaceOrange bool
	var CubiefaceBlue bool
	var CubiefaceRed bool
	var CubiefaceGreen bool
	for true {
		//env.debugPrint(env.currentCube.cube)
		var s []bool
		CubiefaceOrange = ((env.currentCube.cube[ORANGE] >> 8) & 15) == 0
		CubiefaceBlue = ((env.currentCube.cube[BLUE] >> 8) & 15) == 2
		CubiefaceRed = ((env.currentCube.cube[RED] >> 24) & 15) == 5
		CubiefaceGreen = ((env.currentCube.cube[GREEN] >> 8) & 15) == 1
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

	//env.debugPrint(env.currentCube.cube)
	if !CubiefaceOrange && !CubiefaceRed {
		env.exec("F2 U2 B2 U2 F2") // swap ORANGE-back
	} else if !CubiefaceGreen && !CubiefaceBlue {
		env.exec("R2 U2 L2 U2 R2") // swap right-left
	} else if !CubiefaceOrange && !CubiefaceGreen {
		env.exec("F2 U' R2 U F2") //swap ORANGE-right
	} else if !CubiefaceGreen && !CubiefaceRed {
		env.exec("R2 U' B2 U R2") //swap right-back
	} else if !CubiefaceRed && !CubiefaceBlue {
		env.exec("B2 U' L2 U B2") //swap back-left
	} else if !CubiefaceBlue && !CubiefaceOrange {
		env.exec("L2 U' F2 U L2") //swap left-ORANGE
	}
	//env.debugPrint(env.currentCube.cube)
}

func (env *Env) firstlayer() {
	//first objective
	var first bool
	top_orange_one = ((env.currentCube.cube[ORANGE]>>20)&15) == GREEN &&
	((env.currentCube.cube[GREEN]>>28)&15) == ORANGE &&
	((env.currentCube.cube[WHITE]>>12)&15) == YELLOW
	top_orange_two = ((env.currentCube.cube[ORANGE]>>20)&15) == GREEN &&
	((env.currentCube.cube[GREEN]>>28)&15) == ORANGE &&
	((env.currentCube.cube[WHITE]>>12)&15) == YELLOW
	top_orange_one = ((env.currentCube.cube[ORANGE]>>20)&15) == GREEN &&
	((env.currentCube.cube[GREEN]>>28)&15) == ORANGE &&
	((env.currentCube.cube[WHITE]>>12)&15) == YELLOW 
	//test x 4
	//Si c'est bon mais juste inverse algo
	for true {
		top_orange =
		if first {
			env.exec("R' D' R D")
			return
		}
	}
}

func (env *Env) cfop() {

	env.first_cross() // POur gagner beaucoup de coup, possible de faire un A* en - de 10s
	env.firstlayer()
	env.res = env.res[:len(env.res)-1]
	fmt.Println(env.res)

}
