package main

import "fmt"

func reverseLeftWords(s string, n int) string{
	strLeft := s[0:n]
	strRight := s[n:]
	return strRight + strLeft
}

func main() {
	var s string = "abcdefg"
	var k int = 2
	res := reverseLeftWords(s, k)
	fmt.Println(res)
}
