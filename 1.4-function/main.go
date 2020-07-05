package main

import (
	"fmt"
	"strings"
)

func title(name string) {
	msg := strings.ToUpper("Hello " + name)
	fmt.Println(msg)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	title("Por")
	title("May")
	num := add(10, 20)
	fmt.Println(num)
}
