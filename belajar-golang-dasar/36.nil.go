package main

import "fmt"

/**
Di golang saat membuat variable tertentu maka secara otomatis akan ada default valuenya
Namun di golang ada data nil atau kosong
Nil sendiri hanya ada di interface, function, map, slice, pointer, dan channel
*/

//func ContohNilDiStringYangTidakBisa(name string) string {
//	if name == "" {
//		return nil
//	} else {
//		return name
//	}
//}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data["name"])
	}

	fmt.Println(NewMap(""))
	fmt.Println(NewMap("Rizki"))
}
