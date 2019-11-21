package main

func getMoves(currCube CubeEnv, phase int) [6]CubeEnv {
	var cubeList [6]CubeEnv
	for move := 0; move < 6; move++ {
		var newEnvCube CubeEnv
		// Apply move(s) (Only use authorised moved in current phase)
		newEnvCube.cube = rotate(move, 0, currCube.cube)
		var nb string
		if (phase == 1 && (move == 0 || move == 5)) || // <F2 R L U D B2>
			(phase == 2 && (move == 1 || move == 2 || move == 0 || move == 5)) || // <F2 R2 L2 U D B2>
			phase == 3 { // <F2 R2 L2 U2 D2 B2>
			newEnvCube.cube = rotate(move, 0, newEnvCube.cube)
			nb = "2"
		}
		newEnvCube.internationalMove = moves[move] + nb
		newEnvCube.cost = currCube.cost + 1
		newEnvCube.heuristic = globalHeuristic(newEnvCube, phase)
		cubeList[move] = newEnvCube
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
	var latEdges int
	var topDownEdges int
	var faceFacelet int
	for _, face := range [2]int{1, 2} {
		for _, facelet := range [4]int{0, 2, 4, 6} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 3 && int(currCube.cube[face]>>uint(facelet*4))&15 != 4 {
				latEdges++
			}
		}
	}
	for _, face := range [2]int{3, 4} {
		for _, facelet := range [4]int{0, 2, 4, 6} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 != 1 && int(currCube.cube[face]>>uint(facelet*4))&15 != 2 {
				topDownEdges++
			}
		}
	}
	for _, face := range [2]int{0, 5} {
		for _, facelet := range [4]int{0, 2, 4, 6} {
			var compatibleFace1 int
			var oppositeCompatibleFace1 int
			var compatibleFace2 int
			var oppositeCompatibleFace2 int
			if facelet == 0 || facelet == 4 {
				compatibleFace1 = 3
				oppositeCompatibleFace1 = 4
				compatibleFace2 = 0
				oppositeCompatibleFace2 = 5
			} else if facelet == 2 || facelet == 4 {
				compatibleFace1 = 1
				oppositeCompatibleFace1 = 2
				compatibleFace2 = 0
				oppositeCompatibleFace2 = 5
			}
			if int(currCube.cube[face]>>uint(facelet*4))&15 == compatibleFace1 || int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeCompatibleFace1 ||
				int(currCube.cube[face]>>uint(facelet*4))&15 == compatibleFace2 || int(currCube.cube[face]>>uint(facelet*4))&15 == oppositeCompatibleFace2 {
				faceFacelet++
			}
		}
	}
	return 12 - int((latEdges+topDownEdges+faceFacelet)/2)
}

// Fixes UD facelets orientations and midEdges in midLayer
func isInG2(currCube CubeEnv) int {
	var topDownFacelets int
	for _, face := range [2]int{3, 4} {
		for _, facelet := range [8]int{0, 1, 2, 3, 4, 5, 6, 7} {
			if int(currCube.cube[face]>>uint(facelet*4))&15 == 3 || int(currCube.cube[face]>>uint(facelet*4))&15 == 4 {
				topDownFacelets++
			}
		}
	}
	return 8 - int((topDownFacelets)/2)
}

// Fixed all topDown corners and edges orientation and corners parity
func isInG3(currCube CubeEnv) int {
	var facelets int
	var parity int
	for _, face := range [6]int{0, 1, 2, 3, 4, 5} {
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
		for _, facelet := range [8]int{0, 1, 2, 3, 4, 5, 6, 7} {
			if (int(currCube.cube[face]>>uint(facelet*4))&15) == face || (int(currCube.cube[face]>>uint(facelet*4))&15) == oppositeFace {
				facelets++
			}
		}
	}
	for _, face := range [4]int{0, 1, 2, 5} {
		for _, facelet := range [2]int{1, 5} {
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
	for _, face := range [6]int{0, 1, 2, 3, 4, 5} {
		for _, facelet := range [4]int{1, 3, 5, 7} {
			if int((currCube.cube[face]>>uint(facelet*4))&15) == face {
				corners++
			}
		}
	}
	for _, face := range [6]int{0, 1, 2, 3, 4, 5} {
		for _, facelet := range [4]int{0, 2, 4, 6} {
			if (int(currCube.cube[face]>>uint(facelet*4)) & 15) == face {
				edges++
			}
		}
	}
	return 12 - int((corners+edges)/4)
}
