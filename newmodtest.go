package main

import (
	"fmt"

	"github.com/appleboy/com/random"

	"github.com/kamitom/stringrv"
)

func main() {
	fmt.Println("hola, no more gopath approach!")

	fmt.Println("random data: ", random.String(20))

	fmt.Println(stringrv.Reverse("hola, string rv"))
}
