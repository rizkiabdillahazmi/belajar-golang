package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {

	result := getHello("Rizki")
	fmt.Println(result)

	fmt.Println(getHello("Abdillah"))
	fmt.Println(getHello("Azmi"))

}
