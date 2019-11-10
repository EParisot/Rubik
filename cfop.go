package main

import (
	"fmt"
)

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
				if 4 == ((env.currentCube.cube[4]>>24)&15) || 4 == ((env.currentCube.cube[0]>>8)&15) {
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
		if 4 == ((env.currentCube.cube[0] >> 8) & 15) {
			env.exec("F' R U R' F2") // Algo for reverse cubie
		}
		env.exec("D'")
	}
	env.debugPrint(env.currentCube.cube)
	var CubiefaceOrange bool
	var CubiefaceBlue bool
	var CubiefaceRed bool
	var CubiefaceGreen bool
	for true {
		//env.debugPrint(env.currentCube.cube)
		var s []bool
		CubiefaceOrange = ((env.currentCube.cube[0] >> 8) & 15) == 0
		CubiefaceBlue = ((env.currentCube.cube[2] >> 8) & 15) == 2
		CubiefaceRed = ((env.currentCube.cube[5] >> 24) & 15) == 5
		CubiefaceGreen = ((env.currentCube.cube[1] >> 8) & 15) == 1
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

func (env *Env) cfop() {
	fmt.Println((env.currentCube.cube[5] >> 8) & 15)
	env.debugPrint(env.currentCube.cube)
	env.first_cross() // POur gagner beaucoup de coup, possible de faire un A* en - de 10s
	//env.debugPrint(env.currentCube.cube)
}
