package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverseNum())
}

// 88b

//Converts number to a string, splits it to reverse and joins
func reverseNum() int {
	str := strconv.Itoa(getNum())
	arr := reverse(strings.Split(str, ""))
	revStr := strings.Join(arr, "")
	reversNum := strconv.Atoi(revStr)
	return reversNum
}

func reverse(arr []string) []string {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getNum() int {
	fmt.Println("Hey, what's your number?")
	var num int
	fmt.Scanf("%d", &num)
	return num
}
