package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Rizki"
	names[1] = "Abdillah"
	names[2] = "Azmi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)

	// ... hanya bisa digunakan ketika langsung assign valuenya
	var data = [...]int{
		19,
		20,
		21,
		22,
		23,
	}

	fmt.Println(data)
	fmt.Println(data[2])
	fmt.Println(len(data))
}
