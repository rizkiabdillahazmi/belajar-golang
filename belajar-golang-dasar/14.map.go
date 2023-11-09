package main

import "fmt"

func main() {

	//var person map[string]string = map[string]string{}
	//person["name"] = "Rizki"
	//person["city"] = "Medan"

	person := map[string]string{
		"name": "Rizki",
		"city": "Jakarta Pusat",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["city"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Rizki Abdillah Azmi"
	book["stock"] = "20"

	delete(book, "stock")
	fmt.Println(book)
}
