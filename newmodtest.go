package main

import (
	"fmt"

	"github.com/kamitom/go.mod.test/pkg1"
	"github.com/kamitom/go.mod.test/pkg2"
	"github.com/kamitom/go.mod.test/pkg3"

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

	//show pkg2 QuoteToMeHello()
	fmt.Println(pkg2.QuoteToMeHello())

	pkg3.ShowMeOSArgs()

	pkg3.TestArraySlice()

	// test rune
	fmt.Println('a')
	fmt.Println('b')
	fmt.Println('c')
	fmt.Println('d')
	fmt.Println('z')

	fmt.Println("------")

	// use var
	var xx int
	xx = 100
	fmt.Println(xx)
	xx = 200
	fmt.Println(xx)
	// xx = "tom"

	// var ttname = "hood"
	fmt.Print("please input a string: ")
	var yyname string
	fmt.Scanln(&yyname)
	const ccname = "nsa"

	pkg3.KeyValuePairTest(yyname)

	fmt.Print("input two string and 用space 隔開: ")
	var t1, t2 int
	fmt.Scanln(&t1, &t2)
	fmt.Println("2個數字加總: ", t1+t2)

}
