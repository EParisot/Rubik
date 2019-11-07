package main

func replaceBits(src int32, sLoc int, dest int32, dLoc int) int32 {
	mask := -int32(0b1111 << dLoc)-1
	value := (src >> sLoc) & 0b1111
	return (dest & mask) | (value<<dLoc)
}

func (env *Env) rotSide0(cube *[6]int32, way int) {
	sides := [4]int{3, 2, 4, 1}
	mem := cube[sides[0]]
	if way == 0 {
		//cube[sides[0]][2][0] = cube[sides[1]][2][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 12, cube[sides[0]], 4)
		//cube[sides[1]][2][2] = cube[sides[2]][0][2]
		cube[sides[1]] = replaceBits(cube[sides[2]], 20, cube[sides[1]], 12)
		//cube[sides[2]][0][2] = cube[sides[3]][0][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 28, cube[sides[2]], 20)
		//cube[sides[3]][0][0] = mem
		cube[sides[3]] = replaceBits(mem, 4, cube[sides[3]], 28)

		//cube[sides[0]][2][1] = cube[sides[1]][1][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 16, cube[sides[0]], 8)
		//cube[sides[1]][1][2] = cube[sides[2]][0][1]
		cube[sides[1]] = replaceBits(cube[sides[2]], 24, cube[sides[1]], 16)
		//cube[sides[2]][0][1] = cube[sides[3]][1][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 0, cube[sides[2]], 24)
		//cube[sides[3]][1][0] = mem
		cube[sides[3]] = replaceBits(mem, 8, cube[sides[3]], 0)

		//cube[sides[0]][2][2] = cube[sides[1]][0][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 20, cube[sides[0]], 12)
		//cube[sides[1]][0][2] = cube[sides[2]][0][0]
		cube[sides[1]] = replaceBits(cube[sides[2]], 28, cube[sides[1]], 20)
		//cube[sides[2]][0][0] = cube[sides[3]][2][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 4, cube[sides[2]], 28)
		//cube[sides[3]][2][0] = mem
		cube[sides[3]] = replaceBits(mem, 12, cube[sides[3]], 4)
	} else {
		//cube[sides[0]][2][0] = cube[sides[3]][0][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 28, cube[sides[0]], 4)
		//cube[sides[3]][0][0] = cube[sides[2]][0][2]
		cube[sides[3]] = replaceBits(cube[sides[2]], 20, cube[sides[3]], 28)
		//cube[sides[2]][0][2] = cube[sides[1]][2][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 12, cube[sides[2]], 20)
		//cube[sides[1]][2][2] = mem
		cube[sides[1]] = replaceBits(mem, 4, cube[sides[1]], 12)

		//cube[sides[0]][2][1] = cube[sides[3]][1][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 0, cube[sides[0]], 8)
		//cube[sides[3]][1][0] = cube[sides[2]][0][1]
		cube[sides[3]] = replaceBits(cube[sides[2]], 24, cube[sides[3]], 0)
		//cube[sides[2]][0][1] = cube[sides[1]][1][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 16, cube[sides[2]], 24)
		//cube[sides[1]][1][2] = mem
		cube[sides[1]] = replaceBits(mem, 8, cube[sides[1]], 16)

		//cube[sides[0]][2][2] = cube[sides[3]][2][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 4, cube[sides[0]], 12)
		//cube[sides[3]][2][0] = cube[sides[2]][0][0]
		cube[sides[3]] = replaceBits(cube[sides[2]], 28, cube[sides[3]], 4)
		//cube[sides[2]][0][0] = cube[sides[1]][0][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 20, cube[sides[2]], 28)
		//cube[sides[1]][0][2] = mem
		cube[sides[1]] = replaceBits(mem, 12, cube[sides[1]], 20)
	}
}

