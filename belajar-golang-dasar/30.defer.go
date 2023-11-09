package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	// jika terjadi error pun tetap akan dieksekusi di akhir
	defer logging()
	fmt.Println("Application Running")
}

func main() {
	runApplication()
}
