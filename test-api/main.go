package main

import (
	"fmt"
)

const PI = 3.14

type Person struct {
	Name string
	Age int
} 

func main() {
	var x int = 10
	var y int = 20
 	var integer int = 10
	var decimal float64 = 10.10
	var text string = "Hello World"
	var condition bool = true

	// struct
	john := Person{"John", 25}
	fmt.Println("struct: ", john)

	//slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", numbers)

	//map
	person := map[string]string{"name": "John", "age": "25"}
	fmt.Println("map: ", person)

	//map number
	phone := map[int]string{1: "John", 2: "Doe"}
	fmt.Println("map number: ", phone)

	number := 15

	if number%2 == 0 {
		fmt.Println("Ganjil")
	} else {
		fmt.Println("Genap")
	}

	for i, count := 2, 0; count < 5; i++ {
		if isPrime(i){
			fmt.Println(i)
			count++
		}
	}

	// switch 
	day := 7

	switch day {
		case 1:
			fmt.Println("Senin")
		case 2:
			fmt.Println("Selasa")
		case 3:
			fmt.Println("Rabu")
		case 4:
			fmt.Println("Kamis")
		case 5:
			fmt.Println("Jumat")
		case 6:
			fmt.Println("Sabtu")
		case 7:
			fmt.Println("Minggu")
		default:
			fmt.Println("Tidak ada hari")
	}

	xy := 20

	fmt.Println("Nilai xy adalah", xy)

	py := &xy
	*py = 50
	
	fmt.Println("Nilai xy adalah", xy)

	fmt.Println("Nilai x adalah", integer)
	fmt.Println("Nilai x adalah", decimal)
	fmt.Println("Nilai x adalah", text)
	fmt.Println("Nilai x adalah", condition)
	fmt.Println("Nilai x adalah", x)
	fmt.Println("Nilai x adalah", y)

	DisplayNumber(x)

	a := 10
	b := 3

	fmt.Println("Penjumlahan", a + b)
	fmt.Println("Pengurangan", a - b)
	fmt.Println("Perkalian", a * b)
	fmt.Println("Pembagian", a / b)
	fmt.Println("Sisa Bagi", a % b)

	c := true
	d := false

	fmt.Println("AND", c && d)
	fmt.Println("OR", c || d)
	fmt.Println("NOT", !d)

	// perbandingan
	fmt.Println("Sama dengan", a == b)
	fmt.Println("Tidak sama dengan", a != b)
	fmt.Println("Lebih besar dari", a > b)
	fmt.Println("Lebih kecil dari", a < b)
	fmt.Println("Lebih besar sama dengan", a >= b)
	fmt.Println("Lebih kecil sama dengan", a <= b)
}

func DisplayNumber(number int) {
	fmt.Println("angka", number)
}

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}