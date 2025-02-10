package main

import (
	"fmt"
	"os"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	// if i%2 == 0 {
	// 	// 	println("Bilangan anda adalah genap")
	// 	// } else {
	// 	// 	println("Bilangan anda adalah ganjil")
	// 	// }
	// }

	//=======================================================

	// fmt.Println("Countdown")
	// for i := 10; i > 0; i-- {
	// 	fmt.Println(i)
	// }

	// fmt.Println("Lift off!")

	//=======================================================

	// var correctPassword string = "hellstar"
	// var input string

	// fmt.Print("Enter password: ")
	// fmt.Scanln(&input)

	// if input == correctPassword {
	// 	fmt.Println("Correct password")
	// } else {
	// 	fmt.Println("Incorrect password")
	// }

	//=======================================================

	// var correctPassword string = "hellstar"
	// var input string

	// for input != correctPassword {
	// 	fmt.Print("Enter password: ")
	// 	fmt.Scanln(&input)

	// 	if input != correctPassword {
	// 		fmt.Println("Incorrect password")
	// 	}
	// }

	// fmt.Println("Correct password")

	//=======================================================

	for {
		var input int
		fmt.Println("Menu,masakan apa yang ingin anda pesan?")
		fmt.Println("1. Nasi Goreng")
		fmt.Println("2. Mie Goreng")
		fmt.Println("3. Sate Ayam")
		fmt.Println("4. Keluar")

		fmt.Print("Masukan pilihan anda: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Println("Menu makanan yang anda pesan adalah Nasi Goreng")
		case 2:
			fmt.Println("Anda memesan Mie Goreng")
		case 3:
			fmt.Println("Anda memesan Sate Ayam")
		case 4:
			os.Exit(0)

		}
	}
}
