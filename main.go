package main

import "fmt"

func main() {
	var a any
	a = nil
	fmt.Println("2")
	b := fmt.Sprintf("%v", a)
	fmt.Println(b)
	fmt.Println("1")
}
