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

func main() {

	// using the methods
	reverse()
	checkIsPrime()
}
