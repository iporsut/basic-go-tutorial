package main

import "fmt"

func main() {
	num := 10
	num = num + 20
	num = 20 - 5
	num = 20 * 2
	num = 20 / 2
	fmt.Println(num)
	num++
	fmt.Println(num)
	num--
	fmt.Println(num)

	msg := "Hello" + " " + "World"
	fmt.Println(msg)

	status := true
	fmt.Println(status)
	status = !status
	fmt.Println(status)
	status = true && true
	fmt.Println(status)
	status = true && false
	fmt.Println(status)
	status = true || true
	fmt.Println(status)
	status = true || false
	fmt.Println(status)

	status = 5 == 5
	fmt.Println(status)
	status = 5 < 5
	fmt.Println(status)
	status = 6 > 5
	fmt.Println(status)
	status = 5 != 5
	fmt.Println(status)
}
