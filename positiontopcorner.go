package main

func (env *Env) orangecorner() bool {
	one_orange := ((env.currentCube.cube[ORANGE]>>20)&15) == ORANGE &&
		((env.currentCube.cube[GREEN]>>28)&15) == GREEN &&
		((env.currentCube.cube[WHITE]>>12)&15) == WHITE
	two_orange := ((env.currentCube.cube[ORANGE]>>20)&15) == WHITE &&
		((env.currentCube.cube[GREEN]>>28)&15) == ORANGE &&
		((env.currentCube.cube[WHITE]>>12)&15) == GREEN
	three_orange := ((env.currentCube.cube[ORANGE]>>20)&15) == GREEN &&
		((env.currentCube.cube[GREEN]>>28)&15) == WHITE &&
		((env.currentCube.cube[WHITE]>>12)&15) == ORANGE
	if one_orange || two_orange || three_orange {
		return true
	}
	return false
}

func (env *Env) greencorner() bool {
	one_green := ((env.currentCube.cube[GREEN]>>20)&15) == GREEN &&
		((env.currentCube.cube[RED]>>12)&15) == RED &&
		((env.currentCube.cube[WHITE]>>20)&15) == WHITE
	two_green := ((env.currentCube.cube[GREEN]>>20)&15) == WHITE &&
		((env.currentCube.cube[RED]>>12)&15) == GREEN &&
		((env.currentCube.cube[WHITE]>>20)&15) == RED
	three_green := ((env.currentCube.cube[GREEN]>>20)&15) == RED &&
		((env.currentCube.cube[RED]>>12)&15) == WHITE &&
		((env.currentCube.cube[WHITE]>>20)&15) == GREEN
	if one_green || two_green || three_green {
		return true
	}
	return false
}

func (env *Env) redcorner() bool {
	one_red := ((env.currentCube.cube[RED]>>4)&15) == RED &&
		((env.currentCube.cube[BLUE]>>28)&15) == BLUE &&
		((env.currentCube.cube[WHITE]>>28)&15) == WHITE
	two_red := ((env.currentCube.cube[RED]>>4)&15) == WHITE &&
		((env.currentCube.cube[BLUE]>>28)&15) == RED &&
		((env.currentCube.cube[WHITE]>>28)&15) == BLUE
	three_red := ((env.currentCube.cube[RED]>>4)&15) == BLUE &&
		((env.currentCube.cube[BLUE]>>28)&15) == WHITE &&
		((env.currentCube.cube[WHITE]>>28)&15) == RED
	if one_red || two_red || three_red {
		return true
	}
	return false
}

func (env *Env) bluecorner() bool {
	one_blue := ((env.currentCube.cube[BLUE]>>20)&15) == BLUE &&
		((env.currentCube.cube[ORANGE]>>28)&15) == ORANGE &&
		((env.currentCube.cube[WHITE]>>4)&15) == WHITE
	two_blue := ((env.currentCube.cube[BLUE]>>20)&15) == WHITE &&
		((env.currentCube.cube[ORANGE]>>28)&15) == BLUE &&
		((env.currentCube.cube[WHITE]>>4)&15) == ORANGE
	three_blue := ((env.currentCube.cube[BLUE]>>20)&15) == ORANGE &&
		((env.currentCube.cube[ORANGE]>>28)&15) == WHITE &&
		((env.currentCube.cube[WHITE]>>4)&15) == BLUE
	if one_blue || two_blue || three_blue {
		return true
	}
	return false
}

func (env *Env) positiontopcorner() {
	for true {
		if env.orangecorner() {
			for true {
				if env.orangecorner() && env.greencorner() && env.redcorner() && env.bluecorner() {
					return
				}
				env.execFace("U R U' L' U R' U' L", ORANGE)
			}
		}
		if env.greencorner() {
			for true {
				if env.orangecorner() && env.greencorner() && env.redcorner() && env.bluecorner() {
					return
				}
				env.execFace("U R U' L' U R' U' L", GREEN)
			}
		}
		if env.redcorner() {
			for true {
				if env.orangecorner() && env.greencorner() && env.redcorner() && env.bluecorner() {
					return
				}
				env.execFace("U R U' L' U R' U' L", RED)
			}
		}
		if env.bluecorner() {
			for true {
				if env.orangecorner() && env.greencorner() && env.redcorner() && env.bluecorner() {
					return
				}
				env.execFace("U R U' L' U R' U' L", BLUE)
			}
		}
		env.execFace("U R U' L' U R' U' L", ORANGE)
		// if good color, pass
		//else if can inverse with left one , do it (if it good for both)
	}
}

//Chercher un bon corner (ou ya les bonnes couleurs, peux importe l'orientation)
//Si trouver, faire l'algo autant de fois que tout les cubes se mettent en place
//Sinon faire l'ago au pif pui rechercher un bon corner
