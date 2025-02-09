package main

import "fmt"

func main() {
	// var number int

	// fmt.Print("Enter a number: ")
	// fmt.Scanln(&number)

	// if number%2 == 0 {
	// 	fmt.Println("Bilangan anda adalah genap")
	// } else {
	// 	fmt.Println("Bilangan anda adalah ganjil")
	// }

	//=======================================================

	// var number int

	// fmt.Print("Masukan nilai anda: ")
	// fmt.Scanln(&number)

	// if number >= 70 {
	// 	fmt.Println("Selamat anda lulus")
	// } else {
	// 	fmt.Println("Maaf anda tidak lulus")
	// }

	//=======================================================

	var x int

	fmt.Print("Masukan nilai anda: ")
	fmt.Scanln(&x)

	if 100 >= x && x >= 90 {
		fmt.Println("Selamat nilai anda A")
	} else if 90 > x && x >= 80 {
		fmt.Println("Maaf nilai anda B")
	} else if 80 > x && x >= 70 {
		fmt.Println("Maaf nilai anda C")
	} else if 70 > x && x >= 60 {
		fmt.Println("Maaf nilai anda D")
	} else {
		fmt.Println("Maaf nilai anda E")
	}
}
