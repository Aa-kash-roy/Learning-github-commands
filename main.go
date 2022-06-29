package main

import "fmt"

func reverse() {
	// Reverse the list
	var arr []int32
	for index, _ := range arr {
		fmt.Println(index)
	}

	retrun []arr
}

func checkIsPrime() bool {
	// Check number is prime
	var num int64 = 97
	for i:=2; i<num; i++{
		if num % i == 0 {
			return false
		}
	}
	return true
}

func add(a, b int){
	// add two elements
}

func subtract(a, b int){
	// subtract two elements
}

func multiply(a, b int){
	// multiply two numbers
}

func divide(a, b float32){
	// divide two numbers
}

func main() {

	// using the methods
	reverse()
	checkIsPrime()
}
