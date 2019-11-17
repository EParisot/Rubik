package main

func (env *Env) orientedlastcorner() {
	facetop := [4]int32{ORANGE, GREEN, RED, BLUE}
	righttop := [4]int32{GREEN, RED, BLUE, ORANGE}
	for i := 0; i < 4; i++ {
		for true {
			corner := ((env.currentCube.cube[ORANGE]>>20)&15) == facetop[i] &&
				((env.currentCube.cube[GREEN]>>28)&15) == righttop[i] &&
				((env.currentCube.cube[WHITE]>>12)&15) == WHITE
			if corner {
				break
			}
			env.execFace("R' D' R D", ORANGE)
		}
		env.execFace("U", ORANGE)
	}
	return
}
