package main

import "fmt"

func main() {
	fmt.Println("Please, enter current tempreture in Farenheit:")
	var temp int
	fmt.Scanf("%d", &temp)
	celsius := (temp - 32) * 5 / 9
	fmt.Printf("Current temperature in celsius is %d", celsius)
}
