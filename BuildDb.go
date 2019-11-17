package main

import "fmt"

func appendMove(data *map[int]int) {

}

func BuildDB() {
	data := make(map[int]int)
	appendMove(&data)
	fmt.Println(data)
}
