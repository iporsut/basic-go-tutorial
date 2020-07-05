package main

import "fmt"

func main() {
	if num := 5; num == 5 {
		fmt.Println("num is equal 5")
	}

	switch num := 5; num {
	case 5:
		fmt.Println("num is equal 5")
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		i++
	}

	n := 0
	for {
		if n == 5 {
			break
		}
		if n == 3 {
			n++
			continue
		}
		fmt.Println(n)
		n++
	}
}
