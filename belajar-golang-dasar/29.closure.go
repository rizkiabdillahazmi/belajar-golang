package main

import "fmt"

func main() {

	// Harap bijak menggunakan closure
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		// function ini bisa mengakses data di sekitarnya (counter)
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
