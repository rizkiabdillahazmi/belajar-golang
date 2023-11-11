package main

import "fmt"

/**
Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
Semua variable yang mengacu ke data yang sama tidak akan berubah
Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
*/

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{"Medan", "Sumatera Utara", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Tebing Tinggi"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi bandung

	//address2 = &Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	//fmt.Println(address1)
	//fmt.Println(address2)

	*address2 = Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // address1 berubah
	fmt.Println(address2)
}
