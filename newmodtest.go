package main

import (
	"fmt"

	"github.com/kamitom/go.mod.test/pkg1"
	"github.com/kamitom/go.mod.test/pkg2"

	"github.com/appleboy/com/random"

	"github.com/kamitom/stringrv"
)

func main() {
	fmt.Println("hola, no more gopath approach!")

	fmt.Println("random data: ", random.String(20))

	fmt.Println(stringrv.Reverse("hola, github/kamitom"))

	// use my own package: pkg1
	pkg1.HolaPkg1()

	fmt.Println(pkg1.GiveMeFakeFullName())

	fmt.Println("random is: ", pkg2.TryRandom(18))

	//show pkg2 QuoteToMe()
	fmt.Println(pkg2.QuoteToMe())

}
