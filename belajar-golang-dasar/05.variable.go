package main

import "fmt"

func main() {
	var name string

	name = "Rizki Abdillah Azmi"
	fmt.Println(name)

	name = "Rizki Abdillah"
	fmt.Println(name)

	var country = "Indonesia"
	fmt.Println(country)

	// Kebanyakan dev golang membuat seperti ini
	province := "Sumatera Utara"
	fmt.Println(province)

	province = "Jakarta Pusar"
	fmt.Println(province)

	var (
		firstName  = "Rizki"
		middleName = "Abdillah"
		lastName   = "Azmi"
	)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
