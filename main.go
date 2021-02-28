package main

import "fmt"

func main() {
	var arr = make([][]int, 5)
	arr[1] = make([]int, 2)
	fmt.Println(arr)
}
