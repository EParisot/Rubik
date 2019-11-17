package main

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
