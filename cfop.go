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
	//env.debugPrint(env.currentCube.cube)
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

	//env.debugPrint(env.currentCube.cube)
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
	//env.debugPrint(env.currentCube.cube)
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
	var cornercubie bool
	var sidecubie bool
	for true {
		for y := 0; y < 3; y++ {
			for i := 0; i < 4; i++ {
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[UP]>>24)&15) == FRONT &&
					((env.currentCube.cube[BACK]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("00")
					env.exec(all_algo[0])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == UP &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[UP])&15) == RIGHT &&
					((env.currentCube.cube[LEFT]>>24)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("01")
					env.exec(all_algo[1])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("02")
					env.exec(all_algo[2])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == UP &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>16)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("03")
					env.exec(all_algo[3])
					return
				}

				// 2nd case :
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>4)&15) == RIGHT &&
					((env.currentCube.cube[DOWN]>>20)&15) == DOWN
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("04")
					env.exec(all_algo[4])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>4)&15) == RIGHT &&
					((env.currentCube.cube[DOWN]>>20)&15) == DOWN
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>16)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("05")
					env.exec(all_algo[5])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
					((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("06")
					env.exec(all_algo[6])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
					((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>16)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("07")
					env.exec(all_algo[7])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
					((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>16)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("08")
					env.exec(all_algo[8])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
					((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("09")
					env.exec(all_algo[9])
					return
				}

				//3rd case
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
					((env.currentCube.cube[RIGHT])&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("10")
					env.exec(all_algo[10])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT])&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("11")
					env.exec(all_algo[11])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
					((env.currentCube.cube[RIGHT])&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("12")
					env.exec(all_algo[12])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT])&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("13")
					env.exec(all_algo[13])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
					((env.currentCube.cube[RIGHT])&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("14")
					env.exec(all_algo[14])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT])&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("15")
					env.exec(all_algo[15])
					return
				}

				//4th case

				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>16)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("16")
					env.exec(all_algo[16])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>8)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("17")
					env.exec(all_algo[17])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[UP]>>24)&15) == RIGHT &&
					((env.currentCube.cube[BACK]>>8)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("18")
					env.exec(all_algo[18])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[UP])&15) == FRONT &&
					((env.currentCube.cube[LEFT]>>24)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("19")
					env.exec(all_algo[19])
					return
				}

				//Here devrai se declencher
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[UP])&15) == RIGHT &&
					((env.currentCube.cube[LEFT]>>24)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("20")
					env.exec(all_algo[20])
					return
				}

				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[UP]>>24)&15) == FRONT &&
					((env.currentCube.cube[BACK]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("21")
					env.exec(all_algo[21])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>16)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("22")
					env.exec(all_algo[22])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("23")
					env.exec(all_algo[23])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[UP])&15) == FRONT &&
					((env.currentCube.cube[LEFT]>>24)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("24")
					env.exec(all_algo[24])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[UP]>>24)&15) == RIGHT &&
					((env.currentCube.cube[BACK]>>8)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("25")
					env.exec(all_algo[25])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == DOWN &&
					((env.currentCube.cube[UP]>>12)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>8)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("26")
					env.exec(all_algo[26])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>28)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>12)&15) == FRONT
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>16)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("27")
					env.exec(all_algo[27])
					return
				}

				//5th

				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>8)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("28")
					env.exec(all_algo[28])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>16)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("29")
					env.exec(all_algo[29])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[UP])&15) == FRONT &&
					((env.currentCube.cube[LEFT]>>24)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("30")
					env.exec(all_algo[30])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[UP]>>24)&15) == RIGHT &&
					((env.currentCube.cube[BACK]>>8)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("31")
					env.exec(all_algo[31])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[UP]>>24)&15) == FRONT &&
					((env.currentCube.cube[BACK]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("32")
					env.exec(all_algo[32])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[UP])&15) == RIGHT &&
					((env.currentCube.cube[LEFT]>>24)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("33")
					env.exec(all_algo[33])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[RIGHT]>>24)&15) == RIGHT &&
					((env.currentCube.cube[UP]>>16)&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("34")
					env.exec(all_algo[34])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>20)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>28)&15) == FRONT &&
					((env.currentCube.cube[UP]>>12)&15) == DOWN
				sidecubie = ((env.currentCube.cube[FRONT]>>24)&15) == FRONT &&
					((env.currentCube.cube[UP]>>8)&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("35")
					env.exec(all_algo[35])
					return
				}

				//6th case

				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == FRONT &&
					((env.currentCube.cube[RIGHT]>>4)&15) == RIGHT &&
					((env.currentCube.cube[DOWN]>>20)&15) == DOWN
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT])&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("36")
					env.exec(all_algo[36])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
					((env.currentCube.cube[DOWN]>>20)&15) == FRONT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
					((env.currentCube.cube[RIGHT])&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("37")
					env.exec(all_algo[37])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
					((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == FRONT &&
					((env.currentCube.cube[RIGHT])&15) == RIGHT
				if cornercubie && sidecubie {
					fmt.Println("38")
					env.exec(all_algo[38])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT]>>4)&15) == DOWN &&
					((env.currentCube.cube[DOWN]>>20)&15) == FRONT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT])&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("39")
					env.exec(all_algo[39])
					return
				}
				cornercubie = ((env.currentCube.cube[FRONT]>>12)&15) == DOWN &&
					((env.currentCube.cube[RIGHT]>>4)&15) == FRONT &&
					((env.currentCube.cube[DOWN]>>20)&15) == RIGHT
				sidecubie = ((env.currentCube.cube[FRONT]>>16)&15) == RIGHT &&
					((env.currentCube.cube[RIGHT])&15) == FRONT
				if cornercubie && sidecubie {
					fmt.Println("40")
					env.exec(all_algo[40])
					return
				}
				env.exec("U'")
			}
			if y == 0 {
				env.exec("R U R' U'")
			} else if y == 1 {
				env.exec("R' U R U'")
			} else if y == 2 {
				env.exec("L' U L U'")
			}
		}
	}
}

func (env *Env) cfop() {
	//	env.debugPrint(env.currentCube.cube)
	env.first_cross() // POur gagner beaucoup de coup, possible de faire un A* en - de 10s
	//fmt.Println("Cross :")
	//env.debugPrint(env.currentCube.cube)
	env.res = env.res[:len(env.res)-1]
	fmt.Println(env.res)
	env.f2l()
	//fmt.Println("F2l :")
	env.debugPrint(env.currentCube.cube)
}
