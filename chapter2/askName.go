package main

import "fmt"

func main() {
	fmt.Println("Please, enter your name:")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello, %s", name)
}

// why programm didn't wait for user output when line 8 was empty string
