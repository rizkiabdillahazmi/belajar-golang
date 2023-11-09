package main

import "fmt"

func main() {
	name := "Rizki"

	if name == "Rizki" {
		fmt.Println("Hi Rizki")
	} else if name == "Abdillah" {
		fmt.Println("Hi Abdillah")
	} else if name == "Azmi" {
		fmt.Println("Hi Azmi")
	} else {
		fmt.Println("Hi, dengan siapa?")
	}

	// Short If statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
