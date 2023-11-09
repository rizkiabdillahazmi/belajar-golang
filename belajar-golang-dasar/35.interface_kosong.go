package main

import "fmt"

func Ups() any {
	return 1
}

func UpsLagi() interface{} {
	return "Ups"
}

func main() {
	tes := Ups()
	fmt.Println(tes)

	tes2 := UpsLagi()
	fmt.Println(tes2)
}