func (env *Env) rotSide1(cube *[6]int32, way int) {
	sides := [4]int{3, 0, 4, 5}
	mem := cube[sides[0]]
	if way == 0 {
		//cube[sides[0]][2][2] = cube[sides[1]][2][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 12, cube[sides[0]], 12)
		//cube[sides[1]][2][2] = cube[sides[2]][2][2]
		cube[sides[1]] = replaceBits(cube[sides[2]], 12, cube[sides[1]], 12)
		//cube[sides[2]][2][2] = cube[sides[3]][2][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 12, cube[sides[2]], 12)
		//cube[sides[3]][2][2] = mem
		cube[sides[3]] = replaceBits(mem, 12, cube[sides[3]], 12)

		//cube[sides[0]][1][2] = cube[sides[1]][1][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 16, cube[sides[0]], 16)
		//cube[sides[1]][1][2] = cube[sides[2]][1][2]
		cube[sides[1]] = replaceBits(cube[sides[2]], 16, cube[sides[1]], 16)
		//cube[sides[2]][1][2] = cube[sides[3]][1][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 16, cube[sides[2]], 16)
		//cube[sides[3]][1][2] = mem
		cube[sides[3]] = replaceBits(mem, 16, cube[sides[3]], 16)

		//cube[sides[0]][0][2] = cube[sides[1]][0][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 20, cube[sides[0]], 20)
		//cube[sides[1]][0][2] = cube[sides[2]][0][2]
		cube[sides[1]] = replaceBits(cube[sides[2]], 20, cube[sides[1]], 20)
		//cube[sides[2]][0][2] = cube[sides[3]][0][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 20, cube[sides[2]], 20)
		//cube[sides[3]][0][2] = mem
		cube[sides[3]] = replaceBits(mem, 20, cube[sides[3]], 20)
	} else {
		//cube[sides[0]][2][2] = cube[sides[3]][2][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 12, cube[sides[0]], 12)
		//cube[sides[3]][2][2] = cube[sides[2]][2][2]
		cube[sides[3]] = replaceBits(cube[sides[2]], 12, cube[sides[3]], 12)
		//cube[sides[2]][2][2] = cube[sides[1]][2][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 12, cube[sides[2]], 12)
		//cube[sides[1]][2][2] = mem
		cube[sides[1]] = replaceBits(mem, 12, cube[sides[1]], 12)

		//cube[sides[0]][1][2] = cube[sides[3]][1][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 16, cube[sides[0]], 16)
		//cube[sides[3]][1][2] = cube[sides[2]][1][2]
		cube[sides[3]] = replaceBits(cube[sides[2]], 16, cube[sides[3]], 16)
		//cube[sides[2]][1][2] = cube[sides[1]][1][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 16, cube[sides[2]], 16)
		//cube[sides[1]][1][2] = mem
		cube[sides[1]] = replaceBits(mem, 16, cube[sides[1]], 16)

		//cube[sides[0]][0][2] = cube[sides[3]][0][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 20, cube[sides[0]], 20)
		//cube[sides[3]][0][2] = cube[sides[2]][0][2]
		cube[sides[3]] = replaceBits(cube[sides[2]], 20, cube[sides[3]], 20)
		//cube[sides[2]][0][2] = cube[sides[1]][0][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 20, cube[sides[2]], 20)
		//cube[sides[1]][0][2] = mem
		cube[sides[1]] = replaceBits(mem, 20, cube[sides[1]], 20)
	}
}

func (env *Env) rotSide2(cube *[6]int32, way int) {
	sides := [4]int{3, 5, 4, 0}
	mem := cube[sides[0]]
	if way == 0 {
		//cube[sides[0]][2][0] = cube[sides[1]][2][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 4, cube[sides[0]], 4)
		//cube[sides[1]][2][0] = cube[sides[2]][2][0]
		cube[sides[1]] = replaceBits(cube[sides[2]], 4, cube[sides[1]], 4)
		//cube[sides[2]][2][0] = cube[sides[3]][2][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 4, cube[sides[2]], 4)
		//cube[sides[3]][2][0] = mem
		cube[sides[3]] = replaceBits(mem, 4, cube[sides[3]], 4)

		//cube[sides[0]][1][0] = cube[sides[1]][1][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 0, cube[sides[0]], 0)
		//cube[sides[1]][1][0] = cube[sides[2]][1][0]
		cube[sides[1]] = replaceBits(cube[sides[2]], 0, cube[sides[1]], 0)
		//cube[sides[2]][1][0] = cube[sides[3]][1][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 0, cube[sides[2]], 0)
		//cube[sides[3]][1][0] = mem
		cube[sides[3]] = replaceBits(mem, 0, cube[sides[3]], 0)

		//cube[sides[0]][0][0] = cube[sides[1]][0][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 28, cube[sides[0]], 28)
		//cube[sides[1]][0][0] = cube[sides[2]][0][0]
		cube[sides[1]] = replaceBits(cube[sides[2]], 28, cube[sides[1]], 28)
		//cube[sides[2]][0][0] = cube[sides[3]][0][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 28, cube[sides[2]], 28)
		//cube[sides[3]][0][0] = mem
		cube[sides[3]] = replaceBits(mem, 28, cube[sides[3]], 28)
	} else {
		//cube[sides[0]][2][0] = cube[sides[3]][2][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 4, cube[sides[0]], 4)
		//cube[sides[3]][2][0] = cube[sides[2]][2][0]
		cube[sides[3]] = replaceBits(cube[sides[2]], 4, cube[sides[3]], 4)
		//cube[sides[2]][2][0] = cube[sides[1]][2][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 4, cube[sides[2]], 4)
		//cube[sides[1]][2][0] = mem
		cube[sides[1]] = replaceBits(mem, 4, cube[sides[1]], 4)

		//cube[sides[0]][1][0] = cube[sides[3]][1][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 0, cube[sides[0]], 0)
		//cube[sides[3]][1][0] = cube[sides[2]][1][0]
		cube[sides[3]] = replaceBits(cube[sides[2]], 0, cube[sides[3]], 0)
		//cube[sides[2]][1][0] = cube[sides[1]][1][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 0, cube[sides[2]], 0)
		//cube[sides[1]][1][0] = mem
		cube[sides[1]] = replaceBits(mem, 0, cube[sides[1]], 0)

		//cube[sides[0]][0][0] = cube[sides[3]][0][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 28, cube[sides[0]], 28)
		//cube[sides[3]][0][0] = cube[sides[2]][0][0]
		cube[sides[3]] = replaceBits(cube[sides[2]], 28, cube[sides[3]], 28)
		//cube[sides[2]][0][0] = cube[sides[1]][0][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 28, cube[sides[2]], 28)
		//cube[sides[1]][0][0] = mem
		cube[sides[1]] = replaceBits(mem, 28, cube[sides[1]], 28)
	}
}

