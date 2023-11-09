package main

import "fmt"

func main() {
	name := "Rizkiiiiii"

	switch name {
	case "Rizki":
		fmt.Println("Hi Rizki")
	case "Abdillah":
		fmt.Println("Hi Abdillah")
	default:
		fmt.Println("Hi, dengan siapa?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
