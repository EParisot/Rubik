package main

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
