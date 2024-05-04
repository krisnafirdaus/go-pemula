package main

import "fmt"

const Pi = 3.14

type Person struct {
	Name string
	Age int
}

func main(){
	var x int = 0
	fmt.Println("nilai x adalah", x)

	var integer int = 10;
	var float float32 = 10.55;
	var decimal float64 = 15.555555;
	var text string = "Hello World";
	var boolean bool = true;

	fmt.Println("nilai integer adalah", integer)
	fmt.Println("nilai float adalah", float)
	fmt.Println("nilai decimal adalah", decimal)
	fmt.Println("nilai text adalah", text)
	fmt.Println("nilai boolean adalah", boolean)

	// struct
	john := Person{"John Doe", 25}
	fmt.Println("Struch", john)

	// Slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice", numbers)

	// Map
	colors := map [string]string{"red": "#ff0000", "blue": "#0000ff"}
	fmt.Println("Map", colors)

	//if
	angka := 10

	if angka > 5 {
		fmt.Println("angka lebih besar dari 5")
	} else {
		fmt.Println("angka lebih kecil dari 5")
	}

	if angka%2 == 0 {
		fmt.Println("angka genap")
	} else {
		fmt.Println("angka ganjil")
	}

	//for
	for i := 0; i < 5; i++ {
		fmt.Println("perulangan ke", i)
	}

	//swicth
	day := 4 

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

	//pointer
	y := 20
	fmt.Println("nilai y adalah", y)

	p := &y
	*p = 30

	fmt.Println("nilai y adalah", y)

	//operator aritmatika
	a := 10
	b := 5

	fmt.Println("a + b =", a + b)
	fmt.Println("a - b =", a - b)
	fmt.Println("a * b =", a * b)
	fmt.Println("a / b =", a / b)
	fmt.Println("a % b =", a % b)

	// operator logika
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

	// operator perbandingan
	fmt.Println("10 > 5 =", 10 > 5)
	fmt.Println("10 < 5 =", 10 < 5)
	fmt.Println("10 >= 5 =", 10 >= 5)
	fmt.Println("10 <= 5 =", 10 <= 5)
	fmt.Println("10 == 5 =", 10 == 5)
	fmt.Println("10 != 5 =", 10 != 5)

}

func DisplayNumber(number int){
	fmt.Println("nilai number adalah", number)
}