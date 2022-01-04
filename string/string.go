package main

import "fmt"

func main() {
	str := "proddutur nagateja"
	fmt.Println(str)
	s := []rune(str)
	fmt.Println(s)
	i := len(s) - 1
	fmt.Println(i)
	for j := 0; j < i; j++ {
		temp := s[j]
		s[j] = s[i]
		s[i] = temp
		//		s[j], s[i] = s[i], s[j]
		i--

	}
	fmt.Println(string(s))
}
