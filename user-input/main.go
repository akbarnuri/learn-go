package main

import "fmt"

func main() {
	var name string
	fmt.Print("Masukan Nama Anda : ")
	fmt.Scanln(&name)
	fmt.Println("Halo", name)

	var age int
	fmt.Print("Masukan Umur Anda : ")
	fmt.Scanln(&age)
	fmt.Println("Umur anda sekarang " + fmt.Sprint(age))

	var age5Years int = age + 5
	fmt.Println("5 tahun lagi umur anda adalah " + fmt.Sprint(age5Years))

}
