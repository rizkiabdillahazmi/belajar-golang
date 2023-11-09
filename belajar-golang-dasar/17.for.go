package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter < 10 {
	//	fmt.Println("Counter ke ", counter)
	//	counter++
	//}
	//
	//fmt.Println("Perulangan selesai")

	for counter := 1; counter < 10; counter++ {
		fmt.Println("Counter ke ", counter)
	}
	fmt.Println("Perulangan selesai")

	names := []string{"Rizki", "Abdillah", "Azmi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index ke ", index, " = ", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

}
