package main

import "fmt"

func main() {
	str := "abC"
	fmt.Print(convertToCapital(str))
	//-----------
	reverce(str)
	//-----------
	width := 7
	height := 10
	printrect(width, height)
}

func convertToCapital(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] += 'A' - 'a' 
		}
	}
	return str
}

func reverce(str string) {
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(str[i])
	}
}

func printrect(width int, height int) {
	for i := 1; i <= width; i++ {
		for j := 1; j <= height; j++ {
			if i == 1 || i == width || j == 1 || j == height  
				fmt.Print("-")
			else
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
