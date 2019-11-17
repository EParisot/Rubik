package main

func (env *Env) orientedlastcorner() {

	for true {
		one_orange := ((env.currentCube.cube[ORANGE]>>20)&15) == ORANGE &&
			((env.currentCube.cube[GREEN]>>28)&15) == GREEN &&
			((env.currentCube.cube[WHITE]>>12)&15) == WHITE
		if one_orange {
			break
		}
		env.execFace("R' D' R D", ORANGE)
	}
	env.execFace("U", ORANGE)
	for true {
		one_green := ((env.currentCube.cube[ORANGE]>>20)&15) == GREEN &&
			((env.currentCube.cube[GREEN]>>28)&15) == RED &&
			((env.currentCube.cube[WHITE]>>12)&15) == WHITE
		if one_green {
			break
		}
		env.execFace("R' D' R D", ORANGE)
	}
	env.execFace("U", ORANGE)
	for true {
		one_red := ((env.currentCube.cube[ORANGE]>>20)&15) == RED &&
			((env.currentCube.cube[GREEN]>>28)&15) == BLUE &&
			((env.currentCube.cube[WHITE]>>12)&15) == WHITE
		if one_red {
			break
		}
		env.execFace("R' D' R D", ORANGE)
	}
	env.execFace("U", ORANGE)
	for true {
		one_blue := ((env.currentCube.cube[ORANGE]>>20)&15) == BLUE &&
			((env.currentCube.cube[GREEN]>>28)&15) == ORANGE &&
			((env.currentCube.cube[WHITE]>>12)&15) == WHITE
		if one_blue {
			break
		}
		env.execFace("R' D' R D", ORANGE)
	}
	env.execFace("U", ORANGE)
	return
}
