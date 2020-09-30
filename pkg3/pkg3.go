package pkg3

import (
	"fmt"
	"os"
)

// ShowMeOSArgs is show os arguments list
func ShowMeOSArgs() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "=="
	}

	fmt.Println(s)
	fmt.Println("where is ShowMeOSArgs?", len(os.Args))

}
