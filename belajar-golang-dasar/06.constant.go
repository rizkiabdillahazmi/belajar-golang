package main

import "fmt"

func main() {
	const firstName string = "Rizki"
	const lastName = "Abdillah Azmi"

	// error
	//firstName = "Nizam"
	//lastName = "Yazid"

	const (
		city     string = "Indonesia"
		province        = "Jakarta Pusat"
	)

	fmt.Println(city)
	fmt.Println(province)
}
