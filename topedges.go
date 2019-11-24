package main

func (env *Env) topedgesorangeisPlaced() bool {
	if ((env.currentCube.cube[ORANGE] >> 24) & 15) == ORANGE {
		return true
	}
	return false
}

func (env *Env) topedgesgreenisPlaced() bool {
	if ((env.currentCube.cube[GREEN] >> 24) & 15) == GREEN {
		return true
	}
	return false
}

func (env *Env) topedgesredisPlaced() bool {
	if ((env.currentCube.cube[RED] >> 8) & 15) == RED {
		return true
	}
	return false
}

func (env *Env) topedgesblueisPlaced() bool {
	if ((env.currentCube.cube[BLUE] >> 24) & 15) == BLUE {
		return true
	}
	return false
}

func (env *Env) topedgesisFinnished() bool {
	if ((env.currentCube.cube[ORANGE]>>24)&15) == ORANGE &&
		((env.currentCube.cube[GREEN]>>24)&15) == GREEN &&
		((env.currentCube.cube[RED]>>8)&15) == RED &&
		((env.currentCube.cube[BLUE]>>24)&15) == BLUE {
		return true
	}
	return false
}

func (env *Env) topedges() {
	// try to place maximum before
	for i := 0; i < 4; i++ {
		if env.topedgesorangeisPlaced() && (env.topedgesgreenisPlaced() || env.topedgesredisPlaced() || env.topedgesblueisPlaced()) {
			break
		}
		if env.topedgesgreenisPlaced() && (env.topedgesredisPlaced() || env.topedgesblueisPlaced()) {
			break
		}
		if env.topedgesredisPlaced() && env.topedgesblueisPlaced() {
			break
		}
		env.exec("U")
	}
	for true {
		for i := 0; i < 4; i++ {
			if ((env.currentCube.cube[ORANGE] >> 24) & 15) != ORANGE {
				env.execFace("R U R' U R U2 R' U", ORANGE)
			}
			if ((env.currentCube.cube[GREEN] >> 24) & 15) != GREEN {
				env.execFace("R U R' U R U2 R' U", GREEN)
			}
			if ((env.currentCube.cube[RED] >> 8) & 15) != RED {
				env.execFace("R U R' U R U2 R' U", RED)
			}
			if ((env.currentCube.cube[BLUE] >> 24) & 15) != BLUE {
				env.execFace("R U R' U R U2 R' U", BLUE)
			}
			if env.topedgesisFinnished() {
				return
			}
		}
	}
}



func (env *Env) topcrossisFinnished(cube [6]int32) bool {
	if ((cube[WHITE]>>8)&15) == WHITE &&
		((cube[WHITE]>>16)&15) == WHITE &&
		((cube[WHITE]>>24)&15) == WHITE &&
		((cube[WHITE]>>0)&15) == WHITE {
		return true
	}
	return false
}

func (env *Env) topcross() {
	for true {
		if ((env.currentCube.cube[ORANGE] >> 24) & 15) == WHITE {
			env.execFace("F R U R' U' F'", ORANGE)
		}
		if ((env.currentCube.cube[GREEN] >> 24) & 15) == WHITE {
			env.execFace("F R U R' U' F'", GREEN)
		}
		if ((env.currentCube.cube[RED] >> 8) & 15) == WHITE {
			env.execFace("F R U R' U' F'", RED)
		}
		if ((env.currentCube.cube[BLUE] >> 24) & 15) == WHITE {
			env.execFace("F R U R' U' F'", BLUE)
		}
		if env.topcrossisFinnished(env.currentCube.cube) {
			return
		}
	}
}
