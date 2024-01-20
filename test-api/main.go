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
	// var x int = 10
	var decimal float64 = 7.5
	var text string = "Golang"
	var condition bool = true

	fmt.Println("Hello World", PI)
	fmt.Println("Hello World", decimal)
	fmt.Println("Hello World", text)
	fmt.Println("Hello World", condition)

	// struct
	john := Person{Name: "John", Age: 10}
	fmt.Println("struct", john)

	// slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("slice", numbers)

	// map
	colors := map[string]string{"red": "#ff0000"}
	fmt.Println("map:", colors)

	// if
	// number := 15

	// if number % 2 == 0 {
	// 	fmt.Println(number, "is even")
	// } else {
	// 	fmt.Println(number, "is odd")
	// }

	// for
	for i, count := 2, 0; count < 5; i++ {
		if isPrime(i){
			fmt.Println(i)
			count++
		}
	}

	//switch
	day := 3

	switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Other day")
	}

	// pointer
	x := 20
	fmt.Println("nilai awal", x)

	p := &x
	*p = 30

	fmt.Println("nilai baru", x)

}

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}