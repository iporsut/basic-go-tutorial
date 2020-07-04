package display

import (
	"fmt"
	"strings"
)

func Title(name string) {
	msg := strings.ToUpper("Hello " + name)
	fmt.Println(msg)
}
