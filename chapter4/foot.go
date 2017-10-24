package main

import "fmt"

func length() {
	fmt.Println("Please, enter length in meters")
	var length float64
	fmt.Scanf("%d", &length)
	feet := length * 0.3048
	fmt.Printf("The length is %v feet", feet)
}
