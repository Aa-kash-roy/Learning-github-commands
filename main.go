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
	return a + b
}

func subtract(a, b int){
	// subtract two elements
	return a - b
}

func multiply(a, b int){
	// multiply two numbers
	return a * b
}

func divide(a, b float32){
	// divide two numbers
	return a / b
}

func help(){
	// This method is helper function
}

func main() {

	// using the methods
	reverse()
	checkIsPrime()
}