func (env *Env) rotSide3(cube *[6]int32, way int) {
	sides := [4]int{5, 2, 0, 1}
	mem := cube[sides[0]]
	if way == 0 {
		//cube[sides[0]][2][0] = cube[sides[1]][0][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 20, cube[sides[0]], 4)
		//cube[sides[1]][0][2] = cube[sides[2]][0][2]
		cube[sides[1]] = replaceBits(cube[sides[2]], 20, cube[sides[1]], 20)
		//cube[sides[2]][0][2] = cube[sides[3]][0][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 20, cube[sides[2]], 20)
		//cube[sides[3]][0][2] = mem
		cube[sides[3]] = replaceBits(mem, 4, cube[sides[3]], 20)

		//cube[sides[0]][2][1] = cube[sides[1]][0][1]
		cube[sides[0]] = replaceBits(cube[sides[1]], 24, cube[sides[0]], 8)
		//cube[sides[1]][0][1] = cube[sides[2]][0][1]
		cube[sides[1]] = replaceBits(cube[sides[2]], 24, cube[sides[1]], 24)
		//cube[sides[2]][0][1] = cube[sides[3]][0][1]
		cube[sides[2]] = replaceBits(cube[sides[3]], 24, cube[sides[2]], 24)
		//cube[sides[3]][0][1] = mem
		cube[sides[3]] = replaceBits(mem, 8, cube[sides[3]], 24)

		//cube[sides[0]][2][2] = cube[sides[1]][0][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 28, cube[sides[0]], 12)
		//cube[sides[1]][0][0] = cube[sides[2]][0][0]
		cube[sides[1]] = replaceBits(cube[sides[2]], 28, cube[sides[1]], 28)
		//cube[sides[2]][0][0] = cube[sides[3]][0][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 28, cube[sides[2]], 28)
		//cube[sides[3]][0][0] = mem
		cube[sides[3]] = replaceBits(mem, 12, cube[sides[3]], 28)
	} else {
		//cube[sides[0]][2][0] = cube[sides[3]][0][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 20, cube[sides[0]], 4)
		//cube[sides[3]][0][2] = cube[sides[2]][0][2]
		cube[sides[3]] = replaceBits(cube[sides[2]], 20, cube[sides[3]], 20)
		//cube[sides[2]][0][2] = cube[sides[1]][0][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 20, cube[sides[2]], 20)
		//cube[sides[1]][0][2] = mem
		cube[sides[1]] = replaceBits(mem, 4, cube[sides[1]], 20)

		//cube[sides[0]][2][1] = cube[sides[3]][0][1]
		cube[sides[0]] = replaceBits(cube[sides[3]], 24, cube[sides[0]], 8)
		//cube[sides[3]][0][1] = cube[sides[2]][0][1]
		cube[sides[3]] = replaceBits(cube[sides[2]], 24, cube[sides[3]], 24)
		//cube[sides[2]][0][1] = cube[sides[1]][0][1]
		cube[sides[2]] = replaceBits(cube[sides[1]], 24, cube[sides[2]], 24)
		//cube[sides[1]][0][1] = mem
		cube[sides[1]] = replaceBits(mem, 8, cube[sides[1]], 24)

		//cube[sides[0]][2][2] = cube[sides[3]][0][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 28, cube[sides[0]], 12)
		//cube[sides[3]][0][0] = cube[sides[2]][0][0]
		cube[sides[3]] = replaceBits(cube[sides[2]], 28, cube[sides[3]], 28)
		//cube[sides[2]][0][0] = cube[sides[1]][0][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 28, cube[sides[2]], 28)
		//cube[sides[1]][0][0] = mem
		cube[sides[1]] = replaceBits(mem, 12, cube[sides[1]], 28)
	}
}

