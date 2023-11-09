package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}

	// cara yang salah, karena setelah panic, ini tidak dieksekusi lagi
	//message := recover()
	//fmt.Println("Terjadi panic", message)
}

func main() {
	// jika panic program akan langsung dihentikan
	runApp(true)
	runApp(false)
	fmt.Println("Rizki Abdillah Azmi")
}
