package main

import "fmt"

func (env *Env) debugPrint(step string, cube [6]int32) {
	fmt.Println(cube)
}