func (env *Env) rotSide4(cube *[6]int32, way int) {
	sides := [4]int{0, 2, 5, 1}
	mem := cube[sides[0]]
	if way == 0 {
		//cube[sides[0]][2][0] = cube[sides[1]][2][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 4, cube[sides[0]], 4)
		//cube[sides[1]][2][0] = cube[sides[2]][0][2]
		cube[sides[1]] = replaceBits(cube[sides[2]], 20, cube[sides[1]], 4)
		//cube[sides[2]][0][2] = cube[sides[3]][2][0]
		cube[sides[2]] = replaceBits(cube[sides[3]], 4, cube[sides[2]], 20)
		//cube[sides[3]][2][0] = mem
		cube[sides[3]] = replaceBits(mem, 4, cube[sides[3]], 4)

		//cube[sides[0]][2][1] = cube[sides[1]][2][1]
		cube[sides[0]] = replaceBits(cube[sides[1]], 8, cube[sides[0]], 8)
		//cube[sides[1]][2][1] = cube[sides[2]][0][1]
		cube[sides[1]] = replaceBits(cube[sides[2]], 24, cube[sides[1]], 8)
		//cube[sides[2]][0][1] = cube[sides[3]][2][1]
		cube[sides[2]] = replaceBits(cube[sides[3]], 8, cube[sides[2]], 24)
		//cube[sides[3]][2][1] = mem
		cube[sides[3]] = replaceBits(mem, 8, cube[sides[3]], 8)

		//cube[sides[0]][2][2] = cube[sides[1]][2][2]
		cube[sides[0]] = replaceBits(cube[sides[1]], 12, cube[sides[0]], 12)
		//cube[sides[1]][2][2] = cube[sides[2]][0][0]
		cube[sides[1]] = replaceBits(cube[sides[2]], 4, cube[sides[1]], 12)
		//cube[sides[2]][0][0] = cube[sides[3]][2][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 12, cube[sides[2]], 4)
		//cube[sides[3]][2][2] = mem
		cube[sides[3]] = replaceBits(mem, 12, cube[sides[3]], 12)
	} else {
		//cube[sides[0]][2][0] = cube[sides[3]][2][0]
		cube[sides[0]] = replaceBits(cube[sides[3]], 4, cube[sides[0]], 4)
		//cube[sides[3]][2][0] = cube[sides[2]][0][2]
		cube[sides[3]] = replaceBits(cube[sides[2]], 20, cube[sides[3]], 4)
		//cube[sides[2]][0][2] = cube[sides[1]][2][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 4, cube[sides[2]], 20)
		//cube[sides[1]][2][0] = mem
		cube[sides[1]] = replaceBits(mem, 4, cube[sides[1]], 4)

		//cube[sides[0]][2][1] = cube[sides[3]][2][1]
		cube[sides[0]] = replaceBits(cube[sides[3]], 8, cube[sides[0]], 8)
		//cube[sides[3]][2][1] = cube[sides[2]][0][1]
		cube[sides[3]] = replaceBits(cube[sides[2]], 24, cube[sides[3]], 8)
		//cube[sides[2]][0][1] = cube[sides[1]][2][1]
		cube[sides[2]] = replaceBits(cube[sides[1]], 8, cube[sides[2]], 24)
		//cube[sides[1]][2][1] = mem
		cube[sides[1]] = replaceBits(mem, 8, cube[sides[1]], 8)

		//cube[sides[0]][2][2] = cube[sides[3]][2][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 12, cube[sides[0]], 12)
		//cube[sides[3]][2][2] = cube[sides[2]][0][0]
		cube[sides[3]] = replaceBits(cube[sides[2]], 28, cube[sides[3]], 12)
		//cube[sides[2]][0][0] = cube[sides[1]][2][2]
		cube[sides[2]] = replaceBits(cube[sides[1]], 12, cube[sides[2]], 28)
		//cube[sides[1]][2][2] = mem
		cube[sides[1]] = replaceBits(mem, 12, cube[sides[1]], 12)
	}
}

