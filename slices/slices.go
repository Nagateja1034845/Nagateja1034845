package main

import "fmt"

func addElementInmiddle(slice []string, s int, value string) []string {
	slice = append(slice[:s], slice[s-1:]...)
	slice[s] = "10"
	return slice
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	employe := make([]string, 5)
	fmt.Println(employe)
	employe[0] = "1"
	employe[1] = "2"
	employe[2] = "3"
	employe[3] = "4"
	employe[4] = "5"
	fmt.Println(employe)
	fmt.Println(len(employe))

	employe = append(employe, "6")
	employe = append(employe, "7")
	employe = append(employe, "8")
	fmt.Println(employe)
	fmt.Println(len(employe))
	var a int
	fmt.Print("enter a=")
	fmt.Scanln(&a)
	//fmt.Println("a=")
	fmt.Println("a=", a)

	var copyarray = make([]string, len(employe))

	copy(copyarray, employe)
	fmt.Println(copyarray)
	l := employe[2:5]
	fmt.Println("employel1:", l)

	m := employe[2:]
	fmt.Println("sl1:", m)

	n := employe[:5]
	fmt.Println("employen1:", n)

	employe = addElementInmiddle(employe, 3, "10")
	fmt.Println(employe)

	employe = remove(employe, 5)
	fmt.Println(employe)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
