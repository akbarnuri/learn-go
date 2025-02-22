package main

import "fmt"

func main() {
	var hasilTambah float64 = 3 + 5
	var hasilKurang float64 = 50 - 200
	hasilKali := 10 * 10
	var hasilBagi float64 = 10.0 / 3.0

	var complexResult float64 = 1.0/3.0*5 + 100

	var complexResultAgaian float64 = hasilTambah*hasilKurang + complexResult

	var modulus int = 10 % 5

	// 10 % 4 = 10 - 8 = 2
	// 10 % 5 = 10 - 10 = 0

	fmt.Println(hasilTambah)
	fmt.Println(hasilKurang)
	fmt.Println(hasilKali)
	fmt.Println(hasilBagi)
	fmt.Println(complexResult)
	fmt.Println(complexResultAgaian)
	fmt.Println(modulus)

	//=======================================================

	/*var name string = "Akbar" + " " + "Nur"
	var lastName string = "Ichfan"

	var fullName string = name + " " + lastName

	fmt.Println(fullName) */
}
