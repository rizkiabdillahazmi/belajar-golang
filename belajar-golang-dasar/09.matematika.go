package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = 5
	var d = 8
	var e = a + b - c*d

	fmt.Println(e)

	var i = 10
	i += 10
	fmt.Println(i)

	i += 20
	fmt.Println(i)

	// Unary Operator
	var j = 1
	j++
	j++
	fmt.Println(j)

	j--
	j--
	fmt.Println(j)
}
