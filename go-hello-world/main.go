// Learn Go in 12 minutes: https://www.youtube.com/watch?v=C8LgvuEBraI

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// vars
	fmt.Println("VARIABLES:")

	var hello_world string = "Hello, World!"
	fmt.Println(hello_world)

	x := 1
	y := 4
	sum := x + y
	fmt.Println(sum)

	// if conditionals
	fmt.Println("\nCONDITIONALS:")

	if x < 4 {
		fmt.Println("x is less than 4")
	} else if x > 4 {
		fmt.Println("x is greater than 4")
	} else {
		fmt.Println("x is 4")
	}

	// Arrays
	fmt.Println("\nARRAYS & SLICES:")

	var array [5]int
	fmt.Println(array)

	array_2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array_2)

	// slices - abstraction on arrays without fixed number
	array_3 := []int{1, 2, 3, 4, 5}
	array_3 = append(array_3, 6)
	fmt.Println(array_3)

	// maps - key,value pairs
	fmt.Println("\nMAPS:")

	map_1 := make(map[string]int)
	map_1["example_key_1"] = 1
	fmt.Println(map_1["example_key_1"])

	map_1["example_key_2"] = 3
	map_1["example_key_3"] = 5
	fmt.Println(map_1)

	delete(map_1, "example_key_3")
	fmt.Println(map_1)

	// for loops
	fmt.Println("\nLOOPS:")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	array_4 := []int{1, 2, 3}
	for index, value := range array_4 {
		fmt.Println("index:", index, "value:", value)
	}

	map_2 := make(map[string]int)
	map_2["example_key_1"] = 1
	map_2["example_key_2"] = 2
	map_2["example_key_3"] = 3
	for key, value := range map_2 {
		fmt.Println("key:", key, "value:", value)
	}

	// functions
	fmt.Println("\nFUNCTIONS:")

	result := example_sum(2, 3)
	fmt.Println(result)

	result_2, err := sqrt(64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result_2)
	}

	result_3, err := sqrt(-64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result_3)
	}

	// structs
	fmt.Println("\nSTRUCTS:")

	person := person{name: "Fred", age: 23}
	fmt.Println(person)
	fmt.Println(person.name)

	// pointers
	fmt.Println("\nPOINTERS:")

	z := 5
	fmt.Println(z)
	fmt.Println(&z) // memory address
	inc(&z)
	fmt.Println(z)

}

// functions
func example_sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

// structs
type person struct {
	name string
	age  int
}

// pointers
func inc(x *int) {
	*x++ // asterisk to dereference pointer to increment value at memory address
}
