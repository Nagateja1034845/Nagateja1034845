package main

import "fmt"

func main() {
	if num := 78; num < 50 {
		fmt.Println(num, "is lessthan 50")
	} else if num >= 50 && num <= 100 {
		fmt.Println(num, "is less than 100")
	} else {
		fmt.Println(num, "is graterthan 100")
	}
}
