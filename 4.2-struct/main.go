package main

import (
	"examplestruct/profile"
	"fmt"
)

func main() {
	p := profile.Profile{
		username: "Weerasak",
	}
	fmt.Println(p)
	fmt.Println(p.username)
}
