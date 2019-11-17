package main

func (env *Env) PutYellowCubieontheirface() {
	for i := 0; i < 3; i++ {
		if i == 1 {
			env.exec("R' U R F2")
		} else if i == 2 {
			env.exec("L U' L' F2")
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

func (env *Env) firstCross() {
	for i := 0; i < 4; i++ {
		env.PutYellowCubieontheirface()
		if 4 == ((env.currentCube.cube[ORANGE] >> 8) & 15) {
			env.exec("F' R U R' F2") // Algo for reverse cubie
		}
		env.exec("D'")
	}
	// next align yellow aretes cubie every with the correct color of their side
	var CubiefaceOrange bool
	var CubiefaceBlue bool
	var CubiefaceRed bool
	var CubiefaceGreen bool
	for true {
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

	if !CubiefaceOrange && !CubiefaceRed {
		env.exec("F2 U2 B2 U2 F2") // swap ORANGE-RED
	} else if !CubiefaceGreen && !CubiefaceBlue {
		env.exec("R2 U2 L2 U2 R2") // swap GREEN-BLUE
	} else if !CubiefaceOrange && !CubiefaceGreen {
		env.exec("F2 U' R2 U F2") //swap ORANGE-GREEN
	} else if !CubiefaceGreen && !CubiefaceRed {
		env.exec("R2 U' B2 U R2") //swap GREEN-RED
	} else if !CubiefaceRed && !CubiefaceBlue {
		env.exec("B2 U' L2 U B2") //swap RED-BLUE
	} else if !CubiefaceBlue && !CubiefaceOrange {
		env.exec("L2 U' F2 U L2") //swap BLUE-ORANGE
	}
}

func (env *Env) firstPossibleCorner(face int32, cube [6]int32) bool {
	var onecolor bool
	if face == ORANGE {
		onecolor = ((cube[ORANGE]>>20)&15) == YELLOW &&
			((cube[GREEN]>>28)&15) == GREEN &&
			((cube[WHITE]>>12)&15) == ORANGE
	} else if face == GREEN {
		onecolor = ((cube[GREEN]>>20)&15) == YELLOW &&
			((cube[RED]>>12)&15) == RED &&
			((cube[WHITE]>>20)&15) == GREEN
	} else if face == RED {
		onecolor = ((cube[RED]>>4)&15) == YELLOW &&
			((cube[BLUE]>>28)&15) == BLUE &&
			((cube[WHITE]>>28)&15) == RED
	} else if face == BLUE {
		onecolor = ((cube[BLUE]>>20)&15) == YELLOW &&
			((cube[ORANGE]>>28)&15) == ORANGE &&
			((cube[WHITE]>>4)&15) == BLUE
	}
	return onecolor
}

func (env *Env) secondPossibleCorner(face int32, cube [6]int32) bool {
	var twocolor bool
	if face == ORANGE {
		twocolor = ((cube[ORANGE]>>20)&15) == ORANGE &&
			((cube[GREEN]>>28)&15) == YELLOW &&
			((cube[WHITE]>>12)&15) == GREEN
	} else if face == GREEN {
		twocolor = ((cube[GREEN]>>20)&15) == GREEN &&
			((cube[RED]>>12)&15) == YELLOW &&
			((cube[WHITE]>>20)&15) == RED
	} else if face == RED {
		twocolor = ((cube[RED]>>4)&15) == RED &&
			((cube[BLUE]>>28)&15) == YELLOW &&
			((cube[WHITE]>>28)&15) == BLUE
	} else if face == BLUE {
		twocolor = ((cube[BLUE]>>20)&15) == BLUE &&
			((cube[ORANGE]>>28)&15) == YELLOW &&
			((cube[WHITE]>>4)&15) == ORANGE
	}
	return twocolor
}

func (env *Env) thridPossibleCorner(face int32, cube [6]int32) bool {
	var threecolor bool
	if face == ORANGE {
		threecolor = ((cube[ORANGE]>>20)&15) == GREEN &&
			((cube[GREEN]>>28)&15) == ORANGE &&
			((cube[WHITE]>>12)&15) == YELLOW
	} else if face == GREEN {
		threecolor = ((cube[GREEN]>>20)&15) == RED &&
			((cube[RED]>>12)&15) == GREEN &&
			((cube[WHITE]>>20)&15) == YELLOW
	} else if face == RED {
		threecolor = ((cube[RED]>>4)&15) == BLUE &&
			((cube[BLUE]>>28)&15) == RED &&
			((cube[WHITE]>>28)&15) == YELLOW
	} else if face == BLUE {
		threecolor = ((cube[BLUE]>>20)&15) == ORANGE &&
			((cube[ORANGE]>>28)&15) == BLUE &&
			((cube[WHITE]>>4)&15) == YELLOW
	}
	return threecolor
}

func (env *Env) faceFirstLayer(face int32) {
	var first bool
	var second bool
	var third bool

	// Cas orange face :
	for i := 0; i <= 4; i++ {
		for j := 0; j < 4; j++ {
			first = env.firstPossibleCorner(face, env.currentCube.cube)
			second = env.secondPossibleCorner(face, env.currentCube.cube)
			third = env.thridPossibleCorner(face, env.currentCube.cube)
			if first {
				env.execFace("F' U' F", face)
				return
			} else if second {
				env.execFace("R U R'", face)
				return
			} else if third {
				env.execFace("R B U2 B' R'", face)
				return
			} else {
				//sinon change de corner
				env.execFace("U", face)
			}
		}
		// If the corner is block, need to deblock it
		// Warning, ne pas enlever une piece deja mise
		if i == 0 && face == ORANGE {
			env.exec("R U R'")
		} else if i == 1 && (face == ORANGE || face == GREEN) {
			env.exec("R' U R")
		} else if i == 2 && face != BLUE {
			env.exec("L U L'")
		} else if i == 3 {
			env.exec("L' U L")
		}
	}
}
