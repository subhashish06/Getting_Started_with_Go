package main

import "fmt"


func main() {
	var first string
	fmt.Print("What is your first name: ")	
	fmt.Scanln(&first)
	
	var second string
	fmt.Print("What is your second name: ")	
	fmt.Scanln(&second)
	
	fmt.Println("Hello", first, second)
}
