package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya")
	fmt.Println("jumlah Element :", len(buah))
	fmt.Println("isi Element :", (buah))
	fmt.Println("isi Element -: 0", buah[0])
	fmt.Println("isi Element -: 1", buah[1])
	fmt.Println("isi Element -: 2", buah[2])
	fmt.Println("isi Element -: 3", buah[3])
	fmt.Println("isi Element -: 4", buah[4])

}
