//Implement various string function in Golang [substring, replace, length, reverse]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	loop := true
	fmt.Println("Enter the String: ")
	in := bufio.NewReader(os.Stdin)
	value, _ := in.ReadString('\n')
	value = strings.Replace(value, "\n", "", -1)
	for loop {
		fmt.Println("Choose an Operation:")
		fmt.Println("1. Sub String, 2. Replace, 3. Length, 4. Reverse, 5. Exit")
		var a int
		fmt.Scan(&a)
		switch a {
		case 1:
			subString(value)
		case 2:
			replace(value)
		case 3:
			length(value)
		case 4:
			reverse(value)
		case 5:
			loop = false
		default:
			fmt.Println("Invalid Input")
		}
	}
}

func subString(value string) {
	runes := []rune(value)
	fmt.Printf("Enter Starting position for Sub String: ")
	var start int
	fmt.Scan(&start)
	fmt.Printf("Enter Ending position for Sub String: ")
	var end int
	fmt.Scan(&end)
	mySubString := string(runes[start:end])
	fmt.Println("Your Sub String is: " + mySubString)
}

func replace(value string) {
	fmt.Printf("Enter character to be replaced: ")
	var a string
	fmt.Scan(&a)
	fmt.Printf("Enter character to be substituted: ")
	var b string
	fmt.Scan(&b)
	a = strings.Replace(value, a, b, -1)
	fmt.Println("Your New Replaced String is: " + a)
}
func length(value string) {
	a := len(value)
	fmt.Printf("Length of your String is: %v\n", a)
}
func reverse(value string) {
	n := 0
	rune := make([]rune, len(value))
	for _, r := range value {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
	fmt.Println(output)
}
