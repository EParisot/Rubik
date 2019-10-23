package main

import (
	"fmt"
	"os"
)

type Cube struct {
}

type Env struct {
	args string
	cube Cube
	res  string
}

func (env *Env) parseArgs() error {

	//return errors.New("Error : Invalid Argument")
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		arg := string(args[0])
		env := Env{
			args: arg,
		}
		err := env.parseArgs()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error : No args")
	}
}