func (env *Env) rotSide5(cube *[6]int32, way int) {
	sides := [4]int{4, 2, 3, 1}
	mem := cube[sides[0]]
	if way == 0 {
		//cube[sides[0]][2][0] = cube[sides[1]][0][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 28, cube[sides[0]], 4)
		//cube[sides[1]][0][0] = cube[sides[2]][0][2]
		cube[sides[1]] = replaceBits(cube[sides[2]], 20, cube[sides[1]], 28)
		//cube[sides[2]][0][2] = cube[sides[3]][2][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 12, cube[sides[2]], 20)
		//cube[sides[3]][2][2] = mem
		cube[sides[3]] = replaceBits(mem, 4, cube[sides[3]], 12)

		//cube[sides[0]][2][1] = cube[sides[1]][1][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 0, cube[sides[0]], 8)
		//cube[sides[1]][1][0] = cube[sides[2]][0][1]
		cube[sides[1]] = replaceBits(cube[sides[2]], 24, cube[sides[1]], 0)
		//cube[sides[2]][0][1] = cube[sides[3]][1][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 16, cube[sides[2]], 24)
		//cube[sides[3]][1][2] = mem
		cube[sides[3]] = replaceBits(mem, 8, cube[sides[3]], 16)

		//cube[sides[0]][2][2] = cube[sides[1]][2][0]
		cube[sides[0]] = replaceBits(cube[sides[1]], 4, cube[sides[0]], 12)
		//cube[sides[1]][2][0] = cube[sides[2]][0][0]
		cube[sides[1]] = replaceBits(cube[sides[2]], 28, cube[sides[1]], 4)
		//cube[sides[2]][0][0] = cube[sides[3]][0][2]
		cube[sides[2]] = replaceBits(cube[sides[3]], 20, cube[sides[2]], 28)
		//cube[sides[3]][0][2] = mem
		cube[sides[3]] = replaceBits(mem, 12, cube[sides[3]], 20)
	} else {
		//cube[sides[0]][2][0] = cube[sides[3]][2][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 12, cube[sides[0]], 4)
		//cube[sides[3]][2][2] = cube[sides[2]][0][2]
		cube[sides[3]] = replaceBits(cube[sides[2]], 20, cube[sides[3]], 12)
		//cube[sides[2]][0][2] = cube[sides[1]][0][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 28, cube[sides[2]], 20)
		//cube[sides[1]][0][0] = mem
		cube[sides[1]] = replaceBits(mem, 4, cube[sides[1]], 28)

		//cube[sides[0]][2][1] = cube[sides[3]][1][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 16, cube[sides[0]], 8)
		//cube[sides[3]][1][2] = cube[sides[2]][0][1]
		cube[sides[3]] = replaceBits(cube[sides[2]], 24, cube[sides[3]], 16)
		//cube[sides[2]][0][1] = cube[sides[1]][1][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 0, cube[sides[2]], 24)
		//cube[sides[1]][1][0] = mem
		cube[sides[1]] = replaceBits(mem, 8, cube[sides[1]], 0)

		//cube[sides[0]][2][2] = cube[sides[3]][0][2]
		cube[sides[0]] = replaceBits(cube[sides[3]], 20, cube[sides[0]], 12)
		//cube[sides[3]][0][2] = cube[sides[2]][0][0]
		cube[sides[3]] = replaceBits(cube[sides[2]], 28, cube[sides[3]], 20)
		//cube[sides[2]][0][0] = cube[sides[1]][2][0]
		cube[sides[2]] = replaceBits(cube[sides[1]], 4, cube[sides[2]], 28)
		//cube[sides[1]][2][0] = mem
		cube[sides[1]] = replaceBits(mem, 12, cube[sides[1]], 4)
	}
}

func (env *Env) rotSides(cube *[6]int32, face, way int) {
	if face == 0 {
		env.rotSide0(cube, way)
	} else if face == 1 {
		env.rotSide1(cube, way)
	} else if face == 2 {
		env.rotSide2(cube, way)
	} else if face == 3 {
		env.rotSide3(cube, way)
	} else if face == 4 {
		env.rotSide4(cube, way)
	} else if face == 5 {
		env.rotSide5(cube, way)
	}
}

func (env *Env) rotate(face, way int, cube [6]int32) [6]int32 {
	if way == 0 {
		cube[face] = (cube[face] >> 8) | (cube[face] << (32 - 8))
	} else {
		cube[face] = (cube[face] << 8) | (cube[face] >> (32 - 8))
	}
	env.rotSides(&cube, face, way)
	return cube
}
