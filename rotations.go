package main

func (env *Env) rotSide0(cube *[6][3][3]int, way int) {
	sides := [4]int{3, 2, 4, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = mem
	}
}

func (env *Env) rotSide1(cube *[6][3][3]int, way int) {
	sides := [4]int{3, 0, 4, 5}
	if way == 0 {
		mem := cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = cube[sides[2]][2][2]
		cube[sides[2]][2][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = mem
		mem = cube[sides[0]][1][2]
		cube[sides[0]][1][2] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = cube[sides[2]][1][2]
		cube[sides[2]][1][2] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = mem
		mem = cube[sides[0]][0][2]
		cube[sides[0]][0][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = mem
	} else {
		mem := cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = cube[sides[2]][2][2]
		cube[sides[2]][2][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = mem
		mem = cube[sides[0]][1][2]
		cube[sides[0]][1][2] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = cube[sides[2]][1][2]
		cube[sides[2]][1][2] = cube[sides[1]][1][2]
		cube[sides[1]][1][2] = mem
		mem = cube[sides[0]][0][2]
		cube[sides[0]][0][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = mem
	}
}

func (env *Env) rotSide2(cube *[6][3][3]int, way int) {
	sides := [4]int{3, 5, 4, 0}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = cube[sides[2]][2][0]
		cube[sides[2]][2][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = mem
		mem = cube[sides[0]][1][0]
		cube[sides[0]][1][0] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = cube[sides[2]][1][0]
		cube[sides[2]][1][0] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = mem
		mem = cube[sides[0]][0][0]
		cube[sides[0]][0][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = cube[sides[2]][2][0]
		cube[sides[2]][2][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = mem
		mem = cube[sides[0]][1][0]
		cube[sides[0]][1][0] = cube[sides[3]][1][0]
		cube[sides[3]][1][0] = cube[sides[2]][1][0]
		cube[sides[2]][1][0] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = mem
		mem = cube[sides[0]][0][0]
		cube[sides[0]][0][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = mem
	}
}

func (env *Env) rotSide3(cube *[6][3][3]int, way int) {
	sides := [4]int{5, 2, 0, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][0][1]
		cube[sides[1]][0][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][0][1]
		cube[sides[3]][0][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][0][2]
		cube[sides[1]][0][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][0][1]
		cube[sides[3]][0][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][0][1]
		cube[sides[1]][0][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][0][0]
		cube[sides[3]][0][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = mem
	}
}

func (env *Env) rotSide4(cube *[6][3][3]int, way int) {
	sides := [4]int{0, 2, 5, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][2][1]
		cube[sides[1]][2][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][2][1]
		cube[sides[3]][2][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][2][0]
		cube[sides[3]][2][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][2][1]
		cube[sides[3]][2][1] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][2][1]
		cube[sides[1]][2][1] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][2][2]
		cube[sides[1]][2][2] = mem
	}
}

func (env *Env) rotSide5(cube *[6][3][3]int, way int) {
	sides := [4]int{4, 2, 3, 1}
	if way == 0 {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = mem
	} else {
		mem := cube[sides[0]][2][0]
		cube[sides[0]][2][0] = cube[sides[3]][2][2]
		cube[sides[3]][2][2] = cube[sides[2]][0][2]
		cube[sides[2]][0][2] = cube[sides[1]][0][0]
		cube[sides[1]][0][0] = mem
		mem = cube[sides[0]][2][1]
		cube[sides[0]][2][1] = cube[sides[3]][1][2]
		cube[sides[3]][1][2] = cube[sides[2]][0][1]
		cube[sides[2]][0][1] = cube[sides[1]][1][0]
		cube[sides[1]][1][0] = mem
		mem = cube[sides[0]][2][2]
		cube[sides[0]][2][2] = cube[sides[3]][0][2]
		cube[sides[3]][0][2] = cube[sides[2]][0][0]
		cube[sides[2]][0][0] = cube[sides[1]][2][0]
		cube[sides[1]][2][0] = mem
	}
}

func (env *Env) rotSides(cube *[6][3][3]int, face, way int) {
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

func (env *Env) rotate(face, way int) [6][3][3]int {
	var cube [6][3][3]int
	for i := range env.cube {
		for j := range env.cube[i] {
			for k := range env.cube[i][j] {
				cube[i][j][k] = env.cube[i][j][k]
			}
		}
	}
	if way == 0 {
		mem := cube[face][0][0]
		cube[face][0][0] = cube[face][2][0]
		cube[face][2][0] = cube[face][2][2]
		cube[face][2][2] = cube[face][0][2]
		cube[face][0][2] = mem
		mem = cube[face][0][1]
		cube[face][0][1] = cube[face][1][0]
		cube[face][1][0] = cube[face][2][1]
		cube[face][2][1] = cube[face][1][2]
		cube[face][1][2] = mem
	} else {
		mem := cube[face][0][0]
		cube[face][0][0] = cube[face][0][2]
		cube[face][0][2] = cube[face][2][2]
		cube[face][2][2] = cube[face][2][0]
		cube[face][2][0] = mem
		mem = cube[face][0][1]
		cube[face][0][1] = cube[face][1][2]
		cube[face][1][2] = cube[face][2][1]
		cube[face][2][1] = cube[face][1][0]
		cube[face][1][0] = mem
	}
	env.rotSides(&cube, face, way)
	return cube
}
