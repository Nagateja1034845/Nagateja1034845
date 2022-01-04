package main

import "fmt"

func main() {
	str := "abcdefghijklmnopqrstuvwxyz!@#$"
	vowel := 0
	consonents := 0
	specialChar := 0
	vowelslice := make([]string, 0)
	consonentsslice := make([]string, 0)
	specialCharslice := make([]string, 0)
	for _, value := range str {
		switch value >= 33 && value <= 126 {
		case value >= 33 && value <= 47 || value >= 58 && value <= 64 || value >= 91 && value <= 96 || value >= 123 && value < 126:
			specialChar++
			specialCharslice = append(specialCharslice, string(value))
		case value >= 65 && value <= 90 || value >= 97 && value <= 127:

			consonents++
			consonentsslice = append(consonentsslice, string(value))

			if value == 97 || value == 101 || value == 105 || value == 111 || value == 117 {

				vowel++
				vowelslice = append(vowelslice, string(value))
			}
		case value == 97 || value == 101 || value == 105 || value == 111 || value == 117:

			vowel++
			vowelslice = append(vowelslice, string(value))

		}
	}
	fmt.Print("No. of Vowels in string=", vowel)
	fmt.Println("\tvowels Present in string", vowelslice)
	fmt.Print("No. of consonents in string=", consonents)
	fmt.Println("\tconsonents", consonentsslice)
	fmt.Print("No. of Special charaters in string=", specialChar)
	fmt.Println("\tspecial characters", specialCharslice)
}
