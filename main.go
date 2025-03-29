package main

import (
	"fmt"

	"github.com/waditya/go-error-handling/customerror"
)

func main() {
	checkForEvenNumbers(8)
	checkForEvenNumbers(5)
	checkForEvenNumsVerbose(8)
	checkForEvenNumsVerbose(5)
}

func checkForEvenNumbers(number int64) {
	if number%2 == 0 {
		err := customerror.NewError("Only odd numbers are allowed")
		fmt.Println(err)
		return
	}
	fmt.Println("Valid number provided")
}

func checkForEvenNumsVerbose(number int64) {
	if number%2 == 0 {
		err := customerror.NewErrorVerbose("Only odd numbers are allowed", number)
		fmt.Println(err)
		return
	}
	fmt.Println("Valid number provided")
}
