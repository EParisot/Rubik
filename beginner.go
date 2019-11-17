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

func (env *Env) firstlayer() {
	//fmt.Println("Face Orange : ")
	env.faceFirstLayer(ORANGE)
	//fmt.Println("Face Green : ")
	env.faceFirstLayer(GREEN)
	//fmt.Println("Face Red : ")
	env.faceFirstLayer(RED)
	env.faceFirstLayer(BLUE)
}

func (env *Env) beginner() {

	env.first_cross() // POur gagner beaucoup de coup, possible de faire un A* en - de 10s
	env.firstlayer()
	env.secondlayer()
	env.topcross()
	env.topedges()
	env.positiontopcorner()
	env.orientedlastcorner()
	env.res = env.res[:len(env.res)-1]
	fmt.Println(env.res)

}
