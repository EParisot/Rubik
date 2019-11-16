package main

const RIGHT = 1
const LEFT = 2

//https://ruwix.com/the-rubiks-cube/how-to-solve-the-rubiks-cube-beginners-method/

func (env *Env) arreteintop(face int32, cube [6]int32) int32 {
	if face == ORANGE {
		if (((cube[ORANGE]>>24)&15) == ORANGE &&
			(((cube[WHITE]>>8)&15) == GREEN || ((cube[WHITE]>>8)&15) == BLUE)) ||
			(((cube[WHITE]>>8)&15) == ORANGE &&
				(((cube[ORANGE]>>24)&15) == GREEN || ((cube[ORANGE]>>24)&15) == BLUE)) {
			if (((cube[WHITE] >> 8) & 15) == GREEN) || (((cube[ORANGE] >> 24) & 15) == GREEN) {
				return RIGHT
			}
			return LEFT
		}
		return 0
	} else if face == GREEN {
		if (((cube[GREEN]>>24)&15) == GREEN &&
			(((cube[WHITE]>>8)&15) == RED || ((cube[WHITE]>>8)&15) == ORANGE)) ||
			(((cube[WHITE]>>8)&15) == GREEN &&
				(((cube[GREEN]>>24)&15) == RED || ((cube[GREEN]>>24)&15) == ORANGE)) {
			if (((cube[WHITE] >> 8) & 15) == RED) || (((cube[GREEN] >> 24) & 15) == RED) {
				return RIGHT
			}
			return LEFT
		}
		return 0
	} else if face == RED { //really not sure
		if (((cube[RED]>>24)&15) == RED &&
			(((cube[WHITE]>>8)&15) == BLUE || ((cube[WHITE]>>8)&15) == GREEN)) ||
			(((cube[WHITE]>>8)&15) == RED &&
				(((cube[RED]>>24)&15) == BLUE || ((cube[RED]>>24)&15) == GREEN)) {
			if (((cube[WHITE] >> 8) & 15) == BLUE) || (((cube[RED] >> 24) & 15) == BLUE) {
				return RIGHT
			}
			return LEFT
		}
		return 0
	} else if face == BLUE {
		if (((cube[BLUE]>>24)&15) == BLUE &&
			(((cube[WHITE]>>8)&15) == RED || ((cube[WHITE]>>8)&15) == ORANGE)) ||
			(((cube[WHITE]>>8)&15) == BLUE &&
				(((cube[BLUE]>>24)&15) == RED || ((cube[BLUE]>>24)&15) == ORANGE)) {
			if (((cube[WHITE] >> 8) & 15) == ORANGE) || (((cube[BLUE] >> 24) & 15) == ORANGE) {
				return RIGHT
			}
			return LEFT
		}
		return 0
	}
	return 0
}

func (env *Env) secondlayerisbarelyFinnished(cube [6]int32) bool {

	if ((cube[ORANGE]>>0)&15) == ORANGE || ((cube[ORANGE]>>0)&15) == BLUE &&
		((cube[ORANGE]>>16)&15) == ORANGE || ((cube[ORANGE]>>16)&15) == GREEN &&
		((cube[GREEN]>>0)&15) == GREEN || ((cube[GREEN]>>0)&15) == ORANGE &&
		((cube[GREEN]>>16)&15) == GREEN || ((cube[GREEN]>>16)&15) == RED &&
		((cube[RED]>>0)&15) == RED || ((cube[RED]>>0)&15) == GREEN &&
		((cube[RED]>>16)&15) == RED || ((cube[RED]>>16)&15) == BLUE &&
		((cube[BLUE]>>0)&15) == BLUE || ((cube[BLUE]>>0)&15) == RED &&
		((cube[BLUE]>>16)&15) == BLUE || ((cube[BLUE]>>16)&15) == ORANGE {
		return true
	}
	return false
}
func (env *Env) secondlayer() {
	faceSecondLayer := [4]int32{ORANGE, GREEN, RED, BLUE}

	for true {
		for i := 0; i < 4; i++ {
			for _, slayer := range faceSecondLayer {
				value := env.arreteintop(slayer, env.currentCube.cube)
				if value == RIGHT {
					env.execFace("U R U' R' U' F' U F", slayer)
				} else if value == LEFT {
					env.execFace("U' L' U L U F U' F'", slayer)
				}
			}
			if env.secondlayerisbarelyFinnished(env.currentCube.cube) {
				return
			}
			env.exec("U")
		}
	}
	//regarder si on peux faire l'algo a chacun des coins
	//si on peux pas, tourne U

	//si a la fin on se retrouve avec un qui a la mauvaise orientation, il faut appliquer l'algo (lequel ? ) deux fois
}
