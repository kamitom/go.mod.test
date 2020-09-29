package main

import (
	"fmt"

	"github.com/kamitom/go.mod.test/pkg1"

	"github.com/appleboy/com/random"

	"github.com/kamitom/stringrv"
)

func main() {
	fmt.Println("hola, no more gopath approach!")

	fmt.Println("random data: ", random.String(20))

	fmt.Println(stringrv.Reverse("hola, github/kamitom"))

	// use my own package: pkg1
	pkg1.HolaPkg()

	fmt.Println(pkg1.GiveMeFakeFullName())
}
