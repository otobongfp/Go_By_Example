package main

import (
    "fmt"
)

const s string = "constant"

func main() {
    fmt.Println("An Even Number Detector: ")

	num := 23

	if num%2 == 0{
		fmt.Println("This is an Even Number!")
	}else{
		fmt.Println("This is an Odd number")
	}
	
}