package main

import (
	"fmt"
	"log"
	"net/http"
)

const pi = 3.14

type Person struct {
	Name string
	Age int
}

type Employee struct {
	Name string
	Age int
	Address string
	Salary int
}

var fruits = []string{"apple", "grape", "banana", "melon"}

func main() {
	// var x int = 10
	// var decimal float64 = 2.62
	// var text string = "Hello World"
	// var condition bool = true

	// // struct
	// john := Person{Name: "John", Age: 21}
	// fmt.Println("Struct:", john)

	// //slice
	// numbers := []int{1, 2, 3, 4, 5}
	// fmt.Println("Slice:", numbers)

	// //map
	// colors := map[string]string{"red": "#da1337", "orange": "#e95a22"}
	// fmt.Println("Map:", colors)

	// fmt.Println("nilai anda adalah", x)
	// fmt.Println("integer", decimal)
	// fmt.Println("stringg", text)
	// fmt.Println("boolean", condition)
	// fmt.Println("Konstanta", pi)

	// number := 15

	// if number%2 == 0 {
	// 	fmt.Println("Genap")
	// } else {
	// 	fmt.Println("Ganjil")
	// }

		// fmt.Println("Bilangan pertama dari 5 bilangan prima :")

		// for i, count := 2, 0; count < 5; i++ {
		// 	if isPrime(i) {
		// 		fmt.Println(i)
		// 		count++
		// 	}
		// }

		// day := 4

		// switch day {
		// case 1:
		// 	fmt.Println("Senin")
		// case 2:
		// 	fmt.Println("Selasa")
		// case 3:
		// 	fmt.Println("Rabu")
		// case 4:
		// 	fmt.Println("Kamis")
		// case 5:
		// 	fmt.Println("Jumat")
		// case 6:
		// 	fmt.Println("Sabtu")
		// case 7:
		// 	fmt.Println("Minggu")
		// default:
		// 	fmt.Println("Hari tidak ditemukan")
		// }

		// x := 20
		// fmt.Println("Nilai x sebelum diubah", x)

		// p := &x
		// *p = 10
		// fmt.Println("Nilai x setelah diubah", x)

			// selectedFruits := fruits[1:2]

			// fmt.Println("Fruits:", selectedFruits)

			// emp1 := Employee{ "John", 21,"Jakarta", 10000000}
			// fmt.Println("Employee 1:", emp1)

			// emp2 := Employee{Name: "Jane", Age: 21, Address: "Jakarta", Salary: 10000000}
			// fmt.Println("Employee 2:", emp2)

			// var emp3 Employee
			// emp3.Name = "John"
			// emp3.Age = 21
			// emp3.Address = "Jakarta"
			// emp3.Salary = 10000000

			// fmt.Println("Employee 3:", emp3)

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello World")
			})

			fmt.Println("Starting web server at http://localhost:8080/")
			log.Fatal(http.ListenAndServe(":8080", nil))
}

// func isPrime(number int) bool {
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			return false
// 		}
// 	}
// 	return number > 1
// }