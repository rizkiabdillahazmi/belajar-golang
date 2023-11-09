package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	misal := getGoodBye
	contoh := getGoodBye

	fmt.Println(misal("Rizki"))
	fmt.Println(contoh("Rizki"))
}
