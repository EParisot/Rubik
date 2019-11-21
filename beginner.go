package main

import (
	"fmt"
)

const ORANGE = 0
const GREEN = 1
const BLUE = 2
const WHITE = 3
const YELLOW = 4
const RED = 5

func (env *Env) firstLayer() {
	env.faceFirstLayer(ORANGE)
	env.faceFirstLayer(GREEN)
	env.faceFirstLayer(RED)
	env.faceFirstLayer(BLUE)
}

func (env *Env) beginner() {
	// the main documentation are from https://ruwix.com/the-rubiks-cube/how-to-solve-the-rubiks-cube-beginners-method/

	env.firstCross() // for win lot of turns,can make a A* (7 moves max, max 10s)
	env.firstLayer()
	env.secondLayer()
	env.topcross()
	env.topedges()
	env.positiontopcorner()
	env.orientedlastcorner()
	env.res = env.res[:len(env.res)-1]
	env.res = parseOutput(env.res)
	fmt.Println(env.res)

}
