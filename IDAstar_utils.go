package main

func getMoves(currCube CubeEnv, phase int) []CubeEnv {
	var cubeList []CubeEnv
	for face := 0; face <= 5; face++ {
		var newEnvCube CubeEnv
		// Apply move(s) (Only use authorised moved in current phase)
		newEnvCube.cube = rotate(face, 0, currCube.cube)
		var nb string
		if (phase == 1 && (face == 0 || face == 5)) || // <F2 R L U D B2>
			(phase == 2 && (face == 1 || face == 2 || face == 0 || face == 5)) || // <F2 R2 L2 U D B2>
			phase == 3 { // <F2 R2 L2 U2 D2 B2>
			newEnvCube.cube = rotate(face, 0, newEnvCube.cube)
			nb = "2"
		}
		newEnvCube.internationalMove = moves[face] + nb
		newEnvCube.cost = currCube.cost + 1
		newEnvCube.heuristic = globalHeuristic(newEnvCube, phase)
		cubeList = append(cubeList, newEnvCube)
	}
	return cubeList
}

func globalHeuristic(currCube CubeEnv, phase int) int {
	var gHeur int
	if phase == 0 {
		gHeur = isInG1(currCube)
	} else if phase == 1 {
		gHeur = isInG2(currCube)
	} else if phase == 2 {
		gHeur = isInG3(currCube)
	} else if phase == 3 {
		gHeur = isInGc(currCube)
	}
	return gHeur
}

// fixes RL and UD Edges orientation
func isInG1(currCube CubeEnv) int {
	var latFacelets int
	var topDownFacelets int
	for _, face := range []int{1, 2} {
		for _, facelet := range []int{0, 2, 4, 6} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 3 && int(currCube.cube[face]>>uint(facelet*4))&15 != 4 {
				latFacelets++
			}
		}
	}
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 2, 4, 6} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 1 && int(currCube.cube[face]>>uint(facelet*4))&15 != 2 {
				topDownFacelets++
			}
		}
	}
	return 4 - int((latFacelets+topDownFacelets)/4)
}

// Fixes UD facelets orientations and midEdges in midLayer
func isInG2(currCube CubeEnv) int {
	var topDownFacelets int
	for _, face := range []int{3, 4} {
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 == 3 || int(currCube.cube[face]>>uint(facelet*4))&15 == 4 {
				topDownFacelets++
			}
		}
	}
	var midEdges int
	for _, face := range []int{0, 1, 2, 5} {
		for _, facelet := range []int{0, 4} {
			var oppositeFace int
			if face == 0 {
				oppositeFace = 5
			} else if face == 1 {
				oppositeFace = 2
			} else if face == 2 {
				oppositeFace = 1
			} else {
				oppositeFace = 0
			}
			if int(currCube.cube[face]>>uint(facelet*4))&15 == face || int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeFace {
				midEdges++
			}
		}
	}
	return 6 - int((topDownFacelets+midEdges)/4)
}

// Fixed all topDown corners and edges orientation and corners parity
func isInG3(currCube CubeEnv) int {
	var facelets int
	var parity int
	for _, face := range []int{0, 1, 2, 3, 4, 5} {
		var oppositeFace int
		if face%2 == 0 {
			oppositeFace = face - 1
		} else {
			oppositeFace = face + 1
		}
		if oppositeFace == 6 {
			oppositeFace = 0
		} else if oppositeFace == -1 {
			oppositeFace = 5
		}
		for _, facelet := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
			if (int(currCube.cube[face]>>uint(facelet*4))&15) == face || (int(currCube.cube[face]>>uint(facelet*4))&15) == oppositeFace {
				facelets++
			}
		}
	}
	for _, face := range []int{0, 1, 2, 5} {
		for _, facelet := range []int{1, 5} {
			var nextFacelet int
			if facelet == 1 {
				nextFacelet = 3
			} else {
				nextFacelet = 7
			}
			if (int(currCube.cube[face]>>uint(facelet*4)) & 15) == (int(currCube.cube[face]>>uint(nextFacelet*4)) & 15) {
				parity++
			}
		}
	}
	if parity == 0 || parity == 8 {
		parity = 8
	} else {
		parity = 0
	}
	return 14 - int((facelets+parity)/4)
}

// Restore solved cube
func isInGc(currCube CubeEnv) int {
	var corners int
	var edges int
	for _, face := range []int{0, 1, 2, 3, 4, 5} {
		for _, facelet := range []int{1, 3, 5, 7} {
			if int((currCube.cube[face]>>uint(facelet*4))&15) == face {
				corners++
			}
		}
	}
	for _, face := range []int{0, 1, 2, 3, 4, 5} {
		for _, facelet := range []int{0, 2, 4, 6} {
			if (int(currCube.cube[face]>>uint(facelet*4)) & 15) == face {
				edges++
			}
		}
	}
	return 12 - int((corners+edges)/4)
}
