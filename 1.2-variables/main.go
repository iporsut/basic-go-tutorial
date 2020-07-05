package main

import "fmt"

func main() {
	num := 10
	fmt.Printf("%T, %v\n", num, num)
	str := "Hello, World"
	fmt.Printf("%T, %v\n", str, str)
	price := 123.567
	fmt.Printf("%T, %v\n", price, price)
	status := true
	fmt.Printf("%T, %v\n", status, status)
	ch := 'ก'
	fmt.Printf("%T, %v\n", ch, ch)
}
