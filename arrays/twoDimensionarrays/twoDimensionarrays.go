package main

import "fmt"

func main() {
	arr := [2][2][3]int{
		{{1, 2, 3},
			{
				5, 6, 7,
			},
		},
		{{8, 9, 10},
			{
				11, 12, 13,
			},
		},
	}
	fmt.Println(arr)

	for _, value := range arr {
		fmt.Println(value)

	}

}
