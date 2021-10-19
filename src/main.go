package main

import "fmt"

func isPalindrome(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindrome("anilina")
}
